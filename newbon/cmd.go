package newbon

import (
	"fmt"
	"log"
	"regexp"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"github.com/tr00datp00nar/fn"
)

var Cmd = &Z.Cmd{
	Name:        `newbon`,
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_newbon),
	Description: help.D(_newbon),

	Commands: []*Z.Cmd{help.Cmd},

	MaxArgs: 0,

	Call: func(x *Z.Cmd, args ...string) error {
		// Get the branch name
		fmt.Println("Enter the name of your new bonzai branch: ")
		var bonzaiBranchName string
		_, err := fmt.Scanln(&bonzaiBranchName)
		if err != nil {
			log.Fatalf("Could not read user input - %v", err)
		}

		// Get the full path to the template repository
		fmt.Println("Enter the full path to the template repository you would like to use: ")
		var templateName string
		_, err = fmt.Scanln(&templateName)
		if err != nil {
			log.Fatalf("Could not read user input - %v", err)
		}
		pathOK, _ := regexp.MatchString("(^https://)(github\\.com\\/)(\\S*)$", templateName)
		if pathOK == false {
			log.Fatal("Template name is not in the format 'https://github.com/user/repository'. Exiting")
		}

		// Create and close the new repository
		cmd := "gh repo create " + bonzaiBranchName + " -p " + templateName + " --public --clone"
		_, _, err = fn.ExecBash(cmd)
		if err != nil {
			log.Fatalf("Could not create repository - %v", err)
		}

		return nil
	},
}
