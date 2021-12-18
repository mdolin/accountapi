package account

import (
	"accountapi/client"
	"accountapi/model"
	"encoding/json"
	"errors"
)

type RequestCreate struct {
	Host string
	Data []byte
}

func AccountCreate(request *RequestCreate) (*model.AccountData, error) {
	if len(request.Data) == 0 {
		return nil, errors.New("Data required")
	}

	// Crete data
	data, err := client.Post(request.Host, request.Data)

	if err != nil {
		return nil, err
	}

	var response model.AccountData
	json.Unmarshal(data, &response)

	return &response, err
}
