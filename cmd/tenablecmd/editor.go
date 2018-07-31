package tenablecmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var templateType string

var editorRootCmd = &cobra.Command{
	Use:   "editor COMMAND",
	Short: "Use the Tenable editor API",
	Args:  cobra.MinimumNArgs(1),
}

var editorListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List templates",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		_, response, err := client.Editor.List(context.Background(), templateType)
		if err != nil {
			fmt.Println("Error getting templates:", err)
			os.Exit(1)
		}
		outputter.Output(response.BodyJson())
	},
}

func init() {
	// should make cmd 'registering' functions (like many other cobra-based clis)
	editorListCmd.Flags().StringVar(&templateType, "type", "scan", "Type of template to list. Available options are \"scan\" and \"policy\"")

	rootCmd.AddCommand(editorRootCmd)
	editorRootCmd.AddCommand(editorListCmd)
}
