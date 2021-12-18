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

func AccountDelete(request *RequestDelete) error {
	if request.AccountID == "" {
		return errors.New("Account ID required")
	}

	endpoint := request.Host + "/" + request.AccountID + request.Version

	// Delete data
	err := client.Delete(endpoint)

	if err != nil {
		return err
	}

	return err
}
