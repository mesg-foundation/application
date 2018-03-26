package cmdAccount

import (
	"errors"

	"github.com/mesg-foundation/application/account"
	"github.com/mesg-foundation/application/cmd/utils"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

// Export an account into a json file
var Export = &cobra.Command{
	Use:   "export",
	Short: "Export an account",
	Long: `This method creates a file containing the information about your account.
The private key of your account is encrypted with the password you choose.

**Warning:** This method does **NOT** export your password. You have to manage your password yourself.

You can import the backup file on any other MESG Application with the [import method](account/import.md).`,
	Example: `mesg-cli account export
mesg-cli account export --account 0x000000000 --password QWERTY --new-password QWERTY --path ./account-export`,
	Run:               exportHandler,
	DisableAutoGenTag: true,
}

func exportHandler(cmd *cobra.Command, args []string) {
	path := cmd.Flag("path").Value.String()
	acc := cmdUtils.AccountFromFlagOrAsk(cmd, "Choose the account to export:")
	password := cmd.Flag("password").Value.String()
	if password == "" {
		survey.AskOne(&survey.Password{Message: "Type current password:"}, &password, nil)
	}
	newPassword := cmd.Flag("new-password").Value.String()
	if newPassword == "" {
		var passwordConfirmation string
		survey.AskOne(&survey.Password{Message: "Type new password for exportation:"}, &newPassword, nil)
		survey.AskOne(&survey.Password{Message: "Repeat password for exportation:"}, &passwordConfirmation, nil)
		if newPassword != passwordConfirmation {
			panic(errors.New("Passwords are different"))
		}
	}

	err := account.Export(acc, password, newPassword, path)
	if err != nil {
		panic(err)
	}

	// TODO: show confirmation with path
}

func init() {
	cmdUtils.Accountable(Export)
	Export.Flags().StringP("password", "", "", "Current password of the account to export")
	Export.Flags().StringP("new-password", "", "", "New password of the exported account")
	Export.Flags().StringP("path", "p", "./export", "Path of the file your account will be exported in")
}
