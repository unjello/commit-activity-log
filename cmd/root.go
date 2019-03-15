package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

const (
	configFileName = ".commit-activity-log"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use:   "commit-activity-log",
	Short: "cal: Commit Activity Log",
	Long:  "Git Activity Log and Statistics",
}

// Execute run program main loop
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is $HOME/%s", configFileName))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if runtime.GOOS != "windows" {
			viper.AddConfigPath("/etc/commit-activity-log")
		}
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigFile(configFileName)
	}
}
