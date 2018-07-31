package tenablecmd

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var targets string

var scansCmd = &cobra.Command{
	Use:   "scans COMMAND",
	Short: "Use the Tenable scans API",
	Args:  cobra.MinimumNArgs(1),
}

// TODO the usage interface should match the scans one; ie, split this in two (scans info?)
var scansListCmd = &cobra.Command{
	Use:     "list [ID...]",
	Short:   "List scans. Optionally specify specific scan IDs to view details.",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			// at least one ID provided, try to get details for all provided IDs
			for i := 0; i < len(args); i++ {
				scanId, err := strconv.Atoi(args[i])
				if err != nil {
					fmt.Println("ID must be an int. Got:", args[i])
					os.Exit(1)
				}
				_, response, err := client.Scans.Detail(context.Background(), scanId)
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

var scansLaunchCmd = &cobra.Command{
	Use:   "launch ID",
	Short: "Launch a scan",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		scanId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("ID must be an int. Got:", args[0])
			os.Exit(1)
		}
		_, response, err := client.Scans.Launch(context.Background(), scanId, nil)
		if err != nil {
			fmt.Println("Error launching scans:", err)
			os.Exit(1)
		}
		outputter.Output(response.BodyJson())
	},
}

func init() {
	scansListCmd.Flags().StringVar(&targets, "targets", "", "List of targets to scan instead of scan default")

	rootCmd.AddCommand(scansCmd)
	scansCmd.AddCommand(scansListCmd)
	scansCmd.AddCommand(scansLaunchCmd)
}
