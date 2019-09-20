package main

import (
	"log"

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
