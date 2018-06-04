package tenable

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	tenableClient "github.com/mistsys/go-tenable/client"
)

var accessKey, secretKey string
var debug bool

var client *tenableClient.TenableClient

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tenable COMMAND",
	Short: "A CLI for the Tenable API",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&accessKey, "accesskey", "k", "", "Tenable Access Key (required)")
	rootCmd.PersistentFlags().StringVarP(&secretKey, "secretkey", "s", "", "Tenable Secret Key (required)")
	rootCmd.MarkFlagRequired("accesskey")
	rootCmd.MarkFlagRequired("secretkey")

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Run in debug mode (dumps raw request body)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
