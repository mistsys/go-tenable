package tenable

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var workbenchesRootCmd = &cobra.Command{
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
			fmt.Println("Error getting assets:", err)
			os.Exit(1)
		}
		outputter.Output(response.BodyJson())
	},
}

var workbenchesAssetsInfoCmd = &cobra.Command{
	Use:   "info ID",
	Short: "Get general information about an asset",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, response, err := client.Workbenches.AssetsInfo(context.Background(), args[0])
		if err != nil {
			fmt.Println("Error getting asset info:", err)
			os.Exit(1)
		}
		outputter.Output(response.BodyJson())
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
				assetId := args[i]
				allVulns, err := client.Workbenches.AssetVulnerabilityInfoList(context.Background(), assetId)
				if err != nil {
					fmt.Printf("Error getting vulnerabilites for %s, %v", args[i], err)
					os.Exit(1)
				}
				err = outputter.Output(allVulns)
				if err != nil {
					fmt.Printf("Error formatting vulnerabilities for %s, %v", assetId, err)
					os.Exit(1)
				}
			}
		} else {
			_, response, err := client.Workbenches.AssetsVulnerabilities(context.Background())
			if err != nil {
				fmt.Println("Error getting vulnerabilites:", err)
				os.Exit(1)
			}
			outputter.Output(response.BodyJson())
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
			fmt.Println("Error getting vulnerability info:", err)
			os.Exit(1)
		}
		outputter.Output(response.BodyJson())
	},
}

var workbenchesAssetVulnerabilityOutputsCmd = &cobra.Command{
	Use:   "outputs assetId pluginId",
	Short: "Get the vulnerability outputs for a single plugin on a single asset",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		_, response, err := client.Workbenches.AssetVulnerabilityOutputs(context.Background(), args[0], args[1])
		if err != nil {
			fmt.Println("Error getting vulnerability outputs:", err)
			os.Exit(1)
		}
		outputter.Output(response.BodyJson())
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
			fmt.Println("Error getting vulnerabilities list:", err)
			os.Exit(1)
		}
		outputter.Output(response.BodyJson())
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
				fmt.Println("Error getting vulnerability info:", err)
				os.Exit(1)
			}
			outputter.Output(response.BodyJson())
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
				fmt.Println("Error getting vulnerability outputs:", err)
				os.Exit(1)
			}
			outputter.Output(response.BodyJson())
		}
	},
}

var workbenchesVulnerabilitiesFiltersCmd = &cobra.Command{
	Use:   "filters",
	Short: "Get the filters available for vulnerabilities.",
	Run: func(cmd *cobra.Command, args []string) {
		_, response, err := client.Workbenches.VulnerabilitiesFilters(context.Background())
		if err != nil {
			fmt.Println("Error getting vulnerabilities filters:", err)
			os.Exit(1)
		}
		outputter.Output(response.BodyJson())
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
	Run: func(cmd *cobra.Command, args []string) {
		if params == "" {
			fmt.Println("Please provide parameters using the --params flag (see the Tenable API docs)")
			os.Exit(1)
		}

		exportRequest, _, err := client.Workbenches.ExportRequest(context.Background())
		if err != nil {
			fmt.Println("Error initiating export:", err)
			os.Exit(1)
		}

		fileId := exportRequest.File
		fmt.Printf("Started export request with file id: %d\n", fileId)
		fmt.Print("Waiting for file to finish generating.")

		for {
			// poll for export status
			time.Sleep(1 * time.Second)
			fmt.Print(".")
			exportStatus, _, err := client.Workbenches.ExportStatus(context.Background(), fileId)
			if err != nil {
				fmt.Println("Error checking export status::", err)
				fmt.Println("It may still complete successfully, but we're going to stop polling.")
				os.Exit(1)
			}

			if exportStatus.Status == "ready" {
				lol := fmt.Sprintf("/workbenches/export/%d/download", fileId)
				fmt.Printf("\nFile download's not implemented, but the report's ready at %s", lol)
				os.Exit(0)
			} else if exportStatus.Progress == exportStatus.ProgressTotal {
				// For large exports, sometimes the export gets finished but the status doesn't change.
				// It never *does* seem to change once it's stuck, but the file becomes available after a bit of a wait.
				// If you just go for it, sometimes the PDF is empty.
				// bit of repetition
				lol := fmt.Sprintf("/workbenches/export/%d/download", fileId)
				fmt.Printf("\nThe API reports that export progress is complete, but the file may not be ready for download just yet.\n")
				fmt.Printf("File download's not implemented, but the report should be ready soon at %s", lol)
				os.Exit(0)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(workbenchesRootCmd)

	workbenchesRootCmd.AddCommand(workbenchesAssetsRootCmd)
	workbenchesAssetsRootCmd.AddCommand(workbenchesAssetsCmd)
	workbenchesAssetsRootCmd.AddCommand(workbenchesAssetsInfoCmd)
	workbenchesAssetsRootCmd.AddCommand(workbenchesAssetsVulnerabilitiesRootCmd)
	workbenchesAssetsVulnerabilitiesRootCmd.AddCommand(workbenchesAssetsVulnerabilitiesListCmd)
	workbenchesAssetsVulnerabilitiesRootCmd.AddCommand(workbenchesAssetVulnerabilityInfoCmd)
	workbenchesAssetsVulnerabilitiesRootCmd.AddCommand(workbenchesAssetVulnerabilityOutputsCmd)

	workbenchesRootCmd.AddCommand(workbenchesVulnerabilitiesRootCmd)
	workbenchesVulnerabilitiesRootCmd.AddCommand(workbenchesVulnerabilitiesCmd)
	workbenchesVulnerabilitiesRootCmd.AddCommand(workbenchesVulnerabilitiesInfoCmd)
	workbenchesVulnerabilitiesRootCmd.AddCommand(workbenchesVulnerabilityOutputsCmd)
	workbenchesVulnerabilitiesRootCmd.AddCommand(workbenchesVulnerabilitiesFiltersCmd)

	workbenchesRootCmd.AddCommand(workbenchesExportCmd)
}
