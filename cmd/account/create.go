package account

import (
	"errors"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/logrusorgru/aurora"
	"github.com/mesg-foundation/core/account"
	"github.com/mesg-foundation/core/cmd/utils"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

// Create a new account
var Create = &cobra.Command{
	Use:   "create",
	Short: "Create a new account",
	Long: `This method creates a new account secured by a password. We strongly advise to use long randomized password.

**Warning:** Backup your password in a safe place. You will not be able to use the account if you lost the password.

You should also [export your account](mesg-core_account_export.md) to a safe place to prevent losing access to your workflows, services and tokens.`,
	Example: `mesg-core account create
mesg-core account create --password PASSWORD`,
	Run:               createHandler,
	DisableAutoGenTag: true,
}

func createHandler(cmd *cobra.Command, args []string) {
	password := cmd.Flag("password").Value.String()
	if password == "" {
		fmt.Printf("%s\n", aurora.Red("WARNING: Backup your password in a safe place. You will not be able to use the account if you lost the password.").Bold())
		var passwordConfirmation string
		if survey.AskOne(&survey.Password{Message: "Set a password:"}, &password, nil) != nil {
			os.Exit(0)
		}
		if survey.AskOne(&survey.Password{Message: "Repeat password:"}, &passwordConfirmation, nil) != nil {
			os.Exit(0)
		}
		if password != passwordConfirmation {
			panic(errors.New("Passwords are different"))
		}
	}

	s := utils.StartSpinner(utils.SpinnerOptions{Text: "Generating account..."})
	acc, err := account.Generate(password)
	if err != nil {
		panic(err)
	}
	s.Stop()
	fmt.Printf("%s\n", aurora.Green("Account created with success").Bold())

	displaySummary(acc)
}

func displaySummary(acc accounts.Account) {
	fmt.Println("Here is the detail of your account:")
	fmt.Printf("Account public address: %s\n", aurora.Green(acc.Address.String()).Bold())
}

func init() {
	Create.Flags().StringP("password", "p", "", "Password of the account")
}
