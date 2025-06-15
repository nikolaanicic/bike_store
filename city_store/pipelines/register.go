package pipelines

import (
	"bike_store/city_store/handlers"
	"bike_store/database"
	"bike_store/pipeline"
)

func getRegisterPipeline(db database.IDatabase) pipeline.IPipeline {
	return pipeline.New("/users/register", pipeline.POST, handlers.Register, db)
}
