package cmd

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var gitlabToken string
var masterGroupID int
var err error

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
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
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gitlab_cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().Int("gid", masterGroupID, "The topmost group ID to recurse from")
}

func initConfig() {
	viper.SetEnvPrefix("gitlab_cli")
	viper.AutomaticEnv()

	gitlabToken = viper.GetString("token")
	if gitlabToken == "" {
		prompt := promptui.Prompt{
			Label:    "GitLab API token",
			Validate: validate10,
			Mask:     '*',
		}

		gitlabToken, err = prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
	}

	masterGroupID = viper.GetInt("gid")
	if masterGroupID == "" {
		prompt := promptui.Prompt{
			Label:    "Base group ID",
			Validate: validate10,
			Default:  "1111111",
		}

		masterGroupID, err = prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
	}
}

func validate10(input string) error {
	if len(input) < 10 {
		return errors.New("input must have more than 10 characters")
	}
	return nil
}
