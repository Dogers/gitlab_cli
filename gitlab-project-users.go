package main

import (
    "github.com/xanzy/go-gitlab"
    "log"
)

func main() {

    git := gitlab.NewClient(nil, "")

    lSGOpts := &gitlab.ListSubgroupsOptions{}
    groups, _, err := git.Groups.ListSubgroups("1111111", lSGOpts)

    if err != nil {
        log.Fatalf("Error: %o", err)

    }

    for _, group := range groups {
        //log.Println("Found Group: ", group.Name)

        getMembersOfGroup(git, group)

        group_subgroups, _, err := git.Groups.ListSubgroups(group.ID, lSGOpts)
        if err != nil {
            log.Fatalf("Error: %o", err)

        }

        for _, subg := range group_subgroups {
            //log.Println("Found Sub Group: ", subg.Name)
            getMembersOfGroup(git, subg)
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
        log.Println("Members found for group: ", group.Name)
        for _, gMember := range groupMembers {
            log.Printf("\t%s(%s) access: %d", gMember.Name, gMember.Username, gMember.AccessLevel)
        }
    }
}
