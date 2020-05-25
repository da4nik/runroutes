package paths

import stor "github.com/da4nik/runroutes/internal/storage"

type Path struct {
	Distance float32
	Points   []stor.Point
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

func (p *Path) AddWay(way stor.Way) Path {
	// appended way must start from the last point
	if p.Points[len(p.Points)-1].ID != way.From.ID {
		return *p
	}

	points := make([]stor.Point, len(p.Points))
	copy(points, p.Points)

	path := Path{
		Distance: p.Distance + way.Distance,
		Points:   append(points, way.To),
	}

	return path
}

func (p *Path) Append(pth Path) {
	p.Distance += pth.Distance
	for _, point := range pth.Points {
		p.Points = append(p.Points, point)
	}
}
