package handlers

import (
	centralstoreclient "bike_store/central_store_client"
	"bike_store/city_store/data"
	"bike_store/database/models"
	"bike_store/dto"
	"bike_store/handler"
	"net/http"
	"time"
)

func return_bike(l *dto.RentBikeDto, db *data.Database) *dto.Status {

	var bike models.Bike

	err := db.Bikes.GetById(l.BikeID, &bike)
	if err != nil {
		return dto.NewStatus(http.StatusNotFound, "bike not found")
	}

	if bike.Available {
		return dto.NewStatus(http.StatusConflict, "bike is already available")
	}

	status := centralstoreclient.Client.DecrementUserActiveBikes(l)

	if status.Code != http.StatusOK {
		return status
	}

	rental, err := db.Rentals.GetByUserAndBike(l.CitizenID, l.BikeID)
	if err != nil {
		return dto.NewStatus(http.StatusNotFound, "rental not found")
	}

	rental.Completed = true
	rental.ReturnedDate = time.Now()
	bike.Available = true

	tx := db.DB.Begin()

	if err := tx.Save(&rental).Error; err != nil {
		tx.Rollback()
		return dto.NewStatus(http.StatusInternalServerError, "failed to update rental status")
	}

	if err := tx.Save(bike).Error; err != nil {
		tx.Rollback()
		return dto.NewStatus(http.StatusInternalServerError, "failed to update bike availability")
	}

	if err := tx.Commit().Error; err != nil {
		return dto.NewStatus(http.StatusInternalServerError, "failed to commit transaction")
	}

	return status
}

var ReturnBike = handler.AdaptHandlerWithDB(return_bike)
