package cmd

import (
	"avalon/utils"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "avalon",
	Short: "Welcome to Avalon, a tool to help auditors certify audits",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.avalon.yaml)")
	checkDirs()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}

		// Search config in home directory with name ".avalon" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".avalon")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func checkDirs() {
	if !utils.DirectoryExists("keys") {
		color.Red("Error: directory keys is missing, please run make install")
		os.Exit(1)
	}
	if !utils.DirectoryExists("data") {
		color.Red("Error: directory data is missing, please run make install")
		os.Exit(1)
	}
	if !utils.DirectoryExists("signatures") {
		color.Red("Error: directory signatures is missing, please run make install")
		os.Exit(1)
	}
}
