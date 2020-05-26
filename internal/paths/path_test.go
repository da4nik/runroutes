package paths_test

import (
	"github.com/da4nik/runroutes/internal/paths"
	stor "github.com/da4nik/runroutes/internal/storage"
	"testing"
)

func newPath() paths.Path {
	return paths.Path{}
}

func existingPath() paths.Path {
	return paths.Path{
		Distance: 5,
		Points: []stor.Point{
			{
				ID: "1",
			},
			{
				ID: "2",
			},
		},
	}
}

func TestPathToString(t *testing.T) {
	t.Parallel()

	path := existingPath()

	pathString := path.ToString()
	if pathString == "12" {
		return
	}

	t.Errorf("Incorrect string, expected 12, got %s", pathString)
}
