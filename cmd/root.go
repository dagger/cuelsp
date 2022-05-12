package cmd

import (
	// "github.com/rebase-dagger-doc/cmd/dagger/logger"
	"github.com/dagger/dlsp/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "dagger-lsp",
	Short: "Dagger LSP",
}

func init() {
	rootCmd.PersistentFlags().String("log-format", "auto", "Log format (auto, plain, tty, json)")
	rootCmd.PersistentFlags().StringP("log-level", "l", "info", "Log level")
	rootCmd.AddCommand(
		runCmd,
	)

	if err := viper.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		panic(err)
	}
}

func Execute() {
	lg := logger.New()

	if err := rootCmd.Execute(); err != nil {
		lg.Fatal().Err(err).Msg("failed to execute command")
	}
}
