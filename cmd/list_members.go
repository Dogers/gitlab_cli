package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
	"log"
	"strings"
)

var listMembersCmd = &cobra.Command{
	Use:   "list-members",
	Short: "Lists members of subgroups",
	Long: `Recurse through subgroups and lists all members specifically added to subgroups.
Inherited users are not displayed.`,
	Run: func(cmd *cobra.Command, args []string) {

		gitlabClient := gitlab.NewClient(nil, gitlabToken)

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
	rootCmd.AddCommand(listMembersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// groupsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// groupsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getMembersOfGroup(git *gitlab.Client, group *gitlab.Group) {
	lgmOpts := &gitlab.ListGroupMembersOptions{}

	groupMembers, _, err := git.Groups.ListGroupMembers(group.ID, lgmOpts)

	if err != nil {
		log.Fatalf("Error: %o", err)
	}

	if len(groupMembers) > 0 {
		vars := []string{}

		for _, gMember := range groupMembers {
			vars = append(vars, gMember.Name)
			vars = append(vars, gMember.Username)
			vars = append(vars, levelToPerm(gMember.AccessLevel))
		}

		// TODO: this won't work - each item (name, username, level) will be a new line!
		printOut("Members found for group: ", group.FullPath, "\t%s\t%s\t%s\n", "users", vars)
	}
}
