package paths_test

import (
	"github.com/da4nik/runroutes/internal/paths"
	stor "github.com/da4nik/runroutes/internal/storage"
	"testing"
)

func existingPath() paths.Path {
	return paths.Path{
		Distance: 5,
		Ways: []stor.Way{
			{
				From:      stor.Point{ID: "1"},
				To:        stor.Point{ID: "2"},
				Distance:  0,
				Waypoints: nil,
			},
		},
	}
}

func TestPathID(t *testing.T) {
	t.Parallel()

	path := existingPath()

	pathString := path.ID()
	if pathString == "12" {
		return
	}

	t.Errorf("Incorrect string, expected 12, got %s", pathString)
}
