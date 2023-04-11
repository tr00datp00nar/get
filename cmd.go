package get

import (
	Z "github.com/rwxrob/bonzai/z"

	"github.com/rwxrob/help"
	"github.com/rwxrob/yq"

	"github.com/tr00datp00nar/get/geo"
	"github.com/tr00datp00nar/get/joke"
	"github.com/tr00datp00nar/get/manga"
	"github.com/tr00datp00nar/get/url"
	"github.com/tr00datp00nar/get/web"
)

var Cmd = &Z.Cmd{
	Name:        `get`,
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_get),
	Description: help.D(_get),

	Commands: []*Z.Cmd{
		geo.Cmd,
		joke.Cmd,
		manga.Cmd,
		url.Cmd,
		yq.Cmd,
		web.Cmd,
		help.Cmd,
	},
}
