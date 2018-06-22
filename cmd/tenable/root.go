package tenable

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	tenableClient "github.com/mistsys/go-tenable/client"
)

var (
	configFile string
	client     *tenableClient.TenableClient
	verbose    bool
	outputJira bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tenable COMMAND",
	Short: "A CLI for the Tenable API",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		client = tenableClient.NewClient(viper.GetString("accesskey"), viper.GetString("secretkey"))
		// using viper means the cobra *Var flag options don't actually populate the global variables from config files; if you
		// to manually set them anyway, we'll just manually set them...
		debug := viper.GetBool("debug")
		client.Debug = debug
		if debug {
			// regardless of flag, debug mode implies verbose output
			verbose = true
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("accesskey", "k", "", "Tenable Access Key (required)")
	rootCmd.PersistentFlags().StringP("secretkey", "s", "", "Tenable Secret Key (required)")
	rootCmd.MarkFlagRequired("accesskey") // these calls don't do anything if you use viper
	rootCmd.MarkFlagRequired("secretkey") // these calls don't do anything if you use viper
	rootCmd.PersistentFlags().String("impersonate", "", "User to impersonate")

	// TODO
	rootCmd.PersistentFlags().String("query", "", "Query parameters given as a string \"key=value,key=value,...\"")
	// or just json, man, why not
	rootCmd.PersistentFlags().String("payload", "", "JSON payload given as a string '{\"key\": value ... }'")
	rootCmd.PersistentFlags().String("filters", "", "Filters") // TODO doc

	rootCmd.PersistentFlags().BoolVarP(&outputJira, "jira", "j", false, "Produce CSV output suitable for JIRA import. Not available for all commands.")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Produce verbose output. Debug implies verbose.")
	rootCmd.PersistentFlags().Bool("debug", false, "Run in debug mode (dump raw request bodies)")

	rootCmd.PersistentFlags().StringVarP(&configFile, "configFile", "f", "", "Config file to read from")
	flags := rootCmd.PersistentFlags()
	viper.BindPFlag("accesskey", flags.Lookup("accesskey"))
	viper.BindPFlag("secretkey", flags.Lookup("secretkey"))
	viper.BindPFlag("debug", flags.Lookup("debug"))
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("$HOME/.config/tenable/")
		viper.AddConfigPath(".")
	}

	if err := viper.ReadInConfig(); err != nil {
		// fail quietly unless verbose mode is on
		if verbose {
			fmt.Println("Can't read config!", err)
		}
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
