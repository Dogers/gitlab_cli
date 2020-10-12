package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mitchellh/go-homedir"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var masterGID int
var gitlabToken string
var outputType string
var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "gitlab_cli",
	Short: "GitLab CLI is a program for running queries on a given GitLab account",
	Long:  `GitLab CLI is a program that can run various queries on a given GitLab account. Provide an API token and a group ID to start from.`,
}

// Execute kicks it all off
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// read in Viper settings
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory for config file
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}
		viper.SetConfigType("yaml")
		viper.SetConfigName(".gitlab/gitlab_cli")
		viper.AddConfigPath(home)
	}

	if err := viper.ReadInConfig(); err != nil {
		if cfgFile != "" {
			log.Println("config specified but unable to read it, using defaults")
		}
	}

	rootCmd.PersistentFlags().IntVarP(&masterGID, "gid", "g", -1, "The topmost group ID to recurse from")
	_ = rootCmd.MarkPersistentFlagRequired("gid")

	rootCmd.PersistentFlags().StringVarP(&gitlabToken, "token", "t", "", "A valid token for accessing the GitLab API")
	_ = rootCmd.MarkPersistentFlagRequired("token")

	// If token is in config file, set it!
	if viper.IsSet("token") {
		_ = rootCmd.PersistentFlags().Set("token", viper.GetString("token"))
	}

	rootCmd.PersistentFlags().StringVarP(&outputType, "output", "o", "text", "Output format for results, valid options are text, csv and json")

	// Various settings
	cobra.MousetrapDisplayDuration = 10 * time.Second
	cobra.MousetrapHelpText = rootCmd.Long + "\n\nIt is designed to be run via a terminal.\n\nClosing in " + cobra.MousetrapDisplayDuration.String()
}
