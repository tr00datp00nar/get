# :deciduous_tree: get

This is `get`. I have put it into a command branch for inclusion into my c Bonzai stateful command tree.

## Installation

If you just want to try it out, grab the release binary with curl and put into your PATH:

```
curl -L https://github.com/tr00datp00nar/get/releases/latest/download/get-linux-amd64 -o ~/.local/bin/get
curl -L https://github.com/tr00datp00nar/get/releases/latest/download/get-darwin-amd64 -o ~/.local/bin/get
curl -L https://github.com/tr00datp00nar/get/releases/latest/download/get-darwin-arm64 -o ~/.local/bin/get
curl -L https://github.com/tr00datp00nar/get/releases/latest/download/get-windows-amd64 -o ~/.local/bin/get
```

Or with `go`:

```shell
go install github.com/tr00datp00nar/get/cmd/get@latest
```

Composed

```go
package c

import (
	Z "github.com/rwxrob/bonzai/z"
    "github.com/tr00datp00nar/get"
)

var Cmd = &Z.Cmd{
	Name:     'c',
    Commands: []*Z.Cmd{help.Cmd, get.Cmd},
}
```

## Included Branches

### Joke
Queries Sv443's Joke API (https://v2.jokeapi.dev/joke/) and returns a random joke.
### Manga
Uses `mangadex-dl` to manage mangadex downloads.
Expects that `mangadex-dl` has been configured with all desired options and that a list of manga to maintain exists at `$XDG_CONFIG_HOME/mangadex/list.txt`.
### Net
Set of tools for getting network interface information.
### Url
Tools for working with URLs
### Weather
Queries `http://wttr.in` and returns the current weather.
### Web
Tool for making common web requests.

## Resources

To learn more about Bonzai stateful command trees: https://github.com/rwxrob/bonzai

To see my personal Bonzai stateful command tree: https://github.com/tr00datp00nar/c

To see the original Bonzai stateful command tree z: https://github.com/rwxrob/z
