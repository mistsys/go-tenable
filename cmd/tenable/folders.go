package tenable

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var foldersCmd = &cobra.Command{
	Use:   "folders COMMAND",
	Short: "Use the Tenable folders API",
	Args:  cobra.MinimumNArgs(1),
}

var foldersListCmd = &cobra.Command{
	Use:   "list [ID...]",
	Short: "List folders.",
	Run: func(cmd *cobra.Command, args []string) {
		lst, err := client.Folders.List(context.Background())
		if err != nil {
			log.Println("Error getting folders list", err)
		}
		fmt.Printf("%q", lst)
	},
}

func init() {
	rootCmd.AddCommand(foldersCmd)
	foldersCmd.AddCommand(foldersListCmd)
}
