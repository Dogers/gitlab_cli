package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
	"log"
	"strings"
)

var listSubGroupMembersCmd = &cobra.Command{
	Use:   "list-subgroup-members",
	Short: "Lists members of subgroups",
	Long: `Recurse through subgroups and lists all members specifically added to subgroups.
Inherited users are not displayed.`,
	Run: func(cmd *cobra.Command, args []string) {

		gitlabClient, _ := gitlab.NewClient(gitlabToken)

		lSGOpts := &gitlab.ListSubgroupsOptions{}
		groups, _, err := gitlabClient.Groups.ListSubgroups(masterGID, lSGOpts)

		if err != nil {
			if strings.Contains(err.Error(), ": 40") {
				log.Fatal("Invalid token, please check and retry")
			}
			log.Fatalf("Error: %o", err)
		}

		for _, group := range groups {

			getMembersOfGroup(gitlabClient, group)

			subgroups, _, err := gitlabClient.Groups.ListSubgroups(group.ID, lSGOpts)
			if err != nil {
				log.Fatalf("Error: %o", err)
			}

			for _, subg := range subgroups {
				getMembersOfGroup(gitlabClient, subg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listSubGroupMembersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// groupsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// groupsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
