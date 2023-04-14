package manga

import (
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/charmbracelet/log"
)

func PullManga() {
	user, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	homeDir := user.HomeDir
	listFile := ".config/mangadex-dl/list.txt"
	confDir := strings.Join([]string{homeDir, listFile}, "/")

	cmd := exec.Command("mangadex-dl", confDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
