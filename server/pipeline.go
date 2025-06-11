package server

import (
	"net/http"
)

const (
	POST = http.MethodPost
	GET  = http.MethodGet
)

type Pipeline[T JsonModel] struct {
	path string
	// to add handlers
	// middleware
	// and the database interface
}
