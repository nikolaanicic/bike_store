package main

import (
	"bike_store/application"
	centralstoreclient "bike_store/central_store_client"
	"bike_store/city_store/data"
	"bike_store/city_store/pipelines"
)

func main() {
	db := data.NewDB()
	app := application.New(db).SetPipelines(pipelines.GetPipelines(db)...)
	centralstoreclient.Configure(app.Configuration.CentralServerHost)

	app.Start()

}
