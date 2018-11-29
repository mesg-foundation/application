package commands

import (
	"errors"
	"fmt"
	"sync"

	"github.com/mesg-foundation/core/commands/provider"
	"github.com/mesg-foundation/core/utils/pretty"
	"github.com/mesg-foundation/core/x/xerrors"
	"github.com/spf13/cobra"
)

type serviceDeployCmd struct {
	baseCmd

	path string

	e ServiceExecutor
}

func newServiceDeployCmd(e ServiceExecutor) *serviceDeployCmd {
	c := &serviceDeployCmd{e: e}
	c.cmd = newCommand(&cobra.Command{
		Use:   "deploy",
		Short: "Deploy a service",
		Long: `Deploy a service.

To get more information, see the [deploy page from the documentation](https://docs.mesg.com/guide/service/deploy-a-service.html)`,
		Example: `mesg-core service deploy PATH_TO_SERVICE`,
		PreRunE: c.preRunE,
		RunE:    c.runE,
		Args:    cobra.MaximumNArgs(1),
	})
	return c
}

func (c *serviceDeployCmd) preRunE(cmd *cobra.Command, args []string) error {
	c.path = getFirstOrDefault(args)
	return nil
}

func (c *serviceDeployCmd) runE(cmd *cobra.Command, args []string) error {
	var (
		statuses = make(chan provider.DeployStatus)
		wg       sync.WaitGroup
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		printDeployStatuses(statuses)
	}()

	id, validationError, err := c.e.ServiceDeploy(c.path, statuses)
	wg.Wait()

	pretty.DestroySpinner()
	if err != nil {
		return err
	}
	if validationError != nil {
		return xerrors.Errors{
			validationError,
			errors.New("to get more information, run: mesg-core service validate"),
		}
	}
	fmt.Printf("%s Service deployed with ID: %v\n", pretty.SuccessSign, pretty.Success(id))
	fmt.Printf("To start it, run the command:\n\tmesg-core service start %s\n", id)
	return nil
}

func printDeployStatuses(statuses chan provider.DeployStatus) {
	for status := range statuses {
		switch status.Type {
		case provider.Running:
			pretty.UseSpinner(status.Message)
		default:
			var sign string
			switch status.Type {
			case provider.DonePositive:
				sign = pretty.SuccessSign
			case provider.DoneNegative:
				sign = pretty.FailSign
			}
			pretty.DestroySpinner()
			fmt.Printf("%s %s\n", sign, status.Message)
		}
	}
}
