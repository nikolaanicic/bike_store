package handlers

import (
	centralstoreclient "bike_store/central_store_client"
	"bike_store/city_store/data"
	"bike_store/database/models"
	"bike_store/dto"
	"bike_store/handler"
	"net/http"
)

func rent(l *dto.RentBikeDto, db *data.Database) *dto.Status {

	var bike models.Bike

	err := db.Bikes.GetById(l.BikeID, &bike)
	if err != nil {
		return dto.NewStatus(http.StatusNotFound, "bike not found")
	}

	if !bike.Available {
		return dto.NewStatus(http.StatusConflict, "bike is not available")
	}

	status := centralstoreclient.Client.IncrementUserActiveBikes(l)

	if status.Code != http.StatusOK {
		return status
	}

	db.Bikes.GetById(l.BikeID, &bike)

	if !bike.Available {
		centralstoreclient.Client.DecrementUserActiveBikes(l)
		return dto.NewStatus(http.StatusConflict, "bike is not available")
	}

	rental := models.Rental{
		RentalBikeID:    l.BikeID,
		RentalCitizenID: l.CitizenID,
		City:            l.City,
	}
	bike.Available = false

	tx := db.DB.Begin()
	if err := tx.Create(&rental).Error; err != nil {
		tx.Rollback()
		return dto.NewStatus(http.StatusInternalServerError, "failed to create rental record")
	}
	if err := tx.Save(&bike).Error; err != nil {
		tx.Rollback()
		return dto.NewStatus(http.StatusInternalServerError, "failed to update bike availability")
	}
	if err := tx.Commit().Error; err != nil {
		return dto.NewStatus(http.StatusInternalServerError, "failed to commit transaction")
	}

	return dto.NewStatus(http.StatusOK, "bike successfully rented")

}

var RentBike = handler.AdaptHandlerWithDB(rent)
