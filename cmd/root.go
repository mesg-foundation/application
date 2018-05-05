package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mesg-foundation/core/cmd/utils"
	"github.com/mesg-foundation/core/config"

	"github.com/mesg-foundation/core/api"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:               "mesg-cli",
	Short:             "MESG CLI",
	Run:               rootHandler,
	DisableAutoGenTag: true,
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".mesg" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".mesg")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func rootHandler(cmd *cobra.Command, args []string) {
	log.Println("Starting MESG daemon")
	serverTCP := api.Server{
		Network: "tcp",
		Address: viper.GetString(config.APIServerAddress),
	}
	serverSocket := api.Server{
		Network: "unix",
		Address: viper.GetString(config.APIServerSocket),
	}
	go func() {
		err := serverTCP.Serve()
		defer serverTCP.Stop()
		if err != nil {
			log.Panicln(err)
		}
	}()
	go func() {
		err := serverSocket.Serve()
		defer serverSocket.Stop()
		if err != nil {
			log.Panicln(err)
		}
	}()
	<-cmdUtils.WaitForCancel()
}
