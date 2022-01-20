package integration

import (
	"accountapi/account"
	"accountapi/model"
	"encoding/json"
	"testing"
)

const URL = "http://accountapi:8080/"

func AccountCreate(data string) (*model.AccountData, error) {
	var newAccount model.AccountData
	json.Unmarshal([]byte(data), &newAccount)

	var request account.RequestCreate
	request.Host = URL + "v1/organisation/accounts"
	request.Data = &newAccount

	resp, err := account.AccountCreate(&request)

	return resp, err
}

func AccountDelete(id string) error {
	var request account.RequestDelete
	request.Host = URL + "v1/organisation/accounts"
	request.AccountID = id
	request.Version = "?version=0"

	err := account.AccountDelete(&request)

	return err
}

func TestAccountCreateSuccessful(t *testing.T) {
	account_data := `{
		"data": {
		  "id": "123e4567-e89b-12d3-a456-426614174123",
		  "organisation_id": "123e4567-e89b-12d3-a456-426614174123",
		  "type": "accounts",
		  "attributes": {
			 "country": "GB",
			  "base_currency": "GBP",
			  "bank_id": "400302",
			  "bank_id_code": "GBDSC",
			  "account_number": "10000004",
			  "iban": "GB28NWBK40030212764204",
			  "bic": "NWBKGB42",
			  "account_classification": "Personal",
			  "name": ["TestName"]
		  }
		}
	  }`

	resp, err := AccountCreate(account_data)

	if err != nil {
		t.Errorf("Unexpected error: %v", err.Error())
	}

	var account model.AccountData
	json.Unmarshal([]byte(account_data), &account)

	// Test Key-Values
	if resp.Data.ID != account.Data.ID {
		t.Errorf("Wrong ID: got %v expected %v", resp.Data.ID, account.Data.ID)
	}
	if resp.Data.OrganisationID != account.Data.OrganisationID {
		t.Errorf("Wrong OrganisationID: got %v expected %v", resp.Data.OrganisationID, account.Data.OrganisationID)
	}
	if resp.Data.Type != account.Data.Type {
		t.Errorf("Wrong Type: got %v expected %v", resp.Data.Type, account.Data.Type)
	}
	if *resp.Data.Attributes.Country != *account.Data.Attributes.Country {
		t.Errorf("Wrong Country: got %v expected %v", *resp.Data.Attributes.Country, *account.Data.Attributes.Country)
	}
	if resp.Data.Attributes.BaseCurrency != account.Data.Attributes.BaseCurrency {
		t.Errorf("Wrong BaseCurrency: got %v expected %v", resp.Data.Attributes.BaseCurrency, account.Data.Attributes.BaseCurrency)
	}
	if resp.Data.Attributes.BankID != account.Data.Attributes.BankID {
		t.Errorf("Wrong BankID: got %v expected %v", resp.Data.Attributes.BankID, account.Data.Attributes.BankID)
	}
	if resp.Data.Attributes.BankIDCode != account.Data.Attributes.BankIDCode {
		t.Errorf("Wrong BankIDCode: got %v expected %v", resp.Data.Attributes.BankIDCode, account.Data.Attributes.BankIDCode)
	}
	if resp.Data.Attributes.AccountNumber != account.Data.Attributes.AccountNumber {
		t.Errorf("Wrong AccountNumber: got %v expected %v", resp.Data.Attributes.AccountNumber, account.Data.Attributes.AccountNumber)
	}
	if resp.Data.Attributes.Iban != account.Data.Attributes.Iban {
		t.Errorf("Wrong Iban: got %v expected %v", resp.Data.Attributes.Iban, account.Data.Attributes.Iban)
	}
	if resp.Data.Attributes.Bic != account.Data.Attributes.Bic {
		t.Errorf("Wrong Bic: got %v expected %v", resp.Data.Attributes.Bic, account.Data.Attributes.Bic)
	}
	if *resp.Data.Attributes.AccountClassification != *account.Data.Attributes.AccountClassification {
		t.Errorf("Wrong AccountClassification: got %v expected %v", *resp.Data.Attributes.AccountClassification, *account.Data.Attributes.AccountClassification)
	}
	if len(resp.Data.Attributes.Name) == len(account.Data.Attributes.Name) {
		for i := 0; i < len(resp.Data.Attributes.Name); i++ {
			if resp.Data.Attributes.Name[i] != account.Data.Attributes.Name[i] {
				t.Errorf("Wrong Name: got %v expected %v", resp.Data.Attributes.Name[i], account.Data.Attributes.Name[i])
			}
		}
	} else {
		t.Errorf("Name array is not the same")
	}
}

