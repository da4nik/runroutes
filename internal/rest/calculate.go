package rest

import (
	"encoding/json"
	"github.com/da4nik/runroutes/internal/paths"
	stor "github.com/da4nik/runroutes/internal/storage"
	"net/http"
)

type CalculateRequest struct {
	Ways          []stor.Way `json:"ways"`
	StartingPoint stor.Point `json:"starting_point"`
	Distance      float64    `json:"distance"`
}

func (s *Server) calculate(w http.ResponseWriter, r *http.Request) {
	var request CalculateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	calculator := paths.NewPaths(request.Ways)
	result := calculator.Calculate(request.StartingPoint.ID, request.Distance)

	writeJSON(w, http.StatusOK, result)
}
