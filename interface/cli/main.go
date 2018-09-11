package main

import (
	"fmt"
	"os"

	"github.com/mesg-foundation/core/commands"
	"github.com/mesg-foundation/core/commands/provider"
	"github.com/mesg-foundation/core/config"
	"github.com/mesg-foundation/core/interface/grpc/core"
	"github.com/mesg-foundation/core/utils/clierrors"
	"github.com/mesg-foundation/core/version"
	"google.golang.org/grpc"
)

func init() {
}

func main() {
	cfg, err := config.Global()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	connection, err := grpc.Dial(cfg.Client.Address, grpc.WithInsecure())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	p := provider.New(core.NewCoreClient(connection))
	cmd := commands.Build(p)
	cmd.Version = version.Version
	cmd.Short = cmd.Short + " " + version.Version

	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, clierrors.ErrorMessage(err))
		os.Exit(1)
	}
}
