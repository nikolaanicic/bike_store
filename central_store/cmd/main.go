package main

import (
	"bike_store/application"
	"bike_store/central_store/data"
	"bike_store/central_store/pipelines"
)

func main() {
	db := data.NewDB()
	application.New(db).SetPipelines(pipelines.GetPipelines(db)...).Start()
}
