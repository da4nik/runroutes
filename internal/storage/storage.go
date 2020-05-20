package storage

import "github.com/da4nik/runroutes/internal/log"

type Point struct {
	ID             string `json:"id"`
	Lat            string `json:"lat"`
	Long           string `json:"long"`
	AsphaltPercent int    `json:"asphalt_percent"`
	Preferred      int    `json:"preferred"`
}

type Way struct {
	From      Point   `json:"from"`
	To        Point   `json:"to"`
	Distance  float32 `json:"distance"`
	Waypoints []Point `json:"waypoints"`
}

type Storage struct {
	points []Point
	ways   []Way

	logger log.Logger
}

func NewStorage(logger log.Logger) *Storage {
	return &Storage{
		points: []Point{},
		ways:   []Way{},
		logger: logger,
	}
}

func (s *Storage) Points() []Point {
	return s.points
}

func (s *Storage) CreatePoint(point Point) error {
	s.points = append(s.points, point)
	return nil
}

func (s *Storage) Ways() []Way {
	return s.ways
}
