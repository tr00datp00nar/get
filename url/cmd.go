package url

import (
	"fmt"

	"github.com/charmbracelet/log"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:        `url`,
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_url),
	Description: help.D(_url),

	Commands: []*Z.Cmd{
		expandCmd,
		// shortenCmd,
		qrifyCmd,
		help.Cmd,
	},
}

var expandCmd = &Z.Cmd{
	Name:        `expand`,
	Usage:       `URL`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_expand),
	Description: help.D(_expand),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {
		expandedUrl, err := ExpandUrl(args[0])
		if err != nil {
			log.Error(err)
		}
		fmt.Printf("%s\n", args[0])
		fmt.Printf("%s\n", expandedUrl)
		return nil
	},
}

var qrifyCmd = &Z.Cmd{
	Name:        `qrify`,
	Aliases:     []string{`qr`},
	Usage:       `URL [output-file] [size]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_qrify),
	Description: help.D(_qrify),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {
		Qrify(args...)
		return nil
	},
}

// var shortenCmd = &Z.Cmd{
// 	Name:      `shorten`,
// 	Usage:     `URL`,
// 	Version:   `v0.0.1`,
// 	Copyright: `Copyright Micah Nadler 2023`,
// 	License:   `Apache-2.0`,
// 	// Summary:     `shorten urls`,
// 	// Description: ``,
// 	Commands: []*Z.Cmd{help.Cmd},
//
// 	Call: func(x *Z.Cmd, args ...string) error {
// 		ShortenUrl(args[0])
// 		return nil
// 	},
// }
