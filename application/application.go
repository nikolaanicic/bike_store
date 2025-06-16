package application

import (
	"bike_store/configuration"
	"bike_store/database"
	"bike_store/log"
	"bike_store/pipeline"
	"bike_store/server"
)

type Application struct {
	server        *server.Server
	database      database.IDatabase
	Configuration *configuration.Config
}

func New(db database.IDatabase) *Application {

	cfg, err := configuration.Get()

	if err != nil {
		log.Fatalf("failed to read config...shuting down: %v", err)
	}

	if err := db.Configure(&cfg.Database); err != nil {
		log.Fatalf("failed to configure the database connection: %v", err)
	}

	server := server.New(cfg.Server)

	return &Application{database: db, Configuration: &cfg, server: server}
}

func (a *Application) SetPipelines(pipelines ...pipeline.IPipeline) *Application {
	for _, p := range pipelines {
		a.server.RegisterPipelines(p)
		log.Info("%s %s", p.Type(), p.Path())
	}

	return a
}

func (a *Application) Start() {
	a.server.Serve()
}
