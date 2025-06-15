package handlers

import (
	"bike_store/central_store/data"
	"bike_store/database/models"
	"bike_store/dto"
	"bike_store/handler"
	"bike_store/log"
	"net/http"
)

func register(l *dto.RegisterDto, db *data.Database) *dto.Status {
	var user models.User

	err := db.Users.GetById(l.CitizenID, &user)
	if err == nil && user.CitizenID == l.CitizenID {
		return dto.NewStatus(http.StatusConflict, "user already exists")
	}

	user = models.User{Name: l.Name, LastName: l.LastName, Address: l.Address, CitizenID: l.CitizenID}

	if err := db.Users.Create(&user); err != nil {
		log.Error("failed to create the user: %v", err)
		return dto.NewStatus(http.StatusInternalServerError, err.Error())
	}

	return dto.NewStatus(http.StatusOK, "user successfully created")
}

var Register = handler.AdaptHandlerWithDB(register)
