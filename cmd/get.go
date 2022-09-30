package cmd

import (
	"errors"
	"fmt"
	"go_google_images/googleImageCrawler"
	"log"

	"github.com/spf13/cobra"
)

var get = &cobra.Command{
	Use:   "get",
	Short: "Download google images.",
	Long:  "search google and download image files.",
	Example: func() string {
		return fmt.Sprintf("  %s get \"IU\"\n", appName) +
			fmt.Sprintf("  %s get \"IU\" --target ./images\n", appName)
	}(),
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("enter the Query-String")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		targetPath, err := cmd.Flags().GetString("target")
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("search %s download to %s\n", args[0], targetPath)
		googleImageCrawler.Crawler(args[0], targetPath)
	},
}

func init() {
	get.PersistentFlags().String("target", "./images", fmt.Sprintf("%s get \"IU\" --target ./images\n", appName))

	get.SetUsageTemplate(func() string {
		return "Usage:\n" +
			fmt.Sprintf("  %s get \"IU\"\n", appName) +
			fmt.Sprintf("  %s get \"IU\" --target ./images\n", appName)
	}())
}
