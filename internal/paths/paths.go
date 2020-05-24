package paths

import (
	"fmt"
	stor "github.com/da4nik/runroutes/internal/storage"
)

type AdjacencyMatrix map[string][]stor.Way

func NewPaths(points []stor.Point, ways []stor.Way) *Paths {
	return &Paths{
		points: points,
		ways:   ways,
	}
}

func (p *Paths) Calculate(startPointID string) []Path {
	if len(p.ways) == 0 {
		return nil
	}

	var result []Path

	adjMart := make(AdjacencyMatrix)
	for _, way := range p.ways {
		adjMart[way.From.ID] = append(adjMart[way.From.ID], way)
	}

	start, exists := adjMart[startPointID]
	if !exists || len(start) == 0 {
		fmt.Printf("Starting point has no ways")
		return nil
	}

	currentWay := start[0]
	currentPoint := currentWay.From

	resPath := Path{
		Distance: 0,
		Points:   []stor.Point{currentPoint},
	}

	for {
		nextArr := adjMart[currentPoint.ID]
		if len(nextArr) == 0 {
			break
		}

		nextWay := nextArr[0]
		resPath.Points = append(resPath.Points, nextWay.To)
		resPath.Distance += nextWay.Distance

		if nextWay.To.ID == startPointID {
			break
		}

		currentPoint = nextWay.To
	}

	result = append(result, resPath)

	//fmt.Printf("RESULT: %+v\n", result)
	return result
}
