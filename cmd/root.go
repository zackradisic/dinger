package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "dinger",
	Short: "Plays a sound after long command execution to increase productivity",
}

// Execute executes the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(runCmd)
}

func initConfig() {
	viper.SetDefault("sound", os.Getenv("GOPATH")+"/src/github.com/zackradisic/dinger/sounds/ding.mp3")
	viper.SetConfigName("config.json")
	viper.AddConfigPath("$GOPATH/src/github.com/zackradisic/dinger")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.WriteConfig()
		} else {
			log.Fatal(err)
		}
	}
}
