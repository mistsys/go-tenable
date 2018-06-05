package tenable

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	tenableClient "github.com/mistsys/go-tenable/client"
)

var scansCmd = &cobra.Command{
	Use:   "scans COMMAND",
	Short: "Use the Tenable scans API",
	Args:  cobra.MinimumNArgs(1),
}

// fooCmd represents the foo command
var scansListCmd = &cobra.Command{
	Use:   "list [ID...]",
	Short: "List scans. Optionally specify specific scan IDs to view details.",
	Run: func(cmd *cobra.Command, args []string) {
		client = tenableClient.NewClient(accessKey, secretKey)
		client.Debug = debug

		if len(args) > 0 {
			// at least one ID provided, try to get details for all provided IDs
			for i := 0; i < len(args); i++ {
				details, err := client.ScanDetail(context.Background(), args[i])
				if err != nil {
					log.Printf("Error getting scan details for %q, %s", args[i], err)
				}
				fmt.Printf("%v", details)
			}
		} else {
			// no IDs specified, just dump em all
			lst, err := client.ScansList(context.Background())
			if err != nil {
				log.Println("Error getting server scans list", err)
			}
			fmt.Printf("%v", lst)
		}
	},
}

func init() {
	rootCmd.AddCommand(scansCmd)
	scansCmd.AddCommand(scansListCmd)
}
