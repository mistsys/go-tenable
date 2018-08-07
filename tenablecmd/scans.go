package tenablecmd

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/mistsys/go-tenable/outputs"
)

var format string

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
	Use:   "launch SCAN_ID",
	Short: "Launch a scan",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		scanId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("SCAN_ID must be an int. Got:", args[0])
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

var scansExportCmd = &cobra.Command{
	Use:   "export SCAN_ID",
	Short: "Export the results of a scan",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if outputFilename == "-" && strings.ToLower(format) == "jira" {
			fmt.Println("JIRA output requires an output file.")
			os.Exit(1)
		}
		scanId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("SCAN_ID must be an int. Got:", args[0])
			os.Exit(1)
		}

		// exportFormat is what tenable will create, format is what we'll output
		exportFormat := format
		if strings.ToLower(format) == "jira" {
			exportFormat = "csv"
		}
		exportRequest, _, err := client.Scans.ExportRequest(context.Background(), scanId, exportFormat)
		if err != nil {
			fmt.Println("Error initiating export:", err)
			os.Exit(1)
		}

		fileId := exportRequest.File
		fmt.Printf("Started export request with file id: %d\n", fileId)
		fmt.Print("Waiting for file to finish generating.")
		u := fmt.Sprintf("/scans/%d/export/%d/download", scanId, fileId)

		for {
			// poll for export status
			time.Sleep(1 * time.Second)
			fmt.Print(".")
			exportStatus, _, err := client.Scans.ExportStatus(context.Background(), scanId, fileId)
			if err != nil {
				fmt.Println("Error checking export status::", err)
				fmt.Println("It may still complete successfully, but we're going to stop polling.")
				fmt.Printf("Should it complete, the report will be available at %s (must be authenticated)\n", u)
				os.Exit(1)
			}

			if exportStatus.Status == "ready" {
				fmt.Println()
				break
			}
		}

		// if we're here, it's ready
		if outputFilename != "-" {
			fd, err := outputs.NewFile(outputFilename)
			defer fd.Close()
			if err != nil {
				fmt.Printf("File error attempting to open %s: %v\n", outputFilename, err)
			} else {
				resp, _ := client.PlainGet(context.Background(), u)
				defer resp.Body.Close()
				if strings.ToLower(format) == "jira" {
					outputs.WriteTenableToJira(resp.Body, fd)
				} else {
					_, err := io.Copy(fd, resp.Body)
					if err != nil {
						fmt.Printf("Error writing to %s: %v\n", outputFilename, err)
						os.Exit(1)
					}
				}
				fmt.Printf("Wrote to %s\n", outputFilename)
				os.Exit(0)
			}
		}
		fmt.Printf("\nReport (%s) is ready at %s (must be authenticated)", exportFormat, u)
	},
}

func init() {
	scansExportCmd.Flags().StringVar(&format, "format", "csv", "Output format. Available options are csv, pdf, html, jira")
	rootCmd.AddCommand(scansCmd)
	scansCmd.AddCommand(scansListCmd)
	scansCmd.AddCommand(scansLaunchCmd)
	scansCmd.AddCommand(scansExportCmd)
}
