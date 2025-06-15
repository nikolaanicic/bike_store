package pipeline

import (
	"bike_store/dto"
	"encoding/json"
	"io"
	"net/http"
)

func ReadRequestFromBody[T dto.JsonModel](body io.ReadCloser) (T, *dto.Status) {
	var t T

	bodyData, err := io.ReadAll(body)
	if err != nil {
		return t, dto.NewStatus(http.StatusBadRequest, "Bad Request")
	}

	if err := json.Unmarshal(bodyData, &t); err != nil {
		return t, dto.NewStatus(http.StatusBadRequest, "Invalid Json")
	}

	return t, nil
}

func packResponse(w http.ResponseWriter, response *dto.Status) {

	data, _ := json.Marshal(map[string]interface{}{"message": response.Message})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)
	w.Write(data)
}
