package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// All the configuration keys
const (
	APIServerAddress       = "Api.Server.Address"
	APIServerSocket        = "Api.Server.Socket"
	APIClientTarget        = "Api.Client.Target"
	APIServiceTargetPath   = "Api.Service.TargetPath"
	APIServiceTargetSocket = "Api.Service.TargetSocket"
	APIServiceSocketPath   = "Api.Service.SocketPath"
	ServicePathHost        = "Service.Path.Host"
	ServicePathDocker      = "Service.Path.Docker"
	MESGPath               = "MESG.Path"
	CoreImage              = "Core.Image"
)

func init() {
	configDir, _ := getConfigDirectory()

	viper.AddConfigPath(configDir)
	viper.SetConfigName(".mesg")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	viper.SetDefault(MESGPath, configDir)

	viper.SetDefault(APIServerAddress, ":50052")
	viper.SetDefault(APIServerSocket, "/mesg/server.sock")
	os.MkdirAll("/mesg", os.ModePerm)

	viper.SetDefault(APIClientTarget, viper.GetString(APIServerAddress))

	viper.SetDefault(APIServiceSocketPath, filepath.Join(viper.GetString(MESGPath), "server.sock"))
	viper.SetDefault(APIServiceTargetPath, "/mesg/server.sock")
	viper.SetDefault(APIServiceTargetSocket, "unix://"+viper.GetString(APIServiceTargetPath))

	viper.SetDefault(ServicePathHost, filepath.Join(viper.GetString(MESGPath), "services"))
	viper.SetDefault(ServicePathDocker, filepath.Join("/mesg", "services"))
	os.MkdirAll(viper.GetString(ServicePathDocker), os.ModePerm)

	viper.SetDefault(CoreImage, "mesg/core:latest")
}
