package main

import (
	"github.com/mesg-foundation/core/config"
	"github.com/mesg-foundation/core/database"
	"github.com/mesg-foundation/core/interface/grpc"
	"github.com/mesg-foundation/core/logger"
	"github.com/mesg-foundation/core/service"
	"github.com/mesg-foundation/core/version"
	"github.com/sirupsen/logrus"
)

func main() {
	c, err := config.Global()
	if err != nil {
		logrus.Fatalln(err)
	}

	db, err := database.NewServiceDB(c.Database.Path)
	if err != nil {
		logrus.Fatalln(err)
	}

	logger.Init(c.Log.Format, c.Log.Level)

	logrus.Println("Starting MESG Core", version.Version)

	for _, plugin := range c.Server.Plugins {
		if plugin.Path != "" {
			if _, err := service.NewFromPath(plugin.Path); err != nil {
				logrus.Fatalln(err)
			}
		}
	}

	tcpServer := &grpc.Server{
		Network:   "tcp",
		Address:   c.Server.Address,
		ServiceDB: db,
	}

	if err := tcpServer.Serve(); err != nil {
		logrus.Fatalln(err)
	}
	tcpServer.Close()
}
