package main

import "github.com/xanzy/go-gitlab"

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
