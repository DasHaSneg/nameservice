package cli

import (
	"fmt"
	"github.com/cosmos/sdk-application-tutorial/nameservice/x/nameservice/client/cli"
	"nameservice/x/nameservice/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	nameserviceQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Query commands for the nameservice module",
		RunE:                       client.ValidateCmd,
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}
	nameserviceQueryCmd.AddCommand(client.GetCommands(
		cli.GetCmdResolveName(storeKey, cdc),
		cli.GetCmdWhois(storeKey, cdc),
		cli.GetCmdNames(storeKey, cdc),
	)...)
	return nameserviceQueryCmd
}
