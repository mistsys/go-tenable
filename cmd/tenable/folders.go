package tenable

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var foldersCmd = &cobra.Command{
	Use:   "folders COMMAND",
	Short: "Use the Tenable folders API",
	Args:  cobra.MinimumNArgs(1),
}

var foldersListCmd = &cobra.Command{
	Use:     "list [ID...]",
	Short:   "List folders.",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		_, response, err := client.Folders.List(context.Background())
		if err != nil {
			fmt.Println("Error getting folders list", err)
		}
		// fmt.Printf("%q", lst)
		outputter.Output(response.BodyJson())
	},
}

func init() {
	rootCmd.AddCommand(foldersCmd)
	foldersCmd.AddCommand(foldersListCmd)
}
