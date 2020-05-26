package paths

import (
	"fmt"
	stor "github.com/da4nik/runroutes/internal/storage"
)

type Path struct {
	Distance float32
	Points   []stor.Point
	Ways     []stor.Way
	Finished bool
}

func (p *Path) IsCircle() bool {
	if len(p.Points) < 2 {
		return false
	}

	return p.Points[0].ID == p.Points[len(p.Points)-1].ID
}

func (p *Path) ToString() string {
	if len(p.Points) == 0 {
		return ""
	}

	res := ""
	for _, p := range p.Points {
		res += p.ID
	}
	return res
}

func (p *Path) WaysToString() string {
	if len(p.Ways) == 0 {
		return ""
	}

	res := p.Ways[0].From.ID

	for _, w := range p.Ways {
		res += fmt.Sprintf(" >>(%0.2f)>> %s",
			w.Distance, w.To.ID)
	}
	return res
}

func (p *Path) AddWay(way stor.Way) Path {
	// appended way must start from the last point
	if p.Points[len(p.Points)-1].ID != way.From.ID {
		return *p
	}

	points := make([]stor.Point, len(p.Points))
	copy(points, p.Points)

	ways := make([]stor.Way, len(p.Ways))
	copy(ways, p.Ways)

	path := Path{
		Distance: p.Distance + way.Distance,
		Points:   append(points, way.To),
		Ways:     append(ways, way),
	}

	return path
}
