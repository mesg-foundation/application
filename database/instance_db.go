package database

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mesg-foundation/engine/hash"
	"github.com/mesg-foundation/engine/instance"
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	errSaveInstanceWithoutHash = errors.New("database: can't save instance without hash")
)

// InstanceDB describes the API of Instance database.
type InstanceDB interface {
	// Exist check if instance with given hash exist.
	Exist(hash hash.Hash) (bool, error)

	// Get retrives instance by instance hash.
	Get(hash hash.Hash) (*instance.Instance, error)

	// GetAll retrieves all instances.
	GetAll() ([]*instance.Instance, error)

	// Save saves instance to database.
	Save(i *instance.Instance) error

	// Delete an instance by instance hash.
	Delete(hash hash.Hash) error

	// Close closes underlying database connection.
	Close() error
}

// LevelDBInstanceDB is a database for storing services' instances.
type LevelDBInstanceDB struct {
	db *leveldb.DB
}

// NewInstanceDB returns the database which is located under given path.
func NewInstanceDB(path string) (*LevelDBInstanceDB, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &LevelDBInstanceDB{db: db}, nil
}

// marshal returns the byte slice from service.
func (d *LevelDBInstanceDB) marshal(i *instance.Instance) ([]byte, error) {
	return json.Marshal(i)
}

// unmarshal returns the service from byte slice.
func (d *LevelDBInstanceDB) unmarshal(hash hash.Hash, value []byte) (*instance.Instance, error) {
	var s instance.Instance
	if err := json.Unmarshal(value, &s); err != nil {
		return nil, fmt.Errorf("database: could not decode instance %q: %s", hash, err)
	}
	return &s, nil
}

// Exist check if instance with given hash exist.
func (d *LevelDBInstanceDB) Exist(hash hash.Hash) (bool, error) {
	return d.db.Has(hash, nil)
}

// Get retrives instance by instance hash.
func (d *LevelDBInstanceDB) Get(hash hash.Hash) (*instance.Instance, error) {
	b, err := d.db.Get(hash, nil)
	if err != nil {
		return nil, err
	}
	return d.unmarshal(hash, b)
}

// GetAll retrieves all instances.
func (d *LevelDBInstanceDB) GetAll() ([]*instance.Instance, error) {
	instances := []*instance.Instance{}
	iter := d.db.NewIterator(nil, nil)
	for iter.Next() {
		i, err := d.unmarshal(iter.Key(), iter.Value())
		if err != nil {
			iter.Release()
			return nil, err
		}
		instances = append(instances, i)
	}
	iter.Release()
	return instances, iter.Error()
}

// Save saves instance to database.
func (d *LevelDBInstanceDB) Save(i *instance.Instance) error {
	if i.Hash.IsZero() {
		return errSaveInstanceWithoutHash
	}

	// encode service
	b, err := d.marshal(i)
	if err != nil {
		return err
	}

	return d.db.Put(i.Hash, b, nil)
}

// Close closes database.
func (d *LevelDBInstanceDB) Close() error {
	return d.db.Close()
}

// Delete deletes an instance from database.
func (d *LevelDBInstanceDB) Delete(hash hash.Hash) error {
	return d.db.Delete(hash, nil)
}
