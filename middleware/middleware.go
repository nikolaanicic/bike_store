package middleware

import "bike_store/dto"

type Middleware func(dto.JsonModel) *dto.Status
