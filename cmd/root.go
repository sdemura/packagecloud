package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "packagecloud",
	Short:         "CLI for packagecloud.io",
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(ctx context.Context) error {
	return rootCmd.ExecuteContext(ctx)
}
