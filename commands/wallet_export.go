package commands

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

type walletExportCmd struct {
	baseWalletCmd

	address string

	e WalletExecutor
}

func newWalletExportCmd(e WalletExecutor) *walletExportCmd {
	c := &walletExportCmd{e: e}
	c.cmd = newCommand(&cobra.Command{
		Use:     "export",
		Short:   "export an account",
		Long:    "export an existing account in order to backup it and import it in an other wallet",
		Example: "mesg-core wallet export 0x0000000...00001",
		Args:    cobra.ExactArgs(1),
		PreRunE: c.preRunE,
		RunE:    c.runE,
	})
	c.setupFlags()
	return c
}

func (c *walletExportCmd) preRunE(cmd *cobra.Command, args []string) error {
	if err := c.askPassphrase(); err != nil {
		return err
	}
	// TODO: if no address provided, the cli should ask to select one.
	c.address = args[0]
	return nil
}

func (c *walletExportCmd) runE(cmd *cobra.Command, args []string) error {
	account, err := c.e.Export(c.address, c.passphrase)
	if err != nil {
		return err
	}
	b, err := json.Marshal(account)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}
