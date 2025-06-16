package middleware

import (
	"bike_store/dto"
	"net/http"
)

type Middleware func(*http.Request) *dto.Status
