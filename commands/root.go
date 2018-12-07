package commands

import (
	"github.com/mesg-foundation/core/utils/pretty"
	"github.com/spf13/cobra"
)

type rootCmd struct {
	baseCmd

	noColor   bool
	noSpinner bool
}

func newRootCmd(e Executor) *rootCmd {
	c := &rootCmd{}
	c.cmd = newCommand(&cobra.Command{
		Use:              "mesg-core",
		Short:            "MESG Core",
		PersistentPreRun: c.persistentPreRun,
		SilenceUsage:     true,
		SilenceErrors:    true,
	})
	c.cmd.PersistentFlags().BoolVar(&c.noColor, "no-color", c.noColor, "disable colorized output")
	c.cmd.PersistentFlags().BoolVar(&c.noSpinner, "no-spinner", c.noSpinner, "disable spinners")

	c.cmd.AddCommand(
		newStartCmd(e).cmd,
		newStatusCmd(e).cmd,
		newStopCmd(e).cmd,
		newLogsCmd(e).cmd,
		newRootServiceCmd(e).cmd,
		// Workflow system is disable for v0.5. Enable it when ready.
		// newRootWorkflowCmd(e).cmd,
	)
	return c
}

func (c *rootCmd) persistentPreRun(cmd *cobra.Command, args []string) {
	if c.noColor {
		pretty.DisableColor()
	}
	if c.noSpinner {
		pretty.DisableSpinner()
	}
}
