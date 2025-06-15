package pipelines

import (
	"bike_store/central_store/handlers"
	"bike_store/database"
	"bike_store/pipeline"
)

func getReturnPipeline(db database.IDatabase) pipeline.IPipeline {
	return pipeline.New("/users/return_bike", pipeline.POST, handlers.ReturnBike, db)
}
