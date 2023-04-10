package manga

import (
	"os"
	"os/exec"
)

func PullManga() {
	cmd := exec.Command("mangadex-dl", "/home/micah/.config/mangadex-dl/list.txt")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
