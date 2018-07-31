package tenablecmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	tenable "github.com/mistsys/go-tenable/client"
	"github.com/mistsys/go-tenable/outputs"
)

var (
	configFile     string
	client         *tenable.Client
	outputter      *outputs.Outputter
	params         string
	verbose        bool
	outputFilename string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tenable COMMAND",
	Short: "A CLI for the Tenable API",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		client = tenable.NewClient(viper.GetString("accesskey"), viper.GetString("secretkey"))
		// using viper means the cobra *Var flag options don't actually populate the global variables from config files; if you
		// to manually set them anyway, we'll just manually set them...
		debug := viper.GetBool("debug")
		client.Debug = debug
		queryOpts := &tenable.TenableQueryOpts{Params: params}
		client.QueryOpts = queryOpts

		var outputFd *os.File
		var err error // ???
		if outputFilename == "-" {
			outputFd = os.Stdout
		} else {
			outputFd, err = outputs.NewFile(outputFilename)
			if err != nil {
				fmt.Println("Error creating output file:", err)
				os.Exit(1)
			}
		}
		outputter = outputs.NewOutputter(verbose, viper.GetString("format"), outputFd)
	},
}

func init() {
	// authn
	rootCmd.PersistentFlags().StringP("accesskey", "k", "", "Tenable Access Key (required)")
	rootCmd.PersistentFlags().StringP("secretkey", "s", "", "Tenable Secret Key (required)")
	rootCmd.MarkFlagRequired("accesskey") // XXX these calls don't do anything if you use viper
	rootCmd.MarkFlagRequired("secretkey") // XXX these calls don't do anything if you use viper
	rootCmd.PersistentFlags().String("impersonate", "", "User to impersonate")

	// request params
	rootCmd.PersistentFlags().StringVar(&params, "params", "", "Query parameters given as a string of \"key=value,key=value,...\"")
	// TODO next two
	rootCmd.PersistentFlags().String("payload", "", "JSON payload given as a string '{\"key\": value ... }'")
	rootCmd.PersistentFlags().String("filters", "", "Filters")

	// output args
	rootCmd.PersistentFlags().StringVarP(&outputFilename, "output-file", "o", "-", "Output file. Passing `-` writes to stdout (default)")
	rootCmd.PersistentFlags().String("format", "json", "Output format. Available options are `json` (default) and `jira` (not available for all commands)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	rootCmd.PersistentFlags().Bool("debug", false, "Run in debug mode (dump raw request bodies)")

	// viper stuff for using config files
	rootCmd.PersistentFlags().StringVarP(&configFile, "configFile", "f", "", "Config file to read from")
	flags := rootCmd.PersistentFlags()
	viper.BindPFlag("accesskey", flags.Lookup("accesskey"))
	viper.BindPFlag("secretkey", flags.Lookup("secretkey"))
	viper.BindPFlag("format", flags.Lookup("format"))
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
		// fail quietly unless verbose mode is on; config may have been passed as args
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