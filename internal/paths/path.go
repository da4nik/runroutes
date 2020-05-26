package paths

import (
	"fmt"
	stor "github.com/da4nik/runroutes/internal/storage"
)

type Path struct {
	Distance float32
	Ways     []stor.Way
	Finished bool
}

func (p *Path) IsCircle() bool {
	if len(p.Ways) < 2 {
		return false
	}

	return p.Ways[0].From.ID == p.Ways[len(p.Ways)-1].To.ID
}

func (p *Path) ToString() string {
	if len(p.Ways) == 0 {
		return ""
	}

	res := fmt.Sprintf("[%0.2fkm] %s", p.Distance, p.Ways[0].From.ID)
	for _, w := range p.Ways {
		res += fmt.Sprintf(" >>(%0.2f)>> %s",
			w.Distance, w.To.ID)
	}
	return res
}

func (p *Path) AddWay(way stor.Way) Path {
	// appended way must start from the last point
	if p.Ways[len(p.Ways)-1].To.ID != way.From.ID {
		return *p
	}

	ways := make([]stor.Way, len(p.Ways))
	copy(ways, p.Ways)

	path := Path{
		Distance: p.Distance + way.Distance,
		Ways:     append(ways, way),
	}

	return path
}