func TestAccountCreateAlreadyExist(t *testing.T) {

	account_data := `{
		"data": {
		  "id": "123e4567-e89b-12d3-a456-426614174123",
		  "organisation_id": "123e4567-e89b-12d3-a456-426614174123",
		  "type": "accounts",
		  "attributes": {
			 "country": "GB",
			  "base_currency": "GBP",
			  "bank_id": "400302",
			  "bank_id_code": "GBDSC",
			  "account_number": "10000004",
			  "iban": "GB28NWBK40030212764204",
			  "bic": "NWBKGB42",
			  "account_classification": "Personal",
			  "name": ["TestName"]
		  }
		}
	  }`

	resp, err := AccountCreate(account_data)

	if resp != nil {
		t.Errorf("Existing Account ID")
	}

	if err != nil && err.Error() != "409 Conflict" {
		t.Errorf("Unexpected error, got %v expected %v", err.Error(), "409 Conflict")
	}

	AccountDelete("123e4567-e89b-12d3-a456-426614174123")
}

func TestAccountFetchSuccessful(t *testing.T) {
	account_data := `{
		"data": {
		  "id": "123e4567-e89b-12d3-a456-426614174123",
		  "organisation_id": "123e4567-e89b-12d3-a456-426614174123",
		  "type": "accounts",
		  "attributes": {
			 "country": "GB",
			  "base_currency": "GBP",
			  "bank_id": "400302",
			  "bank_id_code": "GBDSC",
			  "account_number": "10000004",
			  "iban": "GB28NWBK40030212764204",
			  "bic": "NWBKGB42",
			  "account_classification": "Personal",
			  "name": ["TestName"]
		  }
		}
	  }`

	_, e := AccountCreate(account_data)
	if e != nil {
		t.Errorf("Unexpected error")
	}

	var request account.RequestFetch
	request.Host = URL + "v1/organisation/accounts"
	request.AccountID = "123e4567-e89b-12d3-a456-426614174123"

	resp, err := account.AccountFetch(&request)

	if err != nil {
		t.Errorf("Unexpected error: %v", err.Error())
	}

	if resp.Data.ID != "123e4567-e89b-12d3-a456-426614174123" {
		t.Errorf("Wrong ID: got %v expected %v", resp.Data.ID, "123e4567-e89b-12d3-a456-426614174123")
	}

	AccountDelete("123e4567-e89b-12d3-a456-426614174123")
}

func TestAccountFetchNotFound(t *testing.T) {
	var request account.RequestFetch
	request.Host = URL + "v1/organisation/accounts"
	request.AccountID = "123e4567-e89b-12d3-a456-426614174000"

	resp, err := account.AccountFetch(&request)
	if resp != nil {
		t.Errorf("Unexpected response")
	}

	if err != nil && err.Error() != "404 Not Found" {
		t.Errorf("Unexpected error, got %v expected %v", err.Error(), "404 Not Found")
	}
}

func TestAccountDeleteSuccessful(t *testing.T) {
	account_data := `{
		"data": {
		  "id": "123e4567-e89b-12d3-a456-426614174123",
		  "organisation_id": "123e4567-e89b-12d3-a456-426614174123",
		  "type": "accounts",
		  "attributes": {
			 "country": "GB",
			  "base_currency": "GBP",
			  "bank_id": "400302",
			  "bank_id_code": "GBDSC",
			  "account_number": "10000004",
			  "iban": "GB28NWBK40030212764204",
			  "bic": "NWBKGB42",
			  "account_classification": "Personal",
			  "name": ["TestName"]
		  }
		}
	  }`

	_, e := AccountCreate(account_data)
	if e != nil {
		t.Errorf("Unexpected error")
	}

	var request account.RequestDelete
	request.Host = URL + "v1/organisation/accounts"
	request.AccountID = "123e4567-e89b-12d3-a456-426614174123"
	request.Version = "?version=0"

	err := account.AccountDelete(&request)

	if err != nil {
		t.Errorf("Unexpected error: %v", err.Error())
	}
}

func TestAccountDeleteNotFound(t *testing.T) {
	var request account.RequestDelete
	request.Host = URL + "v1/organisation/accounts"
	request.AccountID = "123e4567-e89b-12d3-a456-426614174000"
	request.Version = "?version=0"

	err := account.AccountDelete(&request)

	if err != nil && err.Error() != "404 Not Found" {
		t.Errorf("Unexpected error, got %v expected %v", err.Error(), "404 Not Found")
	}
}
