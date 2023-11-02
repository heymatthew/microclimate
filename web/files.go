package web

import (
	"embed"
	"io/fs"
)

//go:embed templates/* static/*
var Files embed.FS
var Static fs.FS

func init() {
	staticDir, err := fs.Sub(Files, "static")
	if err != nil {
		panic("Missing embedded static files")
	}
	Static = staticDir
}
