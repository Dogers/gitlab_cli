package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/xanzy/go-gitlab"
)

func getSubGroups(gitlabClient *gitlab.Client, mastergroupid string) {

	lSGOpts := &gitlab.ListSubgroupsOptions{}
	groups, _, err := gitlabClient.Groups.ListSubgroups(mastergroupid, lSGOpts)

	if err != nil {
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
}

func getMembersOfGroup(git *gitlab.Client, group *gitlab.Group) {
	lgmOpts := &gitlab.ListGroupMembersOptions{}

	groupMembers, _, err := git.Groups.ListGroupMembers(group.ID, lgmOpts)

	if err != nil {
		log.Fatalf("Error: %o", err)
	}

	if len(groupMembers) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Println("Members found for group:", group.Name)

		for _, gMember := range groupMembers {
			fmt.Fprintf(w, "\t%s\t%s\t%s\n", gMember.Name, gMember.Username, levelToPerm(gMember.AccessLevel))
		}

		w.Flush()
		fmt.Println("")
	}
}
