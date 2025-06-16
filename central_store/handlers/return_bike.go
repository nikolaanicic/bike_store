package handlers

import (
	"bike_store/central_store/data"
	"bike_store/database/models"
	"bike_store/dto"
	"bike_store/handler"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

func return_bike(l *dto.RentBikeDto, db *data.Database) *dto.Status {
	var user models.User

	err := db.Users.GetById(l.CitizenID, &user)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.NewStatus(http.StatusNotFound, "user doesn't exist")
	}

	if user.ActiveBikes == 0 {
		return dto.NewStatus(http.StatusConflict, "user doesn't have any rented bikes")
	}

	user.ActiveBikes--
	if err := db.Users.Update(&user); err != nil {
		return dto.NewStatus(http.StatusInternalServerError, "failed to update user active bikes")
	}

	return dto.NewStatus(http.StatusOK, "user successfully returned a bike")
}

var ReturnBike = handler.AdaptHandlerWithDB(return_bike)
