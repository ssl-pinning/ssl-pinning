package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
	"github.com/ssl-pinning/core"
	logger "gopkg.in/slog-handler.v1"
)

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Generate admin token for the UI",
	Run: func(cmd *cobra.Command, args []string) {
		logger.SetGlobalLogger(logger.Options{
			Format: "text",
			Level:  "error",
			Pretty: true,
		})

		token, err := core.GenerateToken()
		if err != nil {
			slog.Error("failed to generate token", "error", err)
			os.Exit(1)
		}

		fmt.Println(token)
	},
}

func init() {
	rootCmd.AddCommand(tokenCmd)
}
