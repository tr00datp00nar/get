package geo

import (
	_ "embed"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:        `geo`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_geo),
	Description: help.D(_geo),

	Commands: []*Z.Cmd{
		wanCmd,
		lanCmd,
		routerCmd,
		dnsCmd,
		macCmd,
		// localCmd,
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

var localCmd = &Z.Cmd{
	Name:        `local`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_local),
	Description: help.D(_local),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		localSearch()
		return nil
	},
}
