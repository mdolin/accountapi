package account

import (
	"accountapi/client"
	"errors"
	"log"
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

	// Create client
	client := client.CreateClient()

	// Delete data
	status, err := client.Delete(endpoint)

	if err != nil {
		return err
	}

	if status == 204 {
		log.Println(status, "Request succeeded")
	}

	return err
}
