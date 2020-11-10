package cli

import (
	"io/ioutil"
	"os"
	"sort"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	atypes "github.com/ovrclk/akash/types"
	"github.com/ovrclk/akash/x/audit/types"
	ptypes "github.com/ovrclk/akash/x/provider/types"
)

// GetTxCmd returns the transaction commands for audit module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Audit transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		cmdAttributes(),
	)

	return cmd
}

func cmdAttributes() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "attr",
		Short: "Manage provider attributes",
	}

	cmd.AddCommand(
		cmdCreateProviderAttributes(),
		cmdDeleteProviderAttributes(),
	)

	return cmd
}

func cmdCreateProviderAttributes() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create/update provider attributes",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.Errorf("invalid arguments")
			}

			if ((len(args) - 1) % 2) != 0 {
				return errors.Errorf("attributes must be provided as pairs")
			}

			providerAddress, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			attr, err := readAttributes(args[1:])
			if err != nil {
				return err
			}

			if len(attr) == 0 {
				return errors.Errorf("no attributes provided")
			}

			clientCtx, err := client.ReadTxCommandFlags(client.GetClientContextFromCmd(cmd), cmd.Flags())
			if err != nil {
				return err
			}

			msg := &types.MsgSignProviderAttributes{
				Validator:  clientCtx.GetFromAddress().String(),
				Owner:      providerAddress.String(),
				Attributes: attr,
			}

			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	setCmdProviderFlags(cmd)

	return cmd
}

func cmdDeleteProviderAttributes() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete provider attributes",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.Errorf("invalid arguments")
			}

			if ((len(args) - 1) % 2) != 0 {
				return errors.Errorf("attributes must be provided as pairs")
			}

			providerAddress, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			attr, err := readAttributes(args[1:])
			if err != nil {
				return err
			}

			clientCtx, err := client.ReadTxCommandFlags(client.GetClientContextFromCmd(cmd), cmd.Flags())
			if err != nil {
				return err
			}

			msg := &types.MsgDeleteProviderAttributes{
				Validator:  clientCtx.GetFromAddress().String(),
				Owner:      providerAddress.String(),
				Attributes: attr,
			}

			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	setCmdProviderFlags(cmd)

	return cmd
}

func setCmdProviderFlags(cmd *cobra.Command) {
	flags.AddTxFlagsToCmd(cmd)

	if err := cmd.MarkFlagRequired(flags.FlagFrom); err != nil {
		panic(err.Error())
	}
}

// readAttributes try read attributes from both cobra arguments and stdin
// len of the args must be even
// read from stdin uses trick to check if it's file descriptor is a pipe
// which happens when some data is piped for example cat attr.yaml | akash ...
func readAttributes(args []string) (ptypes.Attributes, error) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}

	var attr ptypes.Attributes

	if fi.Mode()&os.ModeNamedPipe != 0 {
		// pipe has some data
		var bytes []byte
		if bytes, err = ioutil.ReadAll(os.Stdin); err != nil {
			return nil, err
		}

		stdinAttr := make(map[string]string)

		if err = yaml.Unmarshal(bytes, &attr); err != nil {
			return nil, err
		}

		for k, v := range stdinAttr {
			attr = append(attr, atypes.Attribute{
				Key:   k,
				Value: v,
			})
		}
	}

	for i := 0; i < len(args); i += 2 {
		attr = append(attr, atypes.Attribute{
			Key:   args[i],
			Value: args[i+1],
		})
	}

	sort.SliceStable(attr, func(i, j int) bool {
		return attr[i].Key < attr[j].Value
	})

	if checkAttributeDuplicates(attr) {
		return nil, errors.Errorf("supplied attributes with duplicate keys")
	}

	return attr, nil
}

func checkAttributeDuplicates(attr ptypes.Attributes) bool {
	keys := make(map[string]bool)

	for _, entry := range attr {
		if _, value := keys[entry.Key]; !value {
			keys[entry.Key] = true
		} else {
			return true
		}
	}
	return false
}
