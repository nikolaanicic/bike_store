package pipelines

import (
	"bike_store/database"
	"bike_store/pipeline"
)

func GetPipelines(db database.IDatabase) []pipeline.IPipeline {
	return []pipeline.IPipeline{getRegisterPipeline(db), getRentPipeline(db), getReturnPipeline(db)}
}
