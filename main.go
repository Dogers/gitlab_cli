package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/xanzy/go-gitlab"
)

func main() {
	var gitlabToken string
	var err error

	gitlabToken = ""
	masterGroupId := "1111111"

	if gitlabToken == "" {
		fmt.Println("Gitlab API token required: ")
		reader := bufio.NewReader(os.Stdin)
		gitlabToken, err = reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error: %o", err)
		}
	}

	gitlabClient := gitlab.NewClient(nil, gitlabToken)

	// Generate menu
	getSubGroups(gitlabClient, masterGroupId)
}
