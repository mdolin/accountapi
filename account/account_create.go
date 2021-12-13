package account

import (
	"accountapi/client"
	"accountapi/model"
	"encoding/json"
	"errors"
	"log"
)

type RequestCreate struct {
	Host string
	Data []byte
}

func AccountCreate(request *RequestCreate) (*model.AccountData, error) {
	if len(request.Data) == 0 {
		return nil, errors.New("Data required")
	}

	// Create client
	client := client.CreateClient()

	// Crete data
	status, data, err := client.Post(request.Host, request.Data)

	if err != nil {
		return nil, err
	}

	if status == 201 {
		log.Println(status, "Create succeeded")
	}

	var response model.AccountData
	json.Unmarshal(data, &response)

	return &response, err
}
