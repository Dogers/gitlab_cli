package main

import (
	"fmt"
	"log"
	"os"

	"text/tabwriter"

	"github.com/xanzy/go-gitlab"
)

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
