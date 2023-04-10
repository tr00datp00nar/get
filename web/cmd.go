package web

import (
	"fmt"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
	"github.com/rwxrob/help"
	"github.com/rwxrob/vars"
)

// main branch
var Cmd = &Z.Cmd{

	Name:        `web`,
	Summary:     help.S(_web),
	Description: help.D(_web),
	Version:     `v0.5.0`,
	Copyright:   `Copyright 2021 Robert S Muhlestein`,
	License:     `Apache-2.0`,
	Source:      `git@github.com:rwxrob/web.git`,
	Issues:      `github.com/rwxrob/web/issues`,

	Commands: []*Z.Cmd{
		help.Cmd, conf.Cmd, vars.Cmd, // common
		get, // post, put, del|delete, patch, dl|download
	},
}

var get = &Z.Cmd{

	Name:        `get`,
	Summary:     help.S(_get),
	Description: help.D(_get),
	Commands:    []*Z.Cmd{help.Cmd},

	MinArgs: 1,
	MaxArgs: 2,

	Call: func(_ *Z.Cmd, args ...string) error {
		req := Req{U: args[0], D: ""}
		if err := req.Submit(); err != nil {
			return err
		}
		fmt.Println(req.D)
		return nil
	},
}
