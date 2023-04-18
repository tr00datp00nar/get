package get

import (
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	Z "github.com/rwxrob/bonzai/z"

	"github.com/rwxrob/help"
	"github.com/rwxrob/yq"

	"github.com/briandowns/spinner"
	"github.com/tr00datp00nar/fn"
	"github.com/tr00datp00nar/get/joke"
	"github.com/tr00datp00nar/get/manga"
	"github.com/tr00datp00nar/get/net"
	"github.com/tr00datp00nar/get/url"
	"github.com/tr00datp00nar/get/weather"
	"github.com/tr00datp00nar/get/web"
)

var s1 = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#f4b8e4"))
var s2 = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#8caaee"))

var Cmd = &Z.Cmd{
	Name:        `get`,
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_get),
	Description: help.D(_get),

	Commands: []*Z.Cmd{
		net.Cmd,
		joke.Cmd,
		manga.Cmd,
		url.Cmd,
		yq.Cmd,
		weather.Cmd,
		web.Cmd,
		kittyCmd,
		updateCmd,
		newbonCmd,
		help.Cmd,
	},
}

var kittyCmd = &Z.Cmd{
	Name:      `kitty`,
	Usage:     `[help]`,
	Version:   `v0.0.1`,
	Copyright: `Copyright Micah Nadler 2023`,
	License:   `Apache-2.0`,
	Summary:   `updates the kitty terminal using kitty's update script`,
	// Summary:     help.S(_kitty),
	// Description:  help.D(_kitty),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {

		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Prefix = " "
		s.Suffix = s1.Render(" Updating kitty..... ")
		s.FinalMSG = strings.TrimSpace(s2.Render(" Kitty updated successfully\n"))
		s.Start()

		defer s.Stop()

		cmd := "curl -L https://sw.kovidgoyal.net/kitty/installer.sh | sh /dev/stdin launch=n"
		err := exec.Command("bash", "-c", cmd).Run()
		if err != nil {
			log.Fatal(err)
		}
		return nil
	},
}

var updateCmd = &Z.Cmd{
	Name:      `update`,
	Aliases:   []string{``},
	Usage:     `[help]`,
	Version:   `v0.0.1`,
	Copyright: `Copyright Micah Nadler 2023`,
	License:   `Apache-2.0`,
	Summary:   `runs local instace of topgrade`,
	// Summary:     help.S(_update),
	// Description:  help.D(_update),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		cmd := "topgrade"
		c1 := exec.Command("bash", "-c", cmd)
		c1.Stdout = os.Stdout
		c1.Stderr = os.Stderr
		c1.Run()
		return nil
	},
}

var newbonCmd = &Z.Cmd{
	Name:        `newbon`,
	Usage:       `<bonzaiBranchName> <templateName> | [COMMAND]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_newbon),
	Description: help.D(_newbon),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		bonzaiBranchName := string(args[0])
		templateName := string(args[1])

		cmd := "gh repo create " + bonzaiBranchName + " -p " + templateName + " --public --clone"
		_, _, err := fn.ExecBash(cmd)

		// time.Sleep(5 * time.Second)
		//
		// cmd2 := "git clone git@github.com:tr00datp00nar/" + bonzaiBranchName + " $HOME/Repos/github.com/tr00datp00nar/" + bonzaiBranchName
		// fn.ExecBash(cmd2)

		if err != nil {
			log.Fatal(err)
		}
		return nil
	},
}
