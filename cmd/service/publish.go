package cmdService

import (
	"fmt"
	"time"

	"github.com/logrusorgru/aurora"
	"github.com/mesg-foundation/application/cmd/utils"
	"github.com/mesg-foundation/application/service"
	"github.com/spf13/cobra"
)

// Publish a service to the marketplace
var Publish = &cobra.Command{
	Use:   "publish",
	Short: "Publish a service",
	Long: `Publish a service on the Network.

To get more information, see the [publish page from the documentation](https://docs.mesg.tech/service/develop/publish.html)`,
	Args: cobra.MinimumNArgs(1),
	Example: `mesg-cli service publish
mesg-cli service publish ./SERVICE_FOLDER --account ACCOUNT --confirm`,
	Run:               deployHandler,
	DisableAutoGenTag: true,
}

func deployHandler(cmd *cobra.Command, args []string) {
	account := cmdUtils.AccountFromFlagOrAsk(cmd, "Select an account:")
	if !cmdUtils.Confirm(cmd, "Are you sure?") {
		return
	}
	service, err := service.ImportFromFile(args[0])
	if err != nil {
		fmt.Println(aurora.Red(err))
		return
	}
	s := cmdUtils.StartSpinner(cmdUtils.SpinnerOptions{Text: "Deployment of " + service.Name + " in progress..."})
	time.Sleep(2 * time.Second)
	s.Stop()
	// TODO deploy the service
	fmt.Println("Service deployed with success with account: ", account)
}

func init() {
	cmdUtils.Confirmable(Publish)
	cmdUtils.Accountable(Publish)
}
