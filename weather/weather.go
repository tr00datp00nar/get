package weather

import (
	"fmt"
	"strings"

	"github.com/tr00datp00nar/fn"
)

func getWeather(location string) {
	location = fn.RawUrlEncode(strings.Title(location))
	withoutLocation := "https://wttr.in/"
	withLocation := "https://wttr.in/" + location

	if len(location) > 0 {
		cmd := "curl -L " + withLocation
		fn.ExecBash(cmd)
	} else {
		cmd := "curl -L " + withoutLocation
		fn.ExecBash(cmd)
		fmt.Println("\nProvide a location string as an argument for more accurate responses.")
	}
}
