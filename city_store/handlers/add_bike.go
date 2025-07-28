package handlers

import (
	"bike_store/city_store/data"
	localdto "bike_store/city_store/dto"
	"bike_store/database/models"
	"bike_store/dto"
	"bike_store/handler"
	"bike_store/log"
	"net/http"
	"strings"
)

func addBike(l *localdto.AddBikeDto, db *data.Database) *dto.Status {

	var newBike models.Bike

	newBikeType := strings.ToUpper(l.Type)
	newBikeCity := strings.ToUpper(l.City)

	if newBikeType != string(models.MountainBike) && newBikeType != string(models.CityBike) && newBikeType != string(models.RoadBike) {
		log.Info("bike type is invalid: %v", l)
		return dto.NewStatus(http.StatusBadRequest, "bike type is invalid")
	}

	if newBikeCity != string(models.KRAGUJEVAC) && newBikeCity != string(models.NOVI_SAD) && newBikeCity != string(models.SUBOTICA) {
		log.Info("city is invalid: %v", l)
		return dto.NewStatus(http.StatusBadRequest, "city is invalid")
	}

	newBike.Type = newBikeType
	newBike.City = newBikeCity
	newBike.Available = true

	if err := db.Bikes.Create(&newBike); err != nil {
		log.Error("failed to create a new bike: %v", err)
		return dto.NewStatus(http.StatusInternalServerError, "failed to create a new bike")
	}

	return dto.NewStatus(http.StatusOK, "bike successfully created")

}

var AddBike = handler.AdaptHandlerWithDB(addBike)
