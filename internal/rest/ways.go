package rest

import (
	"encoding/json"
	stor "github.com/da4nik/runroutes/internal/storage"
	"net/http"
)

func (s *Server) ways(w http.ResponseWriter, _ *http.Request) {
	points := s.storage.Ways()
	writeJSON(w, http.StatusOK, points)
}

func (s *Server) createWay(w http.ResponseWriter, r *http.Request) {
	var way stor.Way
	err := json.NewDecoder(r.Body).Decode(&way)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	err = s.storage.CreateWay(way)
	if err != nil {
		writeError(w, http.StatusUnprocessableEntity, err)
		return
	}

	writeJSON(w, http.StatusCreated, way)
}
