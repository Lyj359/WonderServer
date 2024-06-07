package statics

import "embed"

//go:embed all:*
var EmbedStaticDir embed.FS
