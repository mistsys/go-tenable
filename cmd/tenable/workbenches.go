// TODO when getting details about multiple specific resources, need to pack em up into a json array
// TODO need to consider that the tenable API naming kind of sucks, and it might be more natural to use
// a different command hierarchy than obviously implied by the API. e.g. instead of assets, asset-info id, asset-vulnerabilities id,
// maybe assets should just be its own subcommand (default list) and maybe you get something like assets [list], assets info id...,
// assets vulnerabilities id...
package tenable

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var workbenchesCmd = &cobra.Command{
	Use:   "workbenches COMMAND",
	Short: "Use the Tenable workbenches API",
	Args:  cobra.MinimumNArgs(1),
}

var workbenchesAssetsCmd = &cobra.Command{
	Use:   "assets",
	Short: "List (up to) 5000 assets.",
	Run: func(cmd *cobra.Command, args []string) {
		_, response, err := client.Workbenches.Assets(context.Background())
		if err != nil {
			log.Println("Error getting assets list", err)
		}
		// fmt.Printf("%v", assets)
		fmt.Printf(response.BodyJson())
	},
}

var workbenchesAssetInfoCmd = &cobra.Command{
	Use:   "assets-info ID",
	Short: "Get general information about an asset",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, response, err := client.Workbenches.AssetsInfo(context.Background(), args[0])
		if err != nil {
			log.Println("Error getting asset info", err)
		}
		// fmt.Printf("%v", assetInfo)
		fmt.Printf(response.BodyJson())
	},
}

// TODO API returns up to first 5000 recorded vulns; add a note about how to get more
// (once the workbenches export API is implemented)
var workbenchesVulnerabilitiesCmd = &cobra.Command{
	Use:   "vulnerabilities",
	Short: "List (up to) the first 5000 vulnerabilities recorded.",
	Run: func(cmd *cobra.Command, args []string) {
		_, response, err := client.Workbenches.Vulnerabilities(context.Background())
		if err != nil {
			log.Println("Error getting vulnerabilities list", err)
		}
		// fmt.Printf("%v", vulns)
		fmt.Printf(response.BodyJson())
		fmt.Println(cmd.Flags().Lookup("query").Value)
	},
}

var workbenchesVulnerabilitiesInfoCmd = &cobra.Command{
	Use:   "vulnerabilities-info [PLUGIN_ID...]",
	Short: "Get the vulnerability details for a single plugin.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < len(args); i++ {
			_, response, err := client.Workbenches.VulnerabilitiesInfo(context.Background(), args[i])
			if err != nil {
				log.Println("Error getting vulnerability info", err)
			}
			// fmt.Printf("%v", info)
			fmt.Printf(response.BodyJson())
		}
	},
}

var workbenchesVulnerabilitiesOutputsCmd = &cobra.Command{
	Use:   "vulnerabilities-output [PLUGIN_ID...]",
	Short: "Get the vulnerability outputs for a single plugin.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < len(args); i++ {
			_, response, err := client.Workbenches.VulnerabilitiesOutputs(context.Background(), args[i])
			if err != nil {
				log.Println("Error getting vulnerability outputs", err)
			}
			// fmt.Printf("%v", outputs)
			fmt.Printf(response.BodyJson())
		}
	},
}

var workbenchesVulnerabilitiesFiltersCmd = &cobra.Command{
	Use:   "vulnerabilities-filters",
	Short: "Get the vilters available for vulnerabilities.",
	Run: func(cmd *cobra.Command, args []string) {
		_, response, err := client.Workbenches.VulnerabilitiesFilters(context.Background())
		if err != nil {
			log.Println("Error getting vulnerabilities filters", err)
		}
		// fmt.Printf("%v", outputs)
		fmt.Printf(response.BodyJson())
	},
}

func init() {
	rootCmd.AddCommand(workbenchesCmd)
	workbenchesCmd.AddCommand(workbenchesAssetsCmd)
	workbenchesCmd.AddCommand(workbenchesAssetInfoCmd)
	workbenchesCmd.AddCommand(workbenchesVulnerabilitiesCmd)
	workbenchesCmd.AddCommand(workbenchesVulnerabilitiesInfoCmd)
	workbenchesCmd.AddCommand(workbenchesVulnerabilitiesOutputsCmd)
	workbenchesCmd.AddCommand(workbenchesVulnerabilitiesFiltersCmd)
}
