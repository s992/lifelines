package server

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path"
	"strings"
)

// stolen from https://github.com/go-chi/chi/issues/611#issuecomment-1804702959
func SPAHandler(files embed.FS) http.HandlerFunc {
	spaFS, err := fs.Sub(files, "dist/client")
	if err != nil {
		panic(fmt.Errorf("failed getting the sub tree for the site files: %w", err))
	}

	return func(w http.ResponseWriter, r *http.Request) {
		f, err := spaFS.Open(strings.TrimPrefix(path.Clean(r.URL.Path), "/"))

		if err == nil {
			defer f.Close()
		}

		if os.IsNotExist(err) {
			r.URL.Path = "/"
		}

		http.FileServer(http.FS(spaFS)).ServeHTTP(w, r)
	}
}
