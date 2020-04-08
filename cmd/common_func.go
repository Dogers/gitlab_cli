package cmd

import (
	"github.com/xanzy/go-gitlab"
	"log"
)

func levelToPerm(accesslevel gitlab.AccessLevelValue) string {
	switch accesslevel {
	case 10:
		return "Guest"
	case 20:
		return "Reporter"
	case 30:
		return "Developer"
	case 40:
		return "Maintainer"
	case 50:
		return "Owner"
	default:
		return "??"
	}
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
		Printout("Members found for group: ", group.FullPath, "\t%s\t%s\t%s\n", "users", vars)
	}
}
