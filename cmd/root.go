package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var masterGID int
var gitlabToken string
var outputType string

var rootCmd = &cobra.Command{
	Use:   "gitlab_cli",
	Short: "GitLab CLI is a program for running queries on a given GitLab account",
	Long: `GitLab CLI is a program that can run various queries on a given GitLab account.
Just provide an API token and a group ID to start from.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gitlab_cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().IntVarP(&masterGID, "gid", "g", -1, "The topmost group ID to recurse from")
	_ = rootCmd.MarkPersistentFlagRequired("gid")

	rootCmd.PersistentFlags().StringVarP(&gitlabToken, "token", "t", "", "A valid token for accessing the GitLab API")
	_ = rootCmd.MarkPersistentFlagRequired("token")

	rootCmd.PersistentFlags().StringVarP(&outputType, "output", "o", "default", "Output format for results, valid options are default, csv and json")
}

func printOut(intro string, groupPath string, fmtString string, outVars []string) {
	switch outputType {
		case "default":
			// Print plain text
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
			fmt.Println(intro, groupPath)

			for _, item := range outVars {
				_, _ = fmt.Fprintf(w, fmtString, item)
			}

			_ = w.Flush()
			fmt.Println("")

		case "json":
			// Print JSON
			// https://blog.golang.org/json-and-go
		case "csv":
			// Print CSV
			// https://golangcode.com/write-data-to-a-csv-file/
		default:
			// Fix illiteracy, redirect to default
			outputType = "default"
			printOut(intro, groupPath, fmtString, outVars)
	}
}
