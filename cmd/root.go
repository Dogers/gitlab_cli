package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

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
	Long: `GitLab CLI is a program that can run various queries on a given GitLab account. Provide an API token and a group ID to start from.`,
}

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
}

// TODO: this is rubbish. Probably need a print function per activity?
func Printout(intro string, groupPath string, fmtString string, itemType string, outVars []string) {
	switch outputType {
		case "json":
			// Print JSON
			// https://blog.golang.org/json-and-go
			jsonout := json.NewEncoder(os.Stdout)
			for _, item := range outVars {
				_ = jsonout.Encode([]string{groupPath, item})
			}

		case "csv":
			// Print CSV
			// https://golangcode.com/write-data-to-a-csv-file/
			csvout := csv.NewWriter(os.Stdout)
			defer csvout.Flush()

			_ = csvout.Write([]string{"parent_group", itemType})
			for _, item := range outVars{
				_ = csvout.Write([]string{groupPath, item})
			}

		default:
			// Fix illiteracy, redirect to default
			// Print plain text
			tabs := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
			defer tabs.Flush()
			fmt.Println(intro, groupPath)

			for _, item := range outVars {
				_, _ = fmt.Fprintf(tabs, fmtString, item)
			}

			fmt.Println("")
	}
}
