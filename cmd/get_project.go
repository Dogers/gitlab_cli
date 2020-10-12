package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
)

var getGroupCmd = &cobra.Command{
	Use:   "get-group",
	Short: "Gets a given group IDs details",
	Long: `Given a group ID, print details of that group`,
	Run: func(cmd *cobra.Command, args []string) {

		gitlabClient, _ := gitlab.NewClient(gitlabToken)

		group, _, err := gitlabClient.Groups.GetGroup(masterGID)

		if err != nil {
			if strings.Contains(err.Error(), ": 40") {
				log.Fatal("Invalid token, please check and retry")
			}
			log.Fatalf("Error: %o", err)
		}

		Printout("Found group:", group.FullPath, "", string(group.ParentID), nil, 3)
	},
}

func init() {
	rootCmd.AddCommand(getGroupCmd)
}
