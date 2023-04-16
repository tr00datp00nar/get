package weather

import (
	"fmt"
	"strings"

	scriptish "github.com/ganbarodigital/go_scriptish"

	"github.com/tr00datp00nar/fn"
)

func getWeather(location string) {
	location = fn.RawUrlEncode(strings.Title(location))
	withoutLocation := "https://wttr.in/?format=3"
	withLocation := "https://wttr.in/" + location + "?format=3"

	if len(location) > 0 {
		pipeline := scriptish.NewPipeline(
			scriptish.Exec("curl", withLocation),
			scriptish.ToStdout(),
		)
		pipeline.Exec()
	} else {
		pipeline := scriptish.NewPipeline(
			scriptish.Exec("curl", withoutLocation),
			scriptish.ToStdout(),
		)
		pipeline.Exec().Strings()
		fmt.Println("\nProvide a location string as an argument for more accurate responses.")
	}

	// if len(location) > 0 {
	// 	resp, err := http.Get(withLocation)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer resp.Body.Close()
	// 	body, err := io.Copy(os.Stdout, resp.Body)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(body)
	// } else {
	// 	resp, err := http.Get(withoutLocation)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer resp.Body.Close()
	// 	body, err := io.Copy(os.Stdout, resp.Body)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(body)
	// }
}
