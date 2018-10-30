package api

import (
	"os"
	"testing"

	"github.com/mesg-foundation/core/container"
	"github.com/mesg-foundation/core/container/dockertest"
	"github.com/mesg-foundation/core/database"
	"github.com/stretchr/testify/require"
)

const testdbname = "db.test"

func newAPIAndDockerTest(t *testing.T) (*API, *dockertest.Testing, func()) {

	dt := dockertest.New()

	container, err := container.New(container.ClientOption(dt.Client()))
	require.NoError(t, err)

	db, err := database.NewServiceDB(testdbname)
	require.NoError(t, err)

	execDB, err := database.NewExecutionDB("execution" + testdbname)
	require.NoError(t, err)

	a, err := New(db, execDB, ContainerOption(container))
	require.Nil(t, err)

	closer := func() {
		require.NoError(t, db.Close())
		require.NoError(t, execDB.Close())
		require.NoError(t, os.RemoveAll(testdbname))
	}
	return a, dt, closer
}
