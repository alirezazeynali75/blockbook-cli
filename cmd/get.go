package cmd

import (
	"fmt"

	"github.com/alirezazeynali75/blockbook-cli/usecase"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCommand)
	getCommand.AddCommand(getStatusCommand)
}

var getCommand = &cobra.Command{
	Use: "get",
	Short: "send get request",
	Long: "send get request to specified blockbook",
}

var getStatusCommand = &cobra.Command{
	Use: "status",
	Short: "get status of somthing",
	Long: "get status of blockbook, args [key(x.x)]",
	Run: func (cmd *cobra.Command, args []string)  {
		status, err := usecase.GetStatus(BlockbookUrl, BlockbookApiKey)
		if err != nil {
			fmt.Println("There is an error: ", err.Error())
		}
		fmt.Println(status)
	},
}