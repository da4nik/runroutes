package paths_test

import (
	"github.com/da4nik/runroutes/internal/paths"
	"github.com/da4nik/runroutes/internal/storage"
	"testing"
)

func newEmptyPaths() *paths.Paths {
	return paths.NewPaths(nil)
}

func newOneRingPaths() *paths.Paths {
	return paths.NewPaths([]storage.Way{
		{
			From:     storage.Point{ID: "1"},
			To:       storage.Point{ID: "2"},
			Distance: 1,
		},
		{
			From:     storage.Point{ID: "2"},
			To:       storage.Point{ID: "3"},
			Distance: 1,
		},
		{
			From:     storage.Point{ID: "3"},
			To:       storage.Point{ID: "4"},
			Distance: 1,
		},
		{
			From:     storage.Point{ID: "4"},
			To:       storage.Point{ID: "1"},
			Distance: 1,
		},
	})
}

func newTwoRingsPaths() *paths.Paths {
	return paths.NewPaths([]storage.Way{
		{
			From:     storage.Point{ID: "1"},
			To:       storage.Point{ID: "2"},
			Distance: 1,
		},
		{
			From:     storage.Point{ID: "2"},
			To:       storage.Point{ID: "3"},
			Distance: 1,
		},
		{
			From:     storage.Point{ID: "3"},
			To:       storage.Point{ID: "4"},
			Distance: 1,
		},
		{
			From:     storage.Point{ID: "4"},
			To:       storage.Point{ID: "1"},
			Distance: 1,
		},
		{
			From:     storage.Point{ID: "3"},
			To:       storage.Point{ID: "5"},
			Distance: 1,
		},
		{
			From:     storage.Point{ID: "5"},
			To:       storage.Point{ID: "6"},
			Distance: 1,
		},
		{
			From:     storage.Point{ID: "6"},
			To:       storage.Point{ID: "1"},
			Distance: 1,
		},
	})
}

func TestEmptyWays(t *testing.T) {
	t.Parallel()

	p := newEmptyPaths()

	if p.Calculate("1", 1) != nil {
		t.Error("Paths should be empty")
		return
	}
}

func TestOneRingAndDistance(t *testing.T) {
	t.Parallel()

	p := newOneRingPaths()

	result := p.Calculate("1", 1)
	if len(result) != 1 {
		t.Errorf("Should be exactly one path, got %d", len(result))
		return
	}

	path := result[0]

	if path.ID() != "12341" {
		t.Errorf("Wrong path, expected 12341 got %s", path.ID())
	}

	if path.Distance != 4 {
		t.Errorf("Distance should be 4, got %f", path.Distance)
		return
	}

	//fmt.Printf("Paths: %+v\n", result)
}

func TestTwoRingAndDistance(t *testing.T) {
	t.Parallel()

	p := newTwoRingsPaths()

	result := p.Calculate("1", 1)
	if len(result) != 2 {
		t.Errorf("Should be exactly two paths, got %d", len(result))
		return
	}

	pathOne := result[0]
	pathTwo := result[1]

	if pathOne.ID() != "12341" && pathTwo.ID() != "12341" {
		t.Errorf("Wrong path, expected 12341 got %s and %s",
			pathOne.ID(), pathTwo.ID())
	}

	if pathOne.ID() != "123561" && pathTwo.ID() != "123561" {
		t.Errorf("Wrong path, expected 123561 got %s and %s",
			pathOne.ID(), pathTwo.ID())
	}

	if pathOne.Distance != 4 && pathTwo.Distance != 4 {
		t.Errorf("Distance should be 4, got %f and %f",
			pathOne.Distance, pathTwo.Distance)
		return
	}

	if pathOne.Distance != 5 && pathTwo.Distance != 5 {
		t.Errorf("Distance should be 5, got %f and %f",
			pathOne.Distance, pathTwo.Distance)
		return
	}

	//fmt.Printf("Paths: %+v\n", result)
}
