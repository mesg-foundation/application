package main

import (
	"github.com/mesg-foundation/core/config"
	"github.com/mesg-foundation/core/interface/grpc"
	"github.com/mesg-foundation/core/logger"
	"github.com/mesg-foundation/core/version"
	"github.com/sirupsen/logrus"
)

func main() {
	c, err := config.Global()
	if err != nil {
		logrus.Fatalln(err)
	}

	logger.Init(c.Log.Format, c.Log.Level)

	logrus.Println("Starting MESG Core", version.Version)

	tcpServer := &grpc.Server{
		Network: "tcp",
		Address: c.Server.Address,
	}

	if err := tcpServer.Serve(); err != nil {
		logrus.Fatalln(err)
	}
	tcpServer.Close()
}
