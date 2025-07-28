package pipelines

import (
	"bike_store/city_store/handlers"
	"bike_store/database"
	"bike_store/pipeline"
)

func getAddBikePipeline(db database.IDatabase) pipeline.IPipeline {
	return pipeline.New("/bikes/add", pipeline.POST, handlers.AddBike, db)
}
