package tenable

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var scansCmd = &cobra.Command{
	Use:   "scans COMMAND",
	Short: "Use the Tenable scans API",
	Args:  cobra.MinimumNArgs(1),
}

// TODO the usage interface should match the workbenches one; ie, split this in two (scans info?)
var scansListCmd = &cobra.Command{
	Use:     "list [ID...]",
	Short:   "List scans. Optionally specify specific scan IDs to view details.",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			// at least one ID provided, try to get details for all provided IDs
			for i := 0; i < len(args); i++ {
				_, response, err := client.Scans.Detail(context.Background(), args[i])
				if err != nil {
					fmt.Printf("Error getting scan details for %q, %s", args[i], err)
				}
				// fmt.Printf("%v", details)
				outputter.Output(response.BodyJson())
			}
		} else {
			// no IDs specified, just dump em all
			_, response, err := client.Scans.List(context.Background())
			if err != nil {
				fmt.Println("Error getting server scans list", err)
			}
			// fmt.Printf("%v", lst)
			outputter.Output(response.BodyJson())
		}
	},
}

func init() {
	rootCmd.AddCommand(scansCmd)
	scansCmd.AddCommand(scansListCmd)
}
