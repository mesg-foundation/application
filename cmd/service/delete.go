package cmdService

import (
	"context"
	"fmt"

	"github.com/logrusorgru/aurora"
	"github.com/mesg-foundation/core/api/core"
	"github.com/mesg-foundation/core/cmd/utils"
	"github.com/mesg-foundation/core/database/services"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

// Delete a service to the marketplace
var Delete = &cobra.Command{
	Use:   "delete",
	Short: "Delete one or many services",
	Example: `mesg-core service delete SERVICE_ID
mesg-core service delete SERVICE_ID_1 SERVICE_ID_2`,
	Run:               deleteHandler,
	DisableAutoGenTag: true,
}

func init() {
	Delete.Flags().BoolP("all", "", false, "Delete all services")
}

func deleteHandler(cmd *cobra.Command, args []string) {
	if cmd.Flag("all").Value.String() == "true" {
		var confirmed bool
		if survey.AskOne(&survey.Confirm{Message: "Are you sure to delete all services?"}, &confirmed, nil) != nil {
			return
		}
		if confirmed == false {
			return
		}
		fmt.Println("Deleting all services...")
		services, err := services.All() // TODO: this should use the API
		handleError(err)
		if len(services) == 0 {
			fmt.Println("All services are already deleted")
			return
		}
		for _, service := range services {
			args = append(args, service.Hash())
		}
	}
	if len(args) == 0 {
		fmt.Println(aurora.Red("No provided service ID. See help with flag --help"))
	}
	for _, arg := range args {
		var err error
		cmdUtils.ShowSpinnerForFunc(cmdUtils.SpinnerOptions{Text: "Deleting service " + arg + "..."}, func() {
			_, err = cli.DeleteService(context.Background(), &core.DeleteServiceRequest{
				ServiceID: arg,
			})
		})
		cmdUtils.HandleError(err)
		fmt.Println("Service", arg, "deleted")
	}
}
