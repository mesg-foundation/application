package daemon

import (
	"fmt"

	"github.com/logrusorgru/aurora"
	"github.com/mesg-foundation/core/cmd/utils"
	"github.com/mesg-foundation/core/daemon"
	"github.com/spf13/cobra"
)

// Stop the daemon
var Stop = &cobra.Command{
	Use:               "stop",
	Short:             "Stop the daemon",
	Run:               stopHandler,
	DisableAutoGenTag: true,
}

func stopHandler(cmd *cobra.Command, args []string) {
	err := daemon.Stop()
	if err != nil {
		fmt.Println(aurora.Red(err))
		return
	}

	spinner := cmdUtils.StartSpinner(cmdUtils.SpinnerOptions{Text: "Stopping the daemon"})
	wait := daemon.WaitForStopped()
	err = <-wait
	spinner.Stop()
	if err != nil {
		fmt.Println(aurora.Red(err))
		return
	}

	fmt.Println(aurora.Green("Daemon stopped"))
}
