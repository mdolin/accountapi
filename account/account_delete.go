package account

import (
	"accountapi/client"
	"errors"
)

type RequestDelete struct {
	Host      string
	AccountID string
	Version   string
}

func AccountDelete(request *RequestDelete) (int, error) {
	if request.AccountID == "" {
		return 0, errors.New("Account ID required")
	}

	endpoint := request.Host + "/" + request.AccountID + request.Version
	// Create client
	client := client.CreateClient()

	// Get data
	status, err := client.Delete(endpoint)

	if err != nil {
		return 0, err
	}

	return status, err
}
