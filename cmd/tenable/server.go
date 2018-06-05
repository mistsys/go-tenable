package tenable

import (
	"context"
	"fmt"
	"log"

	tenableClient "github.com/mistsys/go-tenable/client"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server COMMAND",
	Short: "Use the Tenable server API",
	Args:  cobra.MinimumNArgs(1),
}

var serverStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Query server status.",
	Run: func(cmd *cobra.Command, args []string) {
		client = tenableClient.NewClient(accessKey, secretKey)
		client.Debug = debug
		status, err := client.ServerStatus(context.Background())
		if err != nil {
			log.Println("Error getting server status.", err)
		}
		fmt.Printf("%v", status)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.AddCommand(serverStatusCmd)
}
