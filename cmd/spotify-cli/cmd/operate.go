package cmd

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/zext/spotify-cli/app"
	"github.com/zext/spotify-cli/client/dbus"
)

func runOperate(cmd *cobra.Command, args []string) error {
	var operator app.Operator

	switch client {
	case "dbus":
		ope, err := dbus.NewDBusOperator()
		if err != nil {
			return fmt.Errorf("failed to initialize DBus operator: %w", err)
		}

		operator = ope

	case "web":
		return fmt.Errorf("not implemented")
	}

	service := app.OperatorService{
		Operator: operator,
	}

	operation := args[0]
	if err := service.Operate(app.OperationMap[operation]); err != nil {
		return fmt.Errorf("failed to operate: %w", err)
	}

	return nil
}

var operateCmd = &cobra.Command{
	Use:       "operate",
	Short:     "Operate tracks",
	ValidArgs: lo.Keys[string, app.Operation](app.OperationMap),
	Args:      cobra.ExactValidArgs(1),
	RunE:      runOperate,
}

func init() {
	rootCmd.AddCommand(operateCmd)
}
