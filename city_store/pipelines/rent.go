package pipelines

import (
	"bike_store/city_store/handlers"
	"bike_store/database"
	"bike_store/pipeline"
)

func getRentPipeline(db database.IDatabase) pipeline.IPipeline {
	return pipeline.New("/users/rent_bike", pipeline.POST, handlers.RentBike, db)
}
