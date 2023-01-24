package cmd

import (
	"proton/internal/services"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the Fiber Server",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := services.Run();
		return err;
	},
}
