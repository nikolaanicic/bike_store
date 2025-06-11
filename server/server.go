package server

import "net/http"

type Server struct {
	host string
	mux  *http.ServeMux
}
