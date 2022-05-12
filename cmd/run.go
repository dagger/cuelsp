package cmd

import (
	"context"
	"fmt"

	"github.com/dagger/dlsp/logger"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var runCmd = &cobra.Command{
	Use:               "run",
	Short:             "Start the Dagger language server",
	PersistentPreRun:  func(*cobra.Command, []string) {},
	PersistentPostRun: func(*cobra.Command, []string) {},
	Args:              cobra.MaximumNArgs(2), // to verify
	PreRun: func(cmd *cobra.Command, args []string) {
		// Fix Viper bug for duplicate flags:
		// https://github.com/spf13/viper/issues/233
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			panic(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		lg := logger.New()
		ctx := lg.WithContext(cmd.Context())

		address := viper.GetString("address")
		protocol := viper.GetString("protocol")

		err := Run(ctx, protocol, address)
		if err != nil {
			lg.Fatal().Err(err).Msg("failed to run LSP")
		}
		fmt.Printf("LSP started at address: %s, with %s protocol\n", address, protocol)
	},
}

func Run(ctx context.Context, protocol string, address string) error {
	lg := log.Ctx(ctx)
	lg.Info().Msg("starting LSP with protocol " + protocol + " at address " + address)
	// server := serverpkg.NewServer(&tosca.Handler, toolName, verbose > 0)

	switch protocol {
	case "stdio":
		return fmt.Errorf("stdio-tata")
		// return server.RunStdio()

	case "tcp":
		return fmt.Errorf("tcp-tata")
		// return server.RunTCP(address)

	case "websocket":
		return fmt.Errorf("websocket-tata")
		// return server.RunWebSocket(address)

	case "nodejs":
		return fmt.Errorf("nodejs-tata")
		// return server.RunNodeJs()

	default:
		return fmt.Errorf("unsupported protocol: %s", protocol)
	}
}

func init() {
	runCmd.Flags().StringP("protocol", "p", "nodejs", "protocol (\"stdio\", \"tcp\", \"websocket\", or \"nodejs\"")
	runCmd.Flags().Uint16P("address", "a", 4389, "address (for \"tcp\" and \"websocket\"")
	if err := viper.BindPFlags(runCmd.Flags()); err != nil {
		panic(err)
	}
}
