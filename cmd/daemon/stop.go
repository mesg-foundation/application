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
	var err error
	cmdUtils.ShowSpinnerForFunc(cmdUtils.SpinnerOptions{Text: "Stopping daemon..."}, func() {
		err = daemon.Stop()
	})
	if err != nil {
		fmt.Println(aurora.Red(err))
		return
	}
	fmt.Println(aurora.Green("Daemon stopped"))
}
