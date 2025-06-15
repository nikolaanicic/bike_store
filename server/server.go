package server

import (
	"bike_store/configuration"
	"bike_store/pipeline"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	host string
	mux  *http.ServeMux
}

func GetCentralServerMethodUrl(path string) string {
	return fmt.Sprintf("http://%s:%s", configuration.CentralServerHost, path)
}

func New(cfg configuration.Server) *Server {
	return &Server{host: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), mux: http.NewServeMux()}
}

func (s *Server) RegisterPipelines(pipelines ...pipeline.IPipeline) {
	for _, p := range pipelines {
		s.mux.HandleFunc(p.Path(), p.Execute)
	}
}

func (s *Server) Serve() {
	if err := http.ListenAndServe(s.host, s.mux); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
