package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
	"log"
	"strings"
)

var listMembersCmd = &cobra.Command{
	Use:   "list-members",
	Short: "Lists members of a group",
	Long:  `Lists all members of specific group`,
	Run: func(cmd *cobra.Command, args []string) {

		gitlabClient, _ := gitlab.NewClient(gitlabToken)

		group, _, err := gitlabClient.Groups.GetGroup(masterGID)

		if err != nil {
			if strings.Contains(err.Error(), ": 40") {
				log.Fatal("Invalid token, please check and retry")
			}
			log.Fatalf("Error: %o", err)
		}

		getMembersOfGroup(gitlabClient, group)
	},
}

func init() {
	rootCmd.AddCommand(listMembersCmd)
}
