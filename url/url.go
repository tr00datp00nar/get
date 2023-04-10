package url

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	qrcode "github.com/skip2/go-qrcode"
)

// ────────────────────────────────────────────────────────────
func ExpandUrl(url string) (string, error) {
	expandedUrl := url

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			expandedUrl = req.URL.String()
			return nil
		},
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	return expandedUrl, nil
}

//      ────────────────────────────────────────────────────────────

func AddExtension(filepath string) (file string) {
	if strings.Contains(filepath, ".png") {
		file := filepath
		return file
	} else {
		file := strings.Join([]string{filepath, ".png"}, "")
		return file
	}
}

func Qrify(args ...string) error {
	l := len(args)

	if l < 1 {
		fmt.Println("Not enough arguments.")
		return nil
	}

	url := args[0]

	var file string

	if l == 1 {
		file = "qr.png"
		err := qrcode.WriteFile(url, qrcode.Medium, 256, file)
		if err != nil {
			return err
		}
	}

	if l == 2 {
		filepath := args[1]
		file := AddExtension(filepath)

		err := qrcode.WriteFile(url, qrcode.Medium, 256, file)
		if err != nil {
			return err
		}
	}

	if l == 3 {
		filepath := args[1]
		size, _ := strconv.Atoi(args[2])
		file := AddExtension(filepath)

		err := qrcode.WriteFile(url, qrcode.Medium, size, file)
		if err != nil {
			return err
		}
	}
	return nil
}

//      ────────────────────────────────────────────────────────────
// curl 'http://api.bit.ly/v3/shorten?login=<bit.ly login>&apiKey=<bit.ly API key>&longURL=<Long URL to be shortened>&format=txt'

// func ShortenUrl(url string) {
// 	var login string
// 	var apiKey string
//
// 	args := "http://api.bit.ly/v3/shorten?login=" + login + "&apiKey=" + apiKey + "&longURL=" + url + "&format=txt"
//
// 	cmd := exec.Command(`curl`, args)
// 	cmd.Run()
// }
