package database

import (
	"os"
	"sync"
	"testing"

	"github.com/mesg-foundation/core/service"
	"github.com/stretchr/testify/require"
	"github.com/syndtr/goleveldb/leveldb"
)

const testdbname = "db.test"

func openServiceDB(t *testing.T) (*LevelDBServiceDB, func()) {
	deleteDBs(t)
	db, err := NewServiceDB(testdbname)
	require.NoError(t, err)
	return db, func() {
		require.NoError(t, db.Close())
		deleteDBs(t)
	}
}

func deleteDBs(t *testing.T) {
	require.NoError(t, os.RemoveAll(testdbname))
}

func TestServiceDBSave(t *testing.T) {
	db, closer := openServiceDB(t)
	defer closer()

	s1 := &service.Service{ID: "00", SID: "1", Name: "test-service"}
	require.NoError(t, db.Save(s1))

	// save same service. should replace
	require.NoError(t, db.Save(s1))
	ss, _ := db.All()
	require.Len(t, ss, 1)

	// different id, same sid. should replace s1.
	s2 := &service.Service{ID: "01", SID: "1", Name: "test-service"}
	require.NoError(t, db.Save(s2))
	_, err := db.Get(s1.ID)
	require.IsType(t, &ErrNotFound{}, err)

	// different id, different sid. should not replace anything.
	s3 := &service.Service{ID: "02", SID: "2", Name: "test-service"}
	require.NoError(t, db.Save(s3))
	ss, _ = db.All()
	require.Len(t, ss, 2)

	// test service without id
	s := &service.Service{Name: "test-service", SID: "SID"}
	require.EqualError(t, db.Save(s), errCannotSaveWithoutID.Error())

	// test service without sid.
	s = &service.Service{Name: "test-service", ID: "id"}
	require.EqualError(t, db.Save(s), errCannotSaveWithoutSID.Error())

	// test service where id has the same length as sid.
	s = &service.Service{Name: "test-service", ID: "sameLength", SID: "sameLength"}
	require.EqualError(t, db.Save(s), errSIDSameLen.Error())
}

func TestServiceDBGet(t *testing.T) {
	db, closer := openServiceDB(t)
	defer closer()

	want := &service.Service{ID: "00", SID: "2", Name: "test-service"}
	require.NoError(t, db.Save(want))
	defer db.Delete(want.ID)

	// id
	got, err := db.Get(want.ID)
	require.NoError(t, err)
	require.Equal(t, want, got)

	// sid.
	got, err = db.Get(want.SID)
	require.NoError(t, err)
	require.Equal(t, want, got)

	// test return err not found
	_, err = db.Get("3")
	require.Error(t, err)
	require.True(t, IsErrNotFound(err))
}

func TestServiceDBDelete(t *testing.T) {
	db, closer := openServiceDB(t)
	defer closer()

	// id
	s := &service.Service{ID: "00", SID: "2", Name: "test-service"}
	require.NoError(t, db.Save(s))
	require.NoError(t, db.Delete(s.ID))
	_, err := db.Get(s.ID)
	require.IsType(t, &ErrNotFound{}, err)
	_, err = db.Get(s.SID)
	require.IsType(t, &ErrNotFound{}, err)

	_, err = db.db.Get([]byte(idKeyPrefix+s.ID), nil)
	require.Equal(t, leveldb.ErrNotFound, err)
	_, err = db.db.Get([]byte(sidKeyPrefix+s.SID), nil)
	require.Equal(t, leveldb.ErrNotFound, err)

	// sid.
	s = &service.Service{ID: "11", SID: "3", Name: "test-service"}
	require.NoError(t, db.Save(s))
	require.NoError(t, db.Delete(s.SID))
	_, err = db.Get(s.SID)
	require.IsType(t, &ErrNotFound{}, err)
	_, err = db.Get(s.ID)
	require.IsType(t, &ErrNotFound{}, err)

	_, err = db.db.Get([]byte(idKeyPrefix+s.ID), nil)
	require.Equal(t, leveldb.ErrNotFound, err)
	_, err = db.db.Get([]byte(sidKeyPrefix+s.SID), nil)
	require.Equal(t, leveldb.ErrNotFound, err)
}

func TestServiceDBDeleteConcurrency(t *testing.T) {
	db, closer := openServiceDB(t)
	defer closer()

	s := &service.Service{ID: "00", SID: "2", Name: "test-service"}
	db.Save(s)

	var wg sync.WaitGroup
	errs := make([]error, 0)
	errsM := &sync.Mutex{}
	n := 10
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			if err := db.Delete(s.ID); err != nil {
				errsM.Lock()
				errs = append(errs, err)
				errsM.Unlock()
			}
		}()
	}

	wg.Wait()
	require.Len(t, errs, n-1)
	for i := 0; i < len(errs); i++ {
		require.IsType(t, &ErrNotFound{}, errs[i])
	}
}

func TestServiceDBAll(t *testing.T) {
	db, closer := openServiceDB(t)
	defer closer()

	s1 := &service.Service{ID: "00", SID: "SID1", Name: "test-service"}
	s2 := &service.Service{ID: "01", SID: "SID2", Name: "test-service"}

	require.NoError(t, db.Save(s1))
	require.NoError(t, db.Save(s2))
	defer db.Delete(s1.ID)
	defer db.Delete(s2.ID)

	services, err := db.All()
	require.NoError(t, err)
	require.Len(t, services, 2)
	require.Contains(t, services, s1)
	require.Contains(t, services, s2)
}

func TestServiceDBAllWithDecodeError(t *testing.T) {
	db, closer := openServiceDB(t)
	defer closer()

	id := "idtest"
	require.NoError(t, db.db.Put([]byte(id), []byte("oaiwdhhiodoihwaiohwa"), nil))
	defer db.db.Delete([]byte(id), nil)

	s1 := &service.Service{ID: "00", SID: "2", Name: "test-service"}
	require.NoError(t, db.Save(s1))
	defer db.Delete(s1.ID)

	services, err := db.All()
	require.NoError(t, err)
	require.Len(t, services, 1)
	require.Contains(t, services, s1)
}

func TestIsErrNotFound(t *testing.T) {
	require.True(t, IsErrNotFound(&ErrNotFound{}))
	require.False(t, IsErrNotFound(nil))
}
