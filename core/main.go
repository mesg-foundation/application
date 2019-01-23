package main

import (
	"log"
	"path/filepath"

	"github.com/mesg-foundation/core/api"
	"github.com/mesg-foundation/core/config"
	"github.com/mesg-foundation/core/database"
	"github.com/mesg-foundation/core/interface/grpc"
	"github.com/mesg-foundation/core/logger"
	"github.com/mesg-foundation/core/systemservices"
	"github.com/mesg-foundation/core/systemservices/deployer"
	"github.com/mesg-foundation/core/version"
	"github.com/mesg-foundation/core/x/xsignal"
	"github.com/sirupsen/logrus"
)

func initGRPCServer(c *config.Config) (*grpc.Server, error) {
	// init services db.
	db, err := database.NewServiceDB(filepath.Join(c.Core.Path, c.Core.Database.ServiceRelativePath))
	if err != nil {
		return nil, err
	}

	// init execution db.
	execDB, err := database.NewExecutionDB(filepath.Join(c.Core.Path, c.Core.Database.ExecutionRelativePath))
	if err != nil {
		return nil, err
	}

	// init api.
	ss := systemservices.New()
	a, err := api.New(db, execDB, ss)
	if err != nil {
		return nil, err
	}

	// init system services.
	systemServicesPath := filepath.Join(c.Core.Path, c.SystemServices.RelativePath)
	d := deployer.New(a, systemServicesPath, ss)
	if err := d.Deploy(systemservices.SystemServicesList); err != nil {
		return nil, err
	}

	return grpc.New(c.Server.Address, a, ss), nil
}

func main() {
	// init configs.
	c, err := config.Global()
	if err != nil {
		log.Fatal(err)
	}

	// init logger.
	logger.Init(c.Log.Format, c.Log.Level, c.Log.ForceColors)

	// init gRPC server.
	server, err := initGRPCServer(c)
	if err != nil {
		logrus.Fatalln(err)
	}

	logrus.Infof("starting MESG Core version %s", version.Version)

	go func() {
		if err := server.Serve(); err != nil {
			logrus.Fatalln(err)
		}
	}()

	<-xsignal.WaitForInterrupt()
	server.Close()
}
