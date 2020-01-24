package static

import (
	"net/http"
	"os"
	"strings"

	"github.com/yarf-framework/yarf"
)

type File struct {
	// Implements Resource
	yarf.Resource

	// Points to www_root
	Path string

	// Prefix to exclude on path construction
	Prefix string
}

// Implement the GET handler
func (f *File) Get(c *yarf.Context) error {
	// Construct path
	path := f.Path + strings.TrimPrefix(c.Request.URL.EscapedPath(), f.Prefix)

	// Check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Check that isn't index request
		if _, err := os.Stat(path + "/index.html"); os.IsNotExist(err) {
			return yarf.ErrorNotFound()
		}
	}

	// Serve the file if we don't end up with any errors.
	http.ServeFile(c.Response, c.Request, path)

	return nil
}

func NewFileServer(y *yarf.Yarf, path, prefix string) *yarf.Yarf {

	// Set follow to file server
	y.Follow = http.StripPrefix(prefix, http.FileServer(CustomFileSystem{http.Dir(path)}))

	// Return without routes.
	return y
}
