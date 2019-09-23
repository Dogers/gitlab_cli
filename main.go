package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

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

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("GL> ")
		scanner.Scan()
		command := scanner.Text()

		switch strings.ToLower(command) {
		case "help":
			printMenu()
		case "exit":
			os.Exit(0)
		case "listusers":
			getSubGroups(gitlabClient, masterGroupID)
		default:
			fmt.Fprintf(os.Stdout, "Unknown command: '%s'\n", command)
			fmt.Println("Type 'help' for available commands")
			fmt.Println()
		}
	}
}

func printMenu() {
	// TODO: can this be auto generated somehow?
	fmt.Println("GitLab CLI")
	fmt.Println("----------")
	fmt.Println()
	fmt.Println()
}
