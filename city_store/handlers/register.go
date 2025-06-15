package handlers

import (
	centralstoreclient "bike_store/central_store_client"
	"bike_store/city_store/data"
	"bike_store/dto"
	"bike_store/handler"
)

func register(l *dto.RegisterDto, _ *data.Database) *dto.Status {
	return centralstoreclient.Client.RegisterUser(l)
}

var Register = handler.AdaptHandlerWithDB(register)
