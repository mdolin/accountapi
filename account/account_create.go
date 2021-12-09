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
	Data string
}

func AccountCreate(request *RequestCreate) (*model.AccountData, error) {
	if request.Data == "" {
		return nil, errors.New("Data required")
	}

	// Convert to byte
	var res map[string]interface{}

	json.Unmarshal([]byte(request.Data), &res)
	body, err := json.Marshal(res)

	// Create client
	client := client.CreateClient()

	// Crete data
	status, data, err := client.Post(request.Host, body)

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
