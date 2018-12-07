package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/mesg-foundation/core/service"
	"github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

const (
	sidKeyPrefix  = "sid_"
	hashKeyPrefix = "hash_"
)

var (
	errCannotSaveWithoutHash = errors.New("database: can't save service without hash")
	errCannotSaveWithoutSID  = errors.New("database: can't save service without sid")
	errSIDSameLen            = errors.New("database: sid can't have the same length as hash")
)

// ServiceDB describes the API of database package.
type ServiceDB interface {
	// Save saves a service to database.
	Save(s *service.Service) error

	// Get gets a service from database by its unique id
	// or unique sid.
	Get(hashOrSID string) (*service.Service, error)

	// Delete deletes a service from database by its unique id
	// or unique sid.
	Delete(hashOrSID string) error

	// All returns all services from database.
	All() ([]*service.Service, error)

	// Close closes underlying database connection.
	Close() error
}

// LevelDBServiceDB is a database for storing service definition.
type LevelDBServiceDB struct {
	db *leveldb.DB
}

// NewServiceDB returns the database which is located under given path.
func NewServiceDB(path string) (*LevelDBServiceDB, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &LevelDBServiceDB{db: db}, nil
}

// marshal returns the byte slice from service.
func (d *LevelDBServiceDB) marshal(s *service.Service) ([]byte, error) {
	return json.Marshal(s)
}

// unmarshal returns the service from byte slice.
func (d *LevelDBServiceDB) unmarshal(id string, value []byte) (*service.Service, error) {
	var s service.Service
	if err := json.Unmarshal(value, &s); err != nil {
		return nil, &DecodeError{ID: id}
	}
	return &s, nil
}

// All returns every service in database.
func (d *LevelDBServiceDB) All() ([]*service.Service, error) {
	var (
		services []*service.Service
		iter     = d.db.NewIterator(util.BytesPrefix([]byte(hashKeyPrefix)), nil)
	)
	for iter.Next() {
		id := strings.TrimPrefix(string(iter.Key()), hashKeyPrefix)
		s, err := d.unmarshal(id, iter.Value())
		if err != nil {
			// NOTE: Ignore all decode errors (possibly due to a service
			// structure change or database corruption)
			if decodeErr, ok := err.(*DecodeError); ok {
				logrus.WithField("service", decodeErr.ID).Warning(decodeErr.Error())
				continue
			}
			iter.Release()
			return nil, err
		}
		services = append(services, s)
	}
	iter.Release()
	return services, iter.Error()
}

// Delete deletes service from database.
func (d *LevelDBServiceDB) Delete(hashOrSID string) error {
	tx, err := d.db.OpenTransaction()
	if err != nil {
		return err
	}
	if err := d.delete(tx, hashOrSID); err != nil {
		tx.Discard()
		return err
	}
	return tx.Commit()
}

// delete deletes service from database by using r reader.
func (d *LevelDBServiceDB) delete(tx *leveldb.Transaction, hashOrSID string) error {
	s, err := d.get(tx, hashOrSID)
	if err != nil {
		return err
	}
	if err := tx.Delete([]byte(hashKeyPrefix+s.Hash), nil); err != nil {
		return err
	}
	return tx.Delete([]byte(sidKeyPrefix+s.SID), nil)
}

// Get retrives service from database.
func (d *LevelDBServiceDB) Get(hashOrSID string) (*service.Service, error) {
	tx, err := d.db.OpenTransaction()
	if err != nil {
		return nil, err
	}
	s, err := d.get(tx, hashOrSID)
	if err != nil {
		tx.Discard()
		return nil, err
	}
	return s, tx.Commit()
}

// get retrives service from database by using r reader.
func (d *LevelDBServiceDB) get(r leveldb.Reader, hashOrSID string) (*service.Service, error) {
	hash := hashOrSID

	// check if key is a sid, if yes then get hash.
	bid, err := r.Get([]byte(sidKeyPrefix+hashOrSID), nil)
	if err != nil && err != leveldb.ErrNotFound {
		return nil, err
	} else if err == nil {
		hash = string(bid)
	}

	// get the service
	b, err := r.Get([]byte(hashKeyPrefix+hash), nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return nil, &ErrNotFound{ID: hashOrSID}
		}
		return nil, err
	}
	return d.unmarshal(hashOrSID, b)
}

// Save stores service in database.
// If there is an another service that uses the same sid, it'll be deleted.
func (d *LevelDBServiceDB) Save(s *service.Service) error {
	// check service
	if s.Hash == "" {
		return errCannotSaveWithoutHash
	}
	if s.SID == "" {
		return errCannotSaveWithoutSID
	}
	if len(s.Hash) == len(s.SID) {
		return errSIDSameLen
	}

	// open database transaction
	tx, err := d.db.OpenTransaction()
	if err != nil {
		return err
	}

	// delete existent service that has the same sid.
	if err := d.delete(tx, s.SID); err != nil && !IsErrNotFound(err) {
		tx.Discard()
		return err
	}

	// encode service
	b, err := d.marshal(s)
	if err != nil {
		tx.Discard()
		return err
	}

	// save service with hash.
	if err := tx.Put([]byte(hashKeyPrefix+s.Hash), b, nil); err != nil {
		tx.Discard()
		return err
	}

	// save sid-hash pair of service.
	if err := tx.Put([]byte(sidKeyPrefix+s.SID), []byte(s.Hash), nil); err != nil {
		tx.Discard()
		return err
	}

	return tx.Commit()
}

// Close closes database.
func (d *LevelDBServiceDB) Close() error {
	return d.db.Close()
}

// ErrNotFound is an not found error.
type ErrNotFound struct {
	ID string
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("database: service %s not found", e.ID)
}

// DecodeError represents a service impossible to decode.
type DecodeError struct {
	ID string
}

func (e *DecodeError) Error() string {
	return fmt.Sprintf("database: could not decode service %q", e.ID)
}

// IsErrNotFound returns true if err is type of ErrNotFound, false otherwise.
func IsErrNotFound(err error) bool {
	_, ok := err.(*ErrNotFound)
	return ok
}
