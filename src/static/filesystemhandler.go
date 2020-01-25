package static

import (
	"net/http"
	"strings"
)

// FileSystem custom file system handler.
type CustomFileSystem struct {
	fs http.FileSystem
}

func (fs CustomFileSystem) Open(path string) (http.File, error) {
	file, err := fs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	stat, err := file.Stat()
	if stat.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := fs.fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}
