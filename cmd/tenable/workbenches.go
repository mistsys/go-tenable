// nice things to have from this point are
// list all assets *and associated vulnerabilities* in one command
// carve out the low utility commands, or leave them but mark them as interfaces to internal/intermediate APIs
package tenable

import (
	"context"
	"encoding/json"
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
var workbenchesAssetsRootCmd = &cobra.Command{
	Use:   "assets",
	Short: "Use the Tenable workbenches/assets API",
	Args:  cobra.MinimumNArgs(1),
}

var workbenchesAssetsCmd = &cobra.Command{
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

var workbenchesAssetsVulnerabilitiesRootCmd = &cobra.Command{
	Use:     "vulnerabilities",
	Short:   "Use the Tenable workbenches assets vulnerabilities API",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"vulns"},
}

var workbenchesAssetsVulnerabilitiesListCmd = &cobra.Command{
	Use:     "list [ID..]",
	Short:   "List (up to) 5000 assets with vulnerabilities. If asset ID(s) are provided, list vulnerabilities for the specified assets",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			// want vulns specific to the given assets; TODO either restrict to one ID or additionally output the ID in some
			// meaningful way
			for i := 0; i < len(args); i++ {
				// note the function name: this is a singular asset, the other branch is plural
				_, response, err := client.Workbenches.AssetVulnerabilities(context.Background(), args[i])
				if err != nil {
					log.Printf("Error getting vulnerabilites for %s, %v", args[i], err)
				}
				fmt.Printf(response.BodyJson())
			}
		} else {
			_, response, err := client.Workbenches.AssetsVulnerabilities(context.Background())
			if err != nil {
				log.Println("Error getting vulnerabilites", err)
			}
			fmt.Printf(response.BodyJson())
		}
	},
}

var workbenchesAssetVulnerabilityInfoCmd = &cobra.Command{
	Use:   "info assetId pluginId",
	Short: "Get the vulnerability details for a single plugin on a single asset",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		_, response, err := client.Workbenches.AssetVulnerabilityInfo(context.Background(), args[0], args[1])
		if err != nil {
			log.Println("Error getting vulnerability info", err)
		}
		fmt.Printf(response.BodyJson())
	},
}

var workbenchesAssetVulnerabilityOutputsCmd = &cobra.Command{
	Use:   "outputs assetId pluginId",
	Short: "Get the vulnerability outputs for a single plugin on a single asset",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		_, response, err := client.Workbenches.AssetVulnerabilityOutputs(context.Background(), args[0], args[1])
		if err != nil {
			log.Println("Error getting vulnerability outputs", err)
		}
		fmt.Printf(response.BodyJson())
	},
}

// Vulnerabilities commands
var workbenchesVulnerabilitiesRootCmd = &cobra.Command{
	Use:     "vulnerabilities",
	Short:   "Use the Tenable workbenches/vulnerabilities API",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"vulns"},
}

var workbenchesVulnerabilitiesCmd = &cobra.Command{
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

var workbenchesVulnerabilityOutputsCmd = &cobra.Command{
	Use:   "outputs [PLUGIN_ID...]",
	Short: "Get the vulnerability outputs for a single plugin.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < len(args); i++ {
			_, response, err := client.Workbenches.VulnerabilityOutputs(context.Background(), args[i])
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

// Export commands
// export request and status can be rolled together or kept separate
// export download maybe shouldn't be this tool's responsibility; maybe an overall request + status command
// outputs the download url when it's ready?
// TODO requires a lot of options so gotta mature-ize that interface first, probably
var workbenchesExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a workbench to a file and print the download link.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implemeneted")
	},
}

// More powerful commands

var allAssetInfoCmd = &cobra.Command{
	Use:     "listv ID",
	Short:   "List all vulnerability info for a given asset",
	Aliases: []string{"lsv"},
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		assetId := args[0]
		allVulns, err := client.Workbenches.AllAssetInfo(context.Background(), assetId)
		if err != nil {
			log.Println("Error getting vulnerability info for %s, %v", assetId, err)
		}
		b, _ := json.MarshalIndent(allVulns, "  ", "  ")
		fmt.Println(string(b))
		j := &JiraTicket{
			Header: defaultJiraTicketHeaders,
			Source: allVulns,
		}
		ticket, err := j.Produce()
		if err != nil {
			log.Println("Failed to produce jira ticket!", err)
		}
		fmt.Println(ticket)
	},
}

func init() {
	rootCmd.AddCommand(workbenchesCmd)

	workbenchesCmd.AddCommand(workbenchesAssetsRootCmd)
	workbenchesAssetsRootCmd.AddCommand(workbenchesAssetsCmd)
	workbenchesAssetsRootCmd.AddCommand(workbenchesAssetsInfoCmd)
	workbenchesAssetsRootCmd.AddCommand(workbenchesAssetsVulnerabilitiesRootCmd)
	workbenchesAssetsVulnerabilitiesRootCmd.AddCommand(workbenchesAssetsVulnerabilitiesListCmd)
	workbenchesAssetsVulnerabilitiesRootCmd.AddCommand(workbenchesAssetVulnerabilityInfoCmd)
	workbenchesAssetsVulnerabilitiesRootCmd.AddCommand(workbenchesAssetVulnerabilityOutputsCmd)

	workbenchesCmd.AddCommand(workbenchesVulnerabilitiesRootCmd)
	workbenchesVulnerabilitiesRootCmd.AddCommand(workbenchesVulnerabilitiesCmd)
	workbenchesVulnerabilitiesRootCmd.AddCommand(workbenchesVulnerabilitiesInfoCmd)
	workbenchesVulnerabilitiesRootCmd.AddCommand(workbenchesVulnerabilityOutputsCmd)
	workbenchesVulnerabilitiesRootCmd.AddCommand(workbenchesVulnerabilitiesFiltersCmd)
}
