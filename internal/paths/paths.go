package paths

import (
	stor "github.com/da4nik/runroutes/internal/storage"
)

type AdjacencyMatrix map[string][]stor.Way

type Paths struct {
	ways    []stor.Way
	adjMart AdjacencyMatrix
}

func NewPaths(ways []stor.Way) *Paths {
	adjMatr := make(AdjacencyMatrix)
	for _, way := range ways {
		adjMatr[way.From.ID] = append(adjMatr[way.From.ID], way)
	}

	return &Paths{
		ways:    ways,
		adjMart: adjMatr,
	}
}

func (p *Paths) Calculate(startPointID string) []Path {
	var resultPaths []Path

	ways, exists := p.adjMart[startPointID]
	if !exists {
		return resultPaths
	}

	if len(ways) == 0 {
		return resultPaths
	}

	resultPaths = p.searchPaths(nil, startPointID)

	for {
		finished := 0
		for _, path := range resultPaths {
			if path.Finished {
				finished += 1
			}
		}
		if finished == len(resultPaths) {
			break
		}

		resultPaths = p.searchPaths(resultPaths, startPointID)
	}

	return resultPaths
}

func (p *Paths) searchPaths(paths []Path, startPointID string) []Path {
	if paths == nil {
		return p.getPaths(nil, startPointID)
	}

	var resultPaths []Path
	for i := range paths {
		workingPath := paths[i]

		if workingPath.Finished {
			resultPaths = append(resultPaths, workingPath)
			continue
		}

		paths := p.getPaths(&workingPath, startPointID)
		if paths == nil {
			workingPath.Finished = true
			resultPaths = append(resultPaths, workingPath)
			continue
		}

		resultPaths = append(resultPaths, paths...)
	}

	return resultPaths
}

func (p *Paths) getPaths(path *Path, startPointID string) []Path {
	var currentPointID string
	if path == nil {
		currentPointID = startPointID
	} else {
		currentPointID = path.Ways[len(path.Ways)-1].To.ID
	}

	ways := p.adjMart[currentPointID]

	if len(ways) == 0 {
		return nil
	}

	if currentPointID == startPointID && path != nil {
		return nil
	}

	paths := make([]Path, 0)
	for i := range ways {
		var newPath Path
		if path == nil {
			newPath = Path{
				Distance: ways[i].Distance,
				Ways:     []stor.Way{ways[i]},
				Finished: false,
			}
		} else {
			newPath = path.AddWay(ways[i])
		}

		paths = append(paths, newPath)
	}

	return paths
}
