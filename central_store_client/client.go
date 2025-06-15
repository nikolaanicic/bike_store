package centralstoreclient

import (
	"bike_store/dto"
	"bike_store/log"
	"bike_store/pipeline"
	"bytes"
	"fmt"
	"net/http"
)

type CentralStoreClient struct {
	centralServerHost string
}

func New(centralStoreHost string) *CentralStoreClient {
	return &CentralStoreClient{
		centralServerHost: centralStoreHost,
	}
}

func (c *CentralStoreClient) getCentralServerMethodUrl(path string) string {
	return fmt.Sprintf("http://%s:%s", c.centralServerHost, path)
}

func retry(fn func(url string, data dto.JsonModel) *dto.Status, retries int, url string, data dto.JsonModel) *dto.Status {
	var status *dto.Status
	for i := 0; i < retries; i++ {
		status = fn(url, data)
		if status.Code == http.StatusOK {
			return status
		}
		log.Error("retry %d/%d failed with status: %s", i+1, retries, status.Message)
	}
	return status
}

func postWithRetry[T dto.JsonModel](url string, data T) *dto.Status {
	return retry(func(u string, d dto.JsonModel) *dto.Status {
		return postData(u, d.(T))
	}, 3, url, data)
}

func postData[T dto.JsonModel](url string, data T) *dto.Status {

	jsonData, err := data.AsJson()
	if err != nil {
		return dto.NewStatus(http.StatusInternalServerError, fmt.Sprintf("failed to serialize data: %v", err))
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return dto.NewStatus(http.StatusInternalServerError, fmt.Sprintf("failed to post data: %v", err))
	}
	defer response.Body.Close()

	status, failedDeserializationStatus := pipeline.ReadRequestFromBody[*dto.Status](response.Body)
	if failedDeserializationStatus != nil {
		return failedDeserializationStatus
	}

	status.Code = response.StatusCode
	return status

}

func (c *CentralStoreClient) IncrementUserActiveBikes(l *dto.RentBikeDto) *dto.Status {
	const rentBikeEndpoint = "/users/rent_bike"
	return postWithRetry(c.getCentralServerMethodUrl(rentBikeEndpoint), l)
}

func (c *CentralStoreClient) DecrementUserActiveBikes(l *dto.RentBikeDto) *dto.Status {
	const returnBikeEndpoint = "/users/return_bike"
	return postWithRetry(c.getCentralServerMethodUrl(returnBikeEndpoint), l)
}

func (c *CentralStoreClient) RegisterUser(l *dto.RegisterDto) *dto.Status {
	const registerEndpoint = "/users/register"
	return postWithRetry(c.getCentralServerMethodUrl(registerEndpoint), l)
}

var Client CentralStoreClient

func Configure(centralServerHost string) {
	Client = *New(centralServerHost)
	log.Info("Central Store Client configured with host: %s", centralServerHost)
}
