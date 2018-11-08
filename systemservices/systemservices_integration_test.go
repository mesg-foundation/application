// +build integration

package systemservices

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mesg-foundation/core/api"
	"github.com/mesg-foundation/core/config"
	"github.com/mesg-foundation/core/database"
)

func TestNew(t *testing.T) {
	conf, err := config.Global()
	require.NoError(t, err)

	var (
		systemServicesPath = filepath.Join(conf.Core.Path, conf.SystemServices.RelativePath)
		databasePath       = filepath.Join(conf.Core.Path, conf.Core.Database.ServiceRelativePath)
		execDatabasePath   = filepath.Join(conf.Core.Path, conf.Core.Database.ExecutionRelativePath)
	)

	serviceDB, err := database.NewServiceDB(databasePath)
	require.NoError(t, err)

	execDB, err := database.NewExecutionDB(execDatabasePath)
	require.NoError(t, err)

	a, err := api.New(serviceDB, execDB)
	require.NoError(t, err)

	s, err := New(a, systemServicesPath)
	require.NoError(t, err)
	require.NotZero(t, s)
}
