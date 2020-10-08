package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/xanzy/go-gitlab"
	"log"
	"os"
	"text/tabwriter"
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
		Printout("Members found for group: ", group.FullPath, "\t%s\t%s\t%s\n", "users", vars, 3)
	}
}

// TODO: this is rubbish. Probably need a print function per activity?
func Printout(intro string, groupPath string, fmtString string, itemType string, outVars []string, rowcount int) {
	switch outputType {
	case "json":
		// Print JSON
		// https://blog.golang.org/json-and-go
		jsonout := json.NewEncoder(os.Stdout)
		for c, item := range outVars {
			if (c+1)%rowcount == 0 {
				_ = jsonout.Encode([]string{groupPath, item})
			}
		}

	case "csv":
		// Print CSV
		// https://golangcode.com/write-data-to-a-csv-file/
		csvout := csv.NewWriter(os.Stdout)
		defer csvout.Flush()

		_ = csvout.Write([]string{"parent_group", itemType})
		for c, item := range outVars{
			if (c+1)%rowcount == 0 {
				_ = csvout.Write([]string{groupPath, item})
			}
		}

	default:
		// Fix illiteracy, redirect to default
		// Print plain text
		tabs := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		defer tabs.Flush()
		fmt.Println(intro, groupPath)

		for c, item := range outVars {
			if (c+1)%rowcount == 0 {
				_, _ = fmt.Fprintf(tabs, fmtString, item)
			}
		}

		fmt.Println("")
	}
}
