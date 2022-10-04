package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const appVersion = "0.0.1"
const appName = "google-image-downloader"

var rootCmd = &cobra.Command{
	Use:   appName + " [command] \"query-string\" [flags]",
	Short: appName + " is search google and download image files.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Usage: %s [command] [flags]\n\nFor more information, use help.\n", appName)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(version)
	rootCmd.AddCommand(get)
}
