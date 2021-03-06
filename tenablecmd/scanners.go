package tenablecmd

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var scannersRootCmd = &cobra.Command{
	Use:   "scanners COMMAND",
	Short: "Use the Tenable scanners API",
	Args:  cobra.MinimumNArgs(1),
}

var scannersListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List scanners",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		_, response, err := client.Scanners.List(context.Background())
		if err != nil {
			fmt.Println("Error getting scanners:", err)
			os.Exit(1)
		}
		outputter.Output(response.BodyJson())
	},
}

var scannersGetAwsTargetsCmd = &cobra.Command{
	Use:   "targets SCANNER_ID",
	Short: "List targets for an AWS scanner",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		scannerId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("SCANNER_ID must be an int. Got:", args[0])
			os.Exit(1)
		}
		_, response, err := client.Scanners.GetAwsTargets(context.Background(), scannerId)
		if err != nil {
			fmt.Println("Error getting targets:", err)
			os.Exit(1)
		}
		outputter.Output(response.BodyJson())
	},
}

func init() {
	rootCmd.AddCommand(scannersRootCmd)
	scannersRootCmd.AddCommand(scannersListCmd)
	scannersRootCmd.AddCommand(scannersGetAwsTargetsCmd)
}
