package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/enkhalifapro/userq/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd base is responsible for config loading and bootstrapping.
var RootCmd = &cobra.Command{
	Use:   "digitalpriv-backend",
	Short: "API server",
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

var (
	cfgName  string
	cfgPaths string
)

func init() {
	flags := RootCmd.PersistentFlags()
	flags.StringVar(&cfgName, "cfg-name", "development", "config file name without path and extension")
	flags.StringVar(&cfgPaths, "cfg-paths", "./etc", "paths where we search config split them by ','")

	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// NOTE: don't user ENV variables overriding until you will not be sure that config variables not match with
	// system default environment variables.
	// viper.AutomaticEnv()
	if err := config.Load(cfgName, strings.Split(cfgPaths, ",")...); err != nil {
		panic(err.Error())
	}
	log.Infof("loaded config: %v", viper.ConfigFileUsed())
}
