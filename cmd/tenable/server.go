package tenable

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server COMMAND",
	Short: "Use the Tenable server API",
	Args:  cobra.MinimumNArgs(1),
}

// serverStatus represents the "server/status" command
var serverStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Query server status.",
	Run: func(cmd *cobra.Command, args []string) {
		status, err := client.ServerStatus(context.Background())
		if err != nil {
			log.Printf("Error getting server status. %s", err)
		}
		fmt.Printf("%v", status)
	},
}

var serverPropertiesCmd = &cobra.Command{
	Use:   "properties",
	Short: "Query server properties.",
	Run: func(cmd *cobra.Command, args []string) {
		properties, err := client.ServerProperties(context.Background())
		if err != nil {
			log.Println("Error getting server properties.", err)
		}
		fmt.Printf("%v", properties)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.AddCommand(serverStatusCmd)
	serverCmd.AddCommand(serverPropertiesCmd)
}