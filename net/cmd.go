package net

import (
	_ "embed"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:        `net`,
	Aliases:     []string{``},
	Usage:       `COMMAND|[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_net),
	Description: help.D(_net),

	Commands: []*Z.Cmd{
		wanCmd,
		lanCmd,
		routerCmd,
		dnsCmd,
		macCmd,
		wifiPasswdCmd,
		checkCmd,
		help.Cmd,
	},
}

var wanCmd = &Z.Cmd{
	Name:        `wan`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_wan),
	Description: help.D(_wan),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		wanSearch()
		return nil
	},
}

var lanCmd = &Z.Cmd{
	Name:        `lan`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_lan),
	Description: help.D(_lan),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		lanSearch()
		return nil
	},
}

var routerCmd = &Z.Cmd{
	Name:        `router`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_router),
	Description: help.D(_router),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		routerSearch()
		return nil
	},
}

var macCmd = &Z.Cmd{
	Name:        `mac`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_mac),
	Description: help.D(_mac),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		macSearch()
		return nil
	},
}

var checkCmd = &Z.Cmd{
	Name:        `check`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_check),
	Description: help.D(_check),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {
		checkNetwork()
		return nil
	},
}

var dnsCmd = &Z.Cmd{
	Name:        `dns`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_dns),
	Description: help.D(_dns),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		dnsSearch()
		return nil
	},
}

var wifiPasswdCmd = &Z.Cmd{
	Name:        `wifipass`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2024`,
	License:     `Apache-2.0`,
	Summary:     help.S(_wifipass),
	Description: help.D(_wifipass),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {
		wifiPasswd()
		return nil
	},
}
