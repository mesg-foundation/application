package cmdAccount

import (
	"errors"
	"fmt"
	"time"

	"github.com/logrusorgru/aurora"
	"github.com/mesg-foundation/application/account"
	"github.com/mesg-foundation/application/cmd/utils"
	survey "gopkg.in/AlecAivazis/survey.v1"

	"github.com/spf13/cobra"
)

// Create run the create command for an account
var Create = &cobra.Command{
	Use:   "create",
	Short: "Create a new account",
	Long: `This method creates a new account secured by your password. We strongly advise to use long randomized password.
Warning: Backup your password in a safe place. You will not be able to use the account if you lost the password.
You should also export your account to a safe place to prevent losing access to your workflows, services and tokens. See account export method.`,
	Example:           "mesg-cli account create --name ACCOUNT_NAME --password ACCOUNT_PASSWORD",
	Run:               createHandler,
	DisableAutoGenTag: true,
}

func init() {
	Create.Flags().StringP("name", "n", "", "Name of the account")
	Create.Flags().StringP("password", "p", "", "Password for the account")
}

func createHandler(cmd *cobra.Command, args []string) {
	account := &account.Account{
		Password: cmd.Flag("password").Value.String(),
		Name:     cmd.Flag("name").Value.String(),
	}
	if err := checkPassword(account); err != nil {
		panic(err)
	}

	if err := checkName(account); err != nil {
		panic(err)
	}

	if err := generateAccount(account); err != nil {
		panic(err)
	}

	displaySummary(account)
}

func checkPassword(account *account.Account) (err error) {
	if account.Password != "" {
		return
	}
	var passwordConfirmation string
	survey.AskOne(&survey.Password{Message: "Please set a password ?"}, &account.Password, nil)
	survey.AskOne(&survey.Password{Message: "Repeat your password ?"}, &passwordConfirmation, nil)
	if account.Password != passwordConfirmation {
		err = errors.New("Password confirmation invalid")
		return
	}
	return
}

func checkName(account *account.Account) (err error) {
	if account.Name != "" {
		return
	}
	survey.AskOne(&survey.Input{Message: "Choose a name for this account"}, &account.Name, nil)
	return
}

func generateAccount(account *account.Account) (err error) {
	s := cmdUtils.StartSpinner(cmdUtils.SpinnerOptions{Text: "Generating secure key..."})
	time.Sleep(time.Second)
	s.Stop()

	err = account.Generate()
	return
}

func displaySummary(account *account.Account) {
	fmt.Println("Here is all the details of your account:")
	fmt.Println()
	fmt.Printf("Account name: %s\n", aurora.Green(account.Name).Bold())
	fmt.Printf("Account address: %s\n", aurora.Green(account.Address).Bold())
	fmt.Printf("Seed: %s\n", aurora.Green(account.Seed).Bold())
	fmt.Println()
	fmt.Printf("%s", aurora.Brown("Please make sure that you save those information and especially the following seed that might be needed to regenerate this address").Bold())
	fmt.Println()
}
