package main

import (
	"os"
	"path"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/lcd"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	"github.com/cosmos/cosmos-sdk/x/genaccounts"
	genaccscmd "github.com/cosmos/cosmos-sdk/x/genaccounts/client/cli"
	genutilcmd "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/mesg-foundation/engine/codec"
	"github.com/mesg-foundation/engine/cosmos"
	"github.com/mesg-foundation/engine/logger"
	enginesdk "github.com/mesg-foundation/engine/sdk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	amino "github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/libs/cli"
	db "github.com/tendermint/tm-db"
)

const (
	cosmosPath = "cosmos"
)

var (
	defaultCLIHome  = os.ExpandEnv("$HOME/.mesgcli")
	defaultNodeHome = os.ExpandEnv("$HOME/.mesg") // TODO: should use config package
	app             *cosmos.App
)

func main() {
	cobra.EnableCommandSorting = false

	// init app factory
	db, err := db.NewGoLevelDB("app", filepath.Join(defaultCLIHome, cosmosPath))
	if err != nil {
		panic(err)
	}
	appFactory := cosmos.NewAppFactory(logger.TendermintLogger(), db)

	// register the backend modules to the app factory.
	enginesdk.NewBackend(appFactory)

	// init cosmos app
	app, err = cosmos.NewApp(appFactory)
	if err != nil {
		panic(err)
	}
	cdc := codec.Codec

	// Read in the configuration file for the sdk
	config := sdk.GetConfig()
	// TODO: do the following in the Engine and change it to MESG!
	config.SetBech32PrefixForAccount(sdk.Bech32PrefixAccAddr, sdk.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(sdk.Bech32PrefixValAddr, sdk.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(sdk.Bech32PrefixConsAddr, sdk.Bech32PrefixConsPub)
	config.Seal()

	rootCmd := &cobra.Command{
		Use:   "cli",
		Short: "Cosmos Client",
	}

	// Add --chain-id to persistent flags and mark it required
	rootCmd.PersistentFlags().String(client.FlagChainID, "", "Chain ID of tendermint node")
	rootCmd.PersistentPreRunE = func(_ *cobra.Command, _ []string) error {
		return initConfig(rootCmd)
	}

	// Construct Root Command
	rootCmd.AddCommand(
		rpc.StatusCommand(),
		client.ConfigCmd(defaultCLIHome),
		queryCmd(cdc),
		txCmd(cdc),
		genesisCmd(cdc),
		lcd.ServeCommand(cdc, registerRoutes),
		keys.Commands(),
		version.Cmd,
		client.NewCompletionCmd(rootCmd, true),
	)

	executor := cli.PrepareMainCmd(rootCmd, "MESG", defaultCLIHome)
	err = executor.Execute()
	if err != nil {
		panic(err)
	}
}

func registerRoutes(rs *lcd.RestServer) {
	client.RegisterRoutes(rs.CliCtx, rs.Mux)
	authrest.RegisterTxRoutes(rs.CliCtx, rs.Mux)
	app.BasicManager().RegisterRESTRoutes(rs.CliCtx, rs.Mux)
}

func queryCmd(cdc *amino.Codec) *cobra.Command {
	queryCmd := &cobra.Command{
		Use:     "query",
		Aliases: []string{"q"},
		Short:   "Querying subcommands",
	}

	queryCmd.AddCommand(
		rpc.ValidatorCommand(cdc),
		rpc.BlockCommand(),
		rpc.StatusCommand(),
		authcmd.QueryTxsByEventsCmd(cdc),
		authcmd.QueryTxCmd(cdc),
	)

	// add modules' query commands
	app.BasicManager().AddQueryCommands(queryCmd, cdc)

	return queryCmd
}

func txCmd(cdc *amino.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:   "tx",
		Short: "Transactions subcommands",
	}

	txCmd.AddCommand(
		authcmd.GetBroadcastCommand(cdc),
		authcmd.GetEncodeCommand(cdc),
	)

	// add modules' tx commands
	app.BasicManager().AddTxCommands(txCmd, cdc)

	return txCmd
}

func genesisCmd(cdc *amino.Codec) *cobra.Command {
	ctx := server.NewDefaultContext()

	genCmd := &cobra.Command{
		Use:               "genesis",
		Short:             "Genesis subcommands",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}

	genCmd.AddCommand(
		genutilcmd.InitCmd(ctx, cdc, app.BasicManager(), defaultNodeHome),
		genutilcmd.CollectGenTxsCmd(ctx, cdc, genaccounts.AppModuleBasic{}, defaultNodeHome),
		genutilcmd.GenTxCmd(
			ctx, cdc, app.BasicManager(), staking.AppModuleBasic{},
			genaccounts.AppModuleBasic{}, defaultNodeHome, defaultCLIHome,
		),
		genutilcmd.ValidateGenesisCmd(ctx, cdc, app.BasicManager()),
		// AddGenesisAccountCmd allows users to add accounts to the genesis file
		genaccscmd.AddGenesisAccountCmd(ctx, cdc, defaultNodeHome, defaultCLIHome),
	)

	return genCmd
}

func initConfig(cmd *cobra.Command) error {
	home, err := cmd.PersistentFlags().GetString(cli.HomeFlag)
	if err != nil {
		return err
	}

	cfgFile := path.Join(home, "config", "config.toml")
	if _, err := os.Stat(cfgFile); err == nil {
		viper.SetConfigFile(cfgFile)

		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}
	if err := viper.BindPFlag(client.FlagChainID, cmd.PersistentFlags().Lookup(client.FlagChainID)); err != nil {
		return err
	}
	if err := viper.BindPFlag(cli.EncodingFlag, cmd.PersistentFlags().Lookup(cli.EncodingFlag)); err != nil {
		return err
	}
	return viper.BindPFlag(cli.OutputFlag, cmd.PersistentFlags().Lookup(cli.OutputFlag))
}

/*
List of commands that could be used in this cli

distribution/client/cli/tx.go
func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
func GetCmdWithdrawRewards(cdc *codec.Codec) *cobra.Command {
func GetCmdWithdrawAllRewards(cdc *codec.Codec, queryRoute string) *cobra.Command {
func GetCmdSetWithdrawAddr(cdc *codec.Codec) *cobra.Command {
func GetCmdSubmitProposal(cdc *codec.Codec) *cobra.Command {

distribution/client/cli/query.go
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryParams(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryValidatorOutstandingRewards(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryValidatorCommission(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryValidatorSlashes(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryDelegatorRewards(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryCommunityPool(queryRoute string, cdc *codec.Codec) *cobra.Command {

gov/client/cli/query.go
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryProposal(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryProposals(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryVote(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryVotes(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryDeposit(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryDeposits(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryTally(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryParams(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryParam(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryProposer(queryRoute string, cdc *codec.Codec) *cobra.Command {

gov/client/cli/tx.go
func GetTxCmd(storeKey string, cdc *codec.Codec, pcmds []*cobra.Command) *cobra.Command {
func GetCmdSubmitProposal(cdc *codec.Codec) *cobra.Command {
func GetCmdDeposit(cdc *codec.Codec) *cobra.Command {
func GetCmdVote(cdc *codec.Codec) *cobra.Command {

bank/client/cli/tx.go
func GetTxCmd(cdc *codec.Codec) *cobra.Command {

params/client/cli/tx.go
func GetCmdSubmitProposal(cdc *codec.Codec) *cobra.Command {

auth/client/cli/query.go
func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
func GetAccountCmd(cdc *codec.Codec) *cobra.Command {

staking/client/cli/query.go
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryValidator(storeName string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryValidators(storeName string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryValidatorUnbondingDelegations(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryValidatorRedelegations(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryDelegation(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryDelegations(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryValidatorDelegations(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryUnbondingDelegation(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryUnbondingDelegations(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryRedelegation(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryRedelegations(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryPool(storeName string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryParams(storeName string, cdc *codec.Codec) *cobra.Command {

staking/client/cli/tx.go
func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
func GetCmdCreateValidator(cdc *codec.Codec) *cobra.Command {
func GetCmdEditValidator(cdc *codec.Codec) *cobra.Command {
func GetCmdDelegate(cdc *codec.Codec) *cobra.Command {
func GetCmdRedelegate(storeName string, cdc *codec.Codec) *cobra.Command {
func GetCmdUnbond(storeName string, cdc *codec.Codec) *cobra.Command {

crisis/client/cli/tx.go
func GetCmdInvariantBroken(cdc *codec.Codec) *cobra.Command {
func GetTxCmd(cdc *codec.Codec) *cobra.Command {

supply/client/cli/query.go
func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
func GetCmdQueryTotalSupply(cdc *codec.Codec) *cobra.Command {

slashing/client/cli/query.go
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
func GetCmdQuerySigningInfo(storeName string, cdc *codec.Codec) *cobra.Command {
func GetCmdQueryParams(cdc *codec.Codec) *cobra.Command {

slashing/client/cli/tx.go
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
func GetCmdUnjail(cdc *codec.Codec) *cobra.Command {

mint/client/cli/query.go
func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
func GetCmdQueryParams(cdc *codec.Codec) *cobra.Command {
func GetCmdQueryInflation(cdc *codec.Codec) *cobra.Command {
func GetCmdQueryAnnualProvisions(cdc *codec.Codec) *cobra.Command {

auth/client/cli/tx.go
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
*/
