package paths

import stor "github.com/da4nik/runroutes/internal/storage"

type Paths struct {
	points []stor.Point
	ways   []stor.Way
}

type Path struct {
	Distance float32
	Points   []stor.Point
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

func (p *Path) Append(pth Path) {
	p.Distance += pth.Distance
	for _, point := range pth.Points {
		p.Points = append(p.Points, point)
	}
}
