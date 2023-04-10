package joke

import _ "embed"

//go:embed text/en/joke.md
var _joke string
