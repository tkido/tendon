// +build release

package assets

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/rakyll/statik/fs"
	_ "github.com/tkido/tendon/assets/statik"
)

var fileSystem http.FileSystem

func init() {
	var err error
	fileSystem, err = fs.New()
	if err != nil {
		log.Fatalf("statik FileSystem.New() failed!!")
	}
}

// Open assets
func Open(path string) (http.File, error) {
	p := filepath.Join("/", path)
	f, err := fileSystem.Open(p)
	if err != nil {
		return nil, err
	}
	return f, nil
}
