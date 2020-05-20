package rest

import (
	"encoding/json"
	stor "github.com/da4nik/runroutes/internal/storage"
	"net/http"
)

func (s *Server) points(w http.ResponseWriter, _ *http.Request) {
	points := s.storage.Points()
	writeJSON(w, http.StatusOK, points)
}

func (s *Server) createPoint(w http.ResponseWriter, r *http.Request) {
	var point stor.Point
	err := json.NewDecoder(r.Body).Decode(&point)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	err = s.storage.CreatePoint(point)
	if err != nil {
		writeError(w, http.StatusUnprocessableEntity, err)
		return
	}

	writeJSON(w, http.StatusCreated, point)
}
