package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/xanzy/go-gitlab"
)

func main() {
	var gitlabToken string
	var err error

	gitlabToken = ""
	masterGroupID := ""

	validate10 := func(input string) error {
		if len(input) < 10 {
			return errors.New("input must have more than 10 characters")
		}
		return nil
	}

	validate7 := func(input string) error {
		if len(input) < 7 {
			return errors.New("input must have more than 7 characters")
		}
		return nil
	}

	// Check required vars!
	if gitlabToken == "" {
		prompt := promptui.Prompt{
			Label:    "GitLab API token",
			Validate: validate10,
			Mask:     '*',
		}

		gitlabToken, err = prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
	}

	if masterGroupID == "" {
		prompt := promptui.Prompt{
			Label:    "Base group ID",
			Validate: validate7,
			Default:  "1111111",
		}

		masterGroupID, err = prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
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
			//printMenu()
		case "exit":
			os.Exit(0)
		case "listusers":
			getSubGroups(gitlabClient, masterGroupID)
		case "":
		default:
			fmt.Fprintf(os.Stdout, "Unknown command: '%s'\n", command)
			fmt.Println("Type 'help' for available commands")
			fmt.Println()
		}
	}
}
