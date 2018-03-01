package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// ProblemWithConfigFile indicates a problem with the config file
var ProblemWithConfigFile bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Dashboard-service",
	Short: "A simple service to host a family dashboard",
	Long: `Dashboard-service is a simple service for hosting the
	family dashboard UI and settings service`,
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/dashboard.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match

	//	Set our defaults
	viper.SetDefault("server.port", "3000")
	viper.SetDefault("server.bind", "")
	viper.SetDefault("server.allowed-origins", "*")
	viper.SetDefault("datastore.database", "config.db")

	viper.SetConfigName("dashboard") // name of config file (without extension)
	viper.AddConfigPath("$HOME")     // adding home directory as first search path
	viper.AddConfigPath(".")         // also look in the working directory

	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	// If a config file is found, read it in
	// otherwise, make note that there was a problem
	if err := viper.ReadInConfig(); err != nil {
		ProblemWithConfigFile = true
	}
}
