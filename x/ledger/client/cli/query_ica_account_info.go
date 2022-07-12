package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/ledger/types"
)

var _ = strconv.Itoa(0)

func CmdIcaAccountInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ica-account-info [owner] [ctrl-connection-id]",
		Short: "Query ica account info",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqOwner := args[0]
			reqCtrlConnectionId := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryIcaAccountInfoRequest{
				Owner:            reqOwner,
				CtrlConnectionId: reqCtrlConnectionId,
			}

			res, err := queryClient.IcaAccountInfo(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}