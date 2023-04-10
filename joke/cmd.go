package joke

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/icelain/jokeapi"
	Z "github.com/rwxrob/bonzai/z"

	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:        `joke`,
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_joke),
	Description: help.D(_joke),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {
		api := jokeapi.New()

		response, err := api.Fetch()
		if err != nil {
			log.Fatal(err)
		}
		if len(response.Joke) > 1 {
			fmt.Println(response.Joke[0])
			fmt.Println(response.Joke[1])

		} else {
			fmt.Println(response.Joke[0])
		}

		return nil
	},
}
