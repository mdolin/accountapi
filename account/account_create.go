package account

import (
	"accountapi/client"
	"accountapi/model"
	"encoding/json"
)

type RequestCreate struct {
	Host string
	Data *model.AccountData `json:"data"`
}

func AccountCreate(request *RequestCreate) (*model.AccountData, error) {
	// Crete data
	body, err := json.Marshal(request.Data)
	if err != nil {
		return nil, err
	}

	// Post data
	data, err := client.Post(request.Host, body)

	if err != nil {
		return nil, err
	}

	var response model.AccountData
	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return &response, err
}
