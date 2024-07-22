package cmd

import (
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds the application or services",
	Long:  `This command is used to build the application or services.`,
	// This command will act as a parent for subcommands and does not necessarily need to implement functionality in Run
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
