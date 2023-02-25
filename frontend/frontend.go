package frontend

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed dist

var frontend embed.FS

func NewFrontend() fs.FS {

	frontendFS, err := fs.Sub(frontend, "dist")

	if err != nil {
		log.Fatalln(err)
	}

	return frontendFS
}
