package utils

import (
	"os"

	"github.com/mesg-foundation/core/conversion"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

// GetOrAskAmount return the amount in MESG based on the flag or the user input
func GetOrAskAmount(cmd *cobra.Command, message string) (amount *conversion.Amount, err error) {
	value := cmd.Flag("amount").Value.String()
	if value == "" && survey.AskOne(&survey.Input{Message: message}, &value, nil) != nil {
		os.Exit(0)
	}
	amount = &conversion.Amount{}
	err = amount.FromString(value)
	return
}

// Payable mark a command as payable
func Payable(cmd *cobra.Command) {
	cmd.Flags().StringP("amount", "", "", "Amount of MESG")
}
