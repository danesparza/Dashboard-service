package cmd

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/hashicorp/logutils"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile               string
	problemWithConfigFile bool
	loglevel              string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Dashboard-service",
	Short: "A simple service to host a dashboard",
	Long:  `Dashboard-service is a simple service for hosting the dashboard UI and settings service`,
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
	rootCmd.PersistentFlags().StringVarP(&loglevel, "loglevel", "l", "WARN", "Log level: DEBUG/INFO/WARN/ERROR")

	//	Bind config flags for optional config file override:
	viper.BindPFlag("loglevel", rootCmd.PersistentFlags().Lookup("loglevel"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("dashboard") // name of config file (without extension)
		viper.AddConfigPath(home)        // adding home directory as first search path
		viper.AddConfigPath(".")         // also look in the working directory
	}

	viper.AutomaticEnv() // read in environment variables that match

	//	Set our defaults
	viper.SetDefault("loglevel", "INFO")
	viper.SetDefault("server.port", "3000")
	viper.SetDefault("server.bind", "")
	viper.SetDefault("server.allowed-origins", "*")
	viper.SetDefault("datastore.system", path.Join(home, "dashboard-service", "db", "system"))

	// If a config file is found, read it in
	// otherwise, make note that there was a problem
	if err := viper.ReadInConfig(); err != nil {
		problemWithConfigFile = true
	}

	//	Set the log level from config (if we have it)
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel(viper.GetString("loglevel")),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)

	//	If we have a config file, report it:
	if viper.ConfigFileUsed() != "" {
		log.Println("[DEBUG] Using config file:", viper.ConfigFileUsed())
	}
}
