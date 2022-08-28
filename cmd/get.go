package cmd

import (
	"fmt"
	"os"

	"github.com/alirezazeynali75/blockbook-cli/usecase"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCommand)
	getTxCommand.PersistentFlags().String("txid", "", "transaction id that to search")
	getCommand.AddCommand(getStatusCommand)
	getCommand.AddCommand(getTxCommand)
}

var getCommand = &cobra.Command{
	Use: "get",
	Short: "send get request",
	Long: "send get request to specified blockbook",
}

var getStatusCommand = &cobra.Command{
	Use: "status",
	Short: "get status of blockbook",
	Long: "get status of blockbook",
	Run: func (cmd *cobra.Command, args []string)  {
		status, err := usecase.GetStatus(BlockbookUrl, BlockbookApiKey)
		if err != nil {
			fmt.Println("There is an error: ", err.Error())
		}
		fmt.Println(status)
	},
}

var getTxCommand = &cobra.Command{
	Use: "tx",
	Short: "get transaction",
	Long: "get transaction with txId",
	Run: func (cmd *cobra.Command, args []string)  {
		txId, err := cmd.Flags().GetString("txid")
		status, err := usecase.GetTransaction(BlockbookUrl, BlockbookApiKey, txId)
		if err != nil {
			fmt.Println("There is an error: ", err.Error())
			os.Exit(1)
		}
		fmt.Println(status)
	},
}