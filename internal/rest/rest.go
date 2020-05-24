package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/da4nik/runroutes/internal/log"
	stor "github.com/da4nik/runroutes/internal/storage"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

type m map[string]interface{}

type Server struct {
	port       int
	storage    *stor.Storage
	router     *chi.Mux
	logger     log.Logger
	httpServer *http.Server
}

func NewServer(port int, storage *stor.Storage, logger log.Logger) *Server {
	return &Server{
		port:    port,
		storage: storage,
		router:  chi.NewRouter(),
		logger:  logger,
	}
}

func (s *Server) Start() error {
	s.logger.Debugf("Starting...")
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Heartbeat("/heartbeat"))

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	s.router.Use(corsMiddleware.Handler)

	s.router.Route("/api/v1", func(r chi.Router) {
		r.Route("/points", func(r chi.Router) {
			r.Get("/", s.points)
			r.Post("/", s.createPoint)
		})

		r.Route("/ways", func(r chi.Router) {
			r.Get("/", s.notImplemented)
			r.Post("/", s.notImplemented)
		})
	})

	port := s.port
	if port == 0 {
		port = 3333
	}

	s.httpServer = &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		Handler:   s.router,
		TLSConfig: nil,
	}

	go func() {
		s.logger.Infof("REST server running at port %s",
			s.httpServer.Addr)

		if err := s.httpServer.ListenAndServe(); err != nil {
			s.logger.Errorf("Failed to start rest server: %s",
				err.Error())
		}
	}()

	s.logger.Debugf("Started")
	return nil
}

func (s *Server) Stop() {
	s.logger.Debugf("Stopping...")
	err := s.httpServer.Shutdown(context.Background())
	if err != nil {
		s.logger.Errorf("Failed to stop rest server: %s", err.Error())
	}
	s.logger.Debugf("Stopped")
}

func (s *Server) notImplemented(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func writeJSON(w http.ResponseWriter, status int, body interface{}) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	_, _ = w.Write(jsonBody)
}

func writeError(w http.ResponseWriter, status int, err error) {
	writeJSON(w, status, m{"error": err.Error()})
}
