package weather

import (
	_ "embed"
	"strings"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:        `weather`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_weather),
	Description: help.D(_weather),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {
		location := strings.Join(args, " ")
		getWeather(location)
		// numArgs := len(args)
		// if numArgs < 1 {
		// 	getWeather("")
		// }
		// if numArgs >= 1 {
		// 	getWeather(args[0])
		// }
		return nil
	},
}
