package manga

import (
	"os"
	"os/exec"
)

func PullManga() {
	// TODO: Use a relative path instead of hardcoding the $HOME dir
	cmd := exec.Command("mangadex-dl", "/home/micah/.config/mangadex-dl/list.txt")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
