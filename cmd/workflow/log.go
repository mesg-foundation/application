package cmdWorkflow

import (
	"fmt"

	"github.com/mesg-foundation/core/cmd/utils"

	"github.com/spf13/cobra"
)

// Log workflow executions
var Log = &cobra.Command{
	Use:   "log WORKFLOW_ID",
	Short: "Log executions of a workflow",
	Args:  cobra.MinimumNArgs(1),
	Example: `mesg-cli workflow log WORKFLOW_ID
	mesg-cli workflow log WORKFLOW_ID --account ACCOUNT_ID
mesg-cli workflow log WORKFLOW_ID --execution EXECUTION_ID
mesg-cli workflow log WORKFLOW_ID --task TASK_ID
mesg-cli workflow log WORKFLOW_ID --from DATE --to DATE`,
	Run:               logHandler,
	DisableAutoGenTag: true,
}

func logHandler(cmd *cobra.Command, args []string) {
	account := cmdUtils.AccountFromFlagOrAsk(cmd, "Select an account:")

	execution := cmd.Flag("execution").Value.String()
	task := cmd.Flag("task").Value.String()
	from := cmd.Flag("from").Value.String()
	to := cmd.Flag("to").Value.String()

	fmt.Printf("Logs with account %s of workflow %s, execution: %s, task: %s, from: %s, to: %s", account, args[0], execution, task, from, to)
}

func init() {
	Log.Flags().StringP("execution", "e", "", "Log a specific execution")
	Log.Flags().StringP("task", "t", "", "Log a specific task")
	Log.Flags().StringP("from", "", "", "Log from date")
	Log.Flags().StringP("to", "", "", "Log to date")
	cmdUtils.Accountable(Log)
}
