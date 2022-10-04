package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = &cobra.Command{
	Use:   appName + " version",
	Short: "Display version number.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s version %s\n", appName, appVersion)
	},
}
