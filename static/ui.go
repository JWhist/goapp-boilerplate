// Package static embeds the UI static files into the Go binary.
package static

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed ui/*
var staticFiles embed.FS

// GetUIHandler returns an http.Handler that serves the embedded static UI files
func GetUIHandler() http.Handler {
	// Get the build subdirectory from the embedded filesystem
	buildFS, err := fs.Sub(staticFiles, "ui")
	if err != nil {
		panic(err)
	}

	return http.FileServer(http.FS(buildFS))
}

// GetUIFileSystem returns the embedded filesystem for the UI build files
func GetUIFileSystem() fs.FS {
	buildFS, err := fs.Sub(staticFiles, "ui")
	if err != nil {
		panic(err)
	}
	return buildFS
}
