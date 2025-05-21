package web

import "embed"

//go:embed dist static
var WebFS embed.FS
