package handlers

import (
	"bike_store/central_store/data"
	"bike_store/database/models"
	"bike_store/dto"
	"bike_store/handler"
	"net/http"
)

func rent(l *dto.RentBikeDto, db *data.Database) *dto.Status {
	var user models.User

	err := db.Users.GetById(l.CitizenID, &user)
	if err == nil && user.CitizenID == l.CitizenID {
		return dto.NewStatus(http.StatusNotFound, "user doesn't exist")
	}

	if user.ActiveBikes >= 2 {
		return dto.NewStatus(http.StatusConflict, "user has already rented 2 bikes")
	}

	user.ActiveBikes++
	if err := db.Users.Update(&user); err != nil {
		return dto.NewStatus(http.StatusInternalServerError, "failed to update user active bikes")
	}

	return dto.NewStatus(http.StatusOK, "user successfully rented a bike")
}

var RentBike = handler.AdaptHandlerWithDB(rent)
