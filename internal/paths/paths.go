package paths

import (
	stor "github.com/da4nik/runroutes/internal/storage"
)

type AdjacencyMatrix map[string][]stor.Way

type Paths struct {
	points  []stor.Point
	ways    []stor.Way
	adjMart AdjacencyMatrix
}

func NewPaths(points []stor.Point, ways []stor.Way) *Paths {
	adjMatr := make(AdjacencyMatrix)
	for _, way := range ways {
		adjMatr[way.From.ID] = append(adjMatr[way.From.ID], way)
	}

	return &Paths{
		points:  points,
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

	resultPaths = append(resultPaths, Path{
		Distance: 0,
		Points:   []stor.Point{ways[0].From},
		Finished: false,
	})

	isFirst := true
	for {
		finishedCount := 0
		newPaths := make([]Path, 0)

		for i := range resultPaths {
			if resultPaths[i].Finished {
				finishedCount += 1
				newPaths = append(newPaths, resultPaths[i])
				continue
			}

			paths := p.getPaths(resultPaths[i], startPointID, isFirst)
			isFirst = false
			if paths == nil {
				resultPaths[i].Finished = true
				newPaths = append(newPaths, resultPaths[i])
				continue
			}

			newPaths = append(newPaths, paths...)
		}
		resultPaths = newPaths

		if finishedCount == len(resultPaths) {
			break
		}
	}

	return resultPaths
}

func (p *Paths) getPaths(path Path, startPointID string, isFirst bool) []Path {
	currentPoint := path.Points[len(path.Points)-1]

	ways := p.adjMart[currentPoint.ID]

	if len(ways) == 0 {
		return nil
	}

	if currentPoint.ID == startPointID && !isFirst {
		return nil
	}

	paths := make([]Path, 0)
	for i := range ways {
		newPath := path.AddWay(ways[i])
		paths = append(paths, newPath)
	}

	return paths
}
