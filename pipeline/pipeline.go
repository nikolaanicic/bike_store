package pipeline

import (
	"bike_store/database"
	"bike_store/dto"
	"bike_store/handler"
	"bike_store/log"
	"bike_store/middleware"
	"net/http"
)

const (
	POST = http.MethodPost
	GET  = http.MethodGet
)

type Pipeline[T dto.JsonModel] struct {
	path        string
	requestType string
	handler     handler.Handler[T]
	middleware  []middleware.Middleware
	database    database.IDatabase
}

type IPipeline interface {
	Execute(http.ResponseWriter, *http.Request)
	Path() string
	Type() string
}

func New[T dto.JsonModel](path string, requestType string, handler handler.Handler[T], db database.IDatabase, middleware ...middleware.Middleware) *Pipeline[T] {
	return &Pipeline[T]{path: path, requestType: requestType, handler: handler, database: db, middleware: middleware}
}

func (p *Pipeline[T]) executeMiddleware(r dto.JsonModel) *dto.Status {
	log.Info("Executing middleware for pipeline %s", p.path)
	for _, middleware := range p.middleware {
		if err := middleware(r); err != nil {
			log.Error("Error executing middleware for pipeline %s: %v", p.path, err)
			return err
		}
	}
	return nil
}

func (p *Pipeline[T]) Execute(w http.ResponseWriter, r *http.Request) {

	if r.Method != p.requestType {
		log.Error("method not allowed: %s", r.Method)
		packResponse(w, dto.NewStatus(http.StatusMethodNotAllowed, "Not Allowed"))
		return
	}

	request, err := ReadRequestFromBody[T](r.Body)

	if err != nil {
		log.Error("Error reading request: %v", err)
		packResponse(w, err)
		return
	}

	defer r.Body.Close()

	if err := p.executeMiddleware(request); err != nil {
		packResponse(w, err)
		return
	}

	log.Info("handling: %s", p.path)
	status := p.handler(request, p.database)
	log.Info("returning status: %v", status)

	packResponse(w, status)
}

func (p *Pipeline[T]) Path() string {
	return p.path
}

func (p *Pipeline[T]) Type() string {
	return p.requestType
}
