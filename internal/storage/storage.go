package storage

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/da4nik/runroutes/internal/log"
	"io/ioutil"
)

type db struct {
	Points []Point `json:"points"`
	Ways   []Way   `json:"ways"`
}

type Point struct {
	ID   string `json:"id"`
	Lat  string `json:"lat"`
	Long string `json:"long"`
}

type Way struct {
	ID             string  `json:"id"`
	From           Point   `json:"from"`
	To             Point   `json:"to"`
	Distance       float32 `json:"distance"`
	AsphaltPercent int     `json:"asphalt_percent"`
	Preferred      int     `json:"preferred"`
	Waypoints      []Point `json:"waypoints"`
}

type Storage struct {
	points []Point
	ways   []Way

	logger log.Logger
}

func NewStorage(logger log.Logger) *Storage {

	data, err := ioutil.ReadFile("db.json")
	if err != nil {
		logger.Errorf("Unable to read db: %s", err.Error())
	}

	var db db
	if err := json.Unmarshal(data, &db); err != nil {
		logger.Errorf("Unable unmarshal db: %s", err.Error())
	} else {
		logger.Infof("Loaded points: %d, ways: %d",
			len(db.Points), len(db.Ways))
	}

	return &Storage{
		points: db.Points,
		ways:   db.Ways,
		logger: logger,
	}
}

func (s *Storage) Points() []Point {
	return s.points
}

func (s *Storage) CreatePoint(point Point) error {
	point.ID = generatePointID(&point)
	s.points = append(s.points, point)

	s.saveDB()

	return nil
}

func (s *Storage) Ways() []Way {
	return s.ways
}

func (s *Storage) CreateWay(way Way) error {
	way.ID = generateWayID(&way)
	s.ways = append(s.ways, way)

	s.saveDB()

	return nil
}

func (s *Storage) saveDB() {
	data, err := json.MarshalIndent(db{
		Points: s.points,
		Ways:   s.ways,
	}, "", "    ")
	if err != nil {
		s.logger.Errorf("Error marshaling db: %s", err.Error())
		return
	}

	err = ioutil.WriteFile("db.json", data, 0644)
	if err != nil {
		s.logger.Errorf("Error saving db: %s", err.Error())
		return
	}
}

func generatePointID(point *Point) string {
	hasher := sha256.New()
	hasher.Write([]byte(fmt.Sprintf("%s%s", point.Lat, point.Long)))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func generateWayID(way *Way) string {
	hasher := sha256.New()
	hasher.Write([]byte(fmt.Sprintf("%s%s", way.From.ID, way.To.ID)))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
