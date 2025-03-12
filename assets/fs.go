package assets

import "embed"

//go:embed all:views
var ViewFS embed.FS

//go:embed all:static
var StaticFS embed.FS
