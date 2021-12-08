package account

import (
	"accountapi/client"
	"accountapi/model"
	"encoding/json"
	"errors"
)

type Request struct {
	Host      string
	AccountID string
}

func AccountFetch(request *Request) (*model.AccountData, error) {
	if request.AccountID == "" {
		return nil, errors.New("Account ID required")
	}

	endpoint := request.Host + "/" + request.AccountID
	// Create client
	client := client.CreateClient()

	// Get data
	byteData, err := client.Get(endpoint)

	if err != nil {
		return nil, err
	}

	// var response Response
	// json.Unmarshal(byteData, &response)
	var response model.AccountData
	json.Unmarshal(byteData, &response)

	return &response, err
}
