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
	Use:     "workbenches COMMAND",
	Short:   "Use the Tenable workbenches API",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"wb"},
}

// Assets commands
var workbenchesAssetsCmd = &cobra.Command{
	Use:   "assets",
	Short: "Use the Tenable workbenches/assets API",
	Args:  cobra.MinimumNArgs(1),
}

var workbenchesAssetsListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List (up to) 5000 assets",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		_, response, err := client.Workbenches.Assets(context.Background())
		if err != nil {
			log.Println("Error getting assets", err)
		}
		fmt.Printf(response.BodyJson())
	},
}

var workbenchesAssetsInfoCmd = &cobra.Command{
	Use:   "info ID",
	Short: "Get general information about an asset",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, response, err := client.Workbenches.AssetsInfo(context.Background(), args[0])
		if err != nil {
			log.Println("Error getting asset info", err)
		}
		fmt.Printf(response.BodyJson())
	},
}

// Assets commands
var workbenchesAssetsVulnerabilitiesCmd = &cobra.Command{
	Use:     "vulnerabilities",
	Short:   "Use the Tenable workbenches assets vulnerabilities API",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"vulns"},
}

var workbenchesAssetsVulnerabilitiesListCmd = &cobra.Command{
	Use:     "list ID",
	Short:   "List (up to) 5000 of the vulnerabilities for an asset",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < len(args); i++ {
			_, response, err := client.Workbenches.AssetsVulnerabilities(context.Background(), args[i])
			if err != nil {
				log.Println("Error getting asset vulnerabilites", err)
			}
			fmt.Printf(response.BodyJson())
		}
	},
}

var workbenchesAssetsVulnerabilitiesInfoCmd = &cobra.Command{
	Use:   "info assetId pluginId",
	Short: "Get the details for a vulnerability recorded on a given asset",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		_, response, err := client.Workbenches.AssetsVulnerabilitiesInfo(context.Background(), args[0], args[1])
		if err != nil {
			log.Println("Error getting vulnerability info", err)
		}
		fmt.Printf(response.BodyJson())
	},
}

// Vulnerabilities commands
var workbenchesVulnerabilitiesCmd = &cobra.Command{
	Use:     "vulnerabilities",
	Short:   "Use the Tenable workbenches/vulnerabilities API",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"vulns"},
}

var workbenchesVulnerabilitiesListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List (up to) the first 5000 vulnerabilities recorded.",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		_, response, err := client.Workbenches.Vulnerabilities(context.Background())
		if err != nil {
			log.Println("Error getting vulnerabilities list", err)
		}
		fmt.Printf(response.BodyJson())
		fmt.Println(cmd.Flags().Lookup("query").Value)
	},
}

var workbenchesVulnerabilitiesInfoCmd = &cobra.Command{
	Use:   "info [PLUGIN_ID...]",
	Short: "Get the vulnerability details for a single plugin.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < len(args); i++ {
			_, response, err := client.Workbenches.VulnerabilitiesInfo(context.Background(), args[i])
			if err != nil {
				log.Println("Error getting vulnerability info", err)
			}
			fmt.Printf(response.BodyJson())
		}
	},
}

var workbenchesVulnerabilitiesOutputsCmd = &cobra.Command{
	Use:   "outputs [PLUGIN_ID...]",
	Short: "Get the vulnerability outputs for a single plugin.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < len(args); i++ {
			_, response, err := client.Workbenches.VulnerabilitiesOutputs(context.Background(), args[i])
			if err != nil {
				log.Println("Error getting vulnerability outputs", err)
			}
			fmt.Printf(response.BodyJson())
		}
	},
}

var workbenchesVulnerabilitiesFiltersCmd = &cobra.Command{
	Use:   "filters",
	Short: "Get the filters available for vulnerabilities.",
	Run: func(cmd *cobra.Command, args []string) {
		_, response, err := client.Workbenches.VulnerabilitiesFilters(context.Background())
		if err != nil {
			log.Println("Error getting vulnerabilities filters", err)
		}
		fmt.Printf(response.BodyJson())
	},
}

func init() {
	rootCmd.AddCommand(workbenchesCmd)

	workbenchesCmd.AddCommand(workbenchesAssetsCmd)
	workbenchesAssetsCmd.AddCommand(workbenchesAssetsListCmd)
	workbenchesAssetsCmd.AddCommand(workbenchesAssetsInfoCmd)
	workbenchesAssetsCmd.AddCommand(workbenchesAssetsVulnerabilitiesCmd)
	workbenchesAssetsVulnerabilitiesCmd.AddCommand(workbenchesAssetsVulnerabilitiesListCmd)
	workbenchesAssetsVulnerabilitiesCmd.AddCommand(workbenchesAssetsVulnerabilitiesInfoCmd)

	workbenchesCmd.AddCommand(workbenchesVulnerabilitiesCmd)
	workbenchesVulnerabilitiesCmd.AddCommand(workbenchesVulnerabilitiesListCmd)
	workbenchesVulnerabilitiesCmd.AddCommand(workbenchesVulnerabilitiesInfoCmd)
	workbenchesVulnerabilitiesCmd.AddCommand(workbenchesVulnerabilitiesOutputsCmd)
	workbenchesVulnerabilitiesCmd.AddCommand(workbenchesVulnerabilitiesFiltersCmd)
}
