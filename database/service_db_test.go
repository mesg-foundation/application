package database

import (
	"os"
	"testing"

	"github.com/mesg-foundation/core/service"
	"github.com/stretchr/testify/require"
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
	require.NoError(t, os.RemoveAll(testdbname+aliasesPathSuffix))
}

func TestDecodeError(t *testing.T) {
	db, closer := openServiceDB(t)
	defer closer()
	_, err := db.unmarshal("IDToTest", []byte("oaiwdhhiodoihwaiohwa"))
	require.IsType(t, &DecodeError{}, err)
}

func TestServiceDBSave(t *testing.T) {
	db, closer := openServiceDB(t)
	defer closer()

	s := &service.Service{ID: "1", Name: "test-service"}
	require.NoError(t, db.Save(s))
	defer db.Delete(s.ID)

	// test service without id
	s = &service.Service{Name: "test-service"}
	require.Error(t, db.Save(s))
}

func TestServiceDBSaveWithAlias(t *testing.T) {
	db, closer := openServiceDB(t)
	defer closer()

	s := &service.Service{ID: "1", Alias: "2", Name: "test-service"}
	require.NoError(t, db.Save(s))
	defer db.Delete(s.ID)

	// try saving a service with the same alias.
	s = &service.Service{ID: "3", Alias: "2", Name: "test-service"}
	require.Error(t, db.Save(s))
}

func TestServiceDBGet(t *testing.T) {
	db, closer := openServiceDB(t)
	defer closer()

	want := &service.Service{ID: "1", Name: "test-service"}
	require.NoError(t, db.Save(want))
	defer db.Delete(want.ID)

	got, err := db.Get(want.ID)
	require.NoError(t, err)
	require.Equal(t, want, got)

	// test return err not found
	_, err = db.Get("2")
	require.Error(t, err)
	require.True(t, IsErrNotFound(err))
}

func TestServiceDBDelete(t *testing.T) {
	db, closer := openServiceDB(t)
	defer closer()

	s := &service.Service{ID: "1", Name: "test-service"}
	require.NoError(t, db.Save(s))
	require.NoError(t, db.Delete(s.ID))
	_, err := db.Get(s.ID)
	require.True(t, IsErrNotFound(err))
}

func TestServiceDBAll(t *testing.T) {
	db, closer := openServiceDB(t)
	defer closer()

	s1 := &service.Service{ID: "1", Name: "test-service"}
	s2 := &service.Service{ID: "2", Name: "test-service"}

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

func TestIsErrNotFound(t *testing.T) {
	require.True(t, IsErrNotFound(&ErrNotFound{}))
	require.False(t, IsErrNotFound(nil))
}
