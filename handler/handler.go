package handler

import (
	centralstoredata "bike_store/central_store/data"
	citystoredata "bike_store/city_store/data"

	"bike_store/database"
	"bike_store/dto"
	"net/http"
)

type Handler[T dto.JsonModel] func(T, database.IDatabase) *dto.Status

type DataBaseTypes interface {
	centralstoredata.Database | citystoredata.Database
}

func AdaptHandlerWithDB[T dto.JsonModel, DB DataBaseTypes](
	fn func(T, *DB) *dto.Status,
) Handler[T] {
	return func(model T, db database.IDatabase) *dto.Status {
		concreteDb, ok := any(db).(*DB)
		if !ok {
			return &dto.Status{
				Message: "Invalid database type",
				Code:    http.StatusInternalServerError,
			}
		}
		return fn(model, concreteDb)
	}
}
