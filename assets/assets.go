// +build !release

package assets

import (
	"net/http"
	"os"
	"path/filepath"
)

// Open assets
func Open(path string) (http.File, error) {
	p := filepath.Join("_assets", path)
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	return f, nil
}
