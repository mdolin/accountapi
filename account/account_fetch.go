package account

import (
	"accountapi/client"
	"accountapi/model"
	"encoding/json"
	"errors"
)

type RequestFetch struct {
	Host      string
	AccountID string
}

func AccountFetch(request *RequestFetch) (*model.AccountData, error) {
	if request.AccountID == "" {
		return nil, errors.New("Account ID required")
	}

	endpoint := request.Host + "/" + request.AccountID

	// Get data
	data, err := client.Get(endpoint)

	if err != nil {
		return nil, err
	}

	var response model.AccountData
	er := json.Unmarshal(data, &response)

	if er != nil {
		return nil, er
	}

	return &response, err
}
