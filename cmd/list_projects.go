package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
)

var listProjectsCmd = &cobra.Command{
	Use:   "list-projects",
	Short: "Lists projects under a group",
	Long: `Given a group ID, recurse through and list all projects`,
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

			getProjectsInGroup(gitlabClient, group)

			subgroups, _, err := gitlabClient.Groups.ListSubgroups(group.ID, lSGOpts)
			if err != nil {
				log.Fatalf("Error: %o", err)
			}

			for _, subg := range subgroups {
				getProjectsInGroup(gitlabClient, subg)
			}
		}
		// Also need the group itself
		group, _, err := gitlabClient.Groups.GetGroup(masterGID)
		getProjectsInGroup(gitlabClient, group)
	},
}

func init() {
	rootCmd.AddCommand(listProjectsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// groupsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// groupsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getProjectsInGroup(git *gitlab.Client, group *gitlab.Group) {
	lgpoOpts := &gitlab.ListGroupProjectsOptions{}

	groupProjects, _, err := git.Groups.ListGroupProjects(group.ID, lgpoOpts)

	if err != nil {
		log.Fatalf("Error: %o", err)
	}

	if len(groupProjects) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Println("Projects found for group:", group.FullPath)

		for _, gProject := range groupProjects {
			_, _ = fmt.Fprintf(w, "\t%s\n", gProject.Name)
		}

		_ = w.Flush()
		fmt.Println("")
	}
}
