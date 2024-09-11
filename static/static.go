package static

import "embed"

//go:embed *
var STATIC_FILES embed.FS
