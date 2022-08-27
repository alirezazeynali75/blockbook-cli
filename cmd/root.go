package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	BlockbookUrl string
	BlockbookApiKey string
	rootCmd = &cobra.Command{
		Use:   "blockbook",
		Short: "blockbook cli client",
		Long:  `blockbook cli client to call Trezor Blockbook Methods`,
	}
)


func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&BlockbookUrl, "url", "", "base url for blockbook")
	rootCmd.PersistentFlags().StringVar(&BlockbookApiKey, "apikey", "", "api key for blockbook")
}

func initConfig() {
	if BlockbookUrl == "" {
		fmt.Println("url not specified use --url for specification")
		os.Exit(1)
	}
}