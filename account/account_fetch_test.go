package account

import (
	"accountapi/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchAccountNoAccountID(t *testing.T) {
	var request RequestFetch
	request.Host = "http://api/accounts"
	request.AccountID = ""

	_, err := AccountFetch(&request)
	if err == nil {
		t.Errorf("We go unexpected error: %v", err.Error())
	}
}

func TestFetchAccountNotFound(t *testing.T) {
	testServer := httptest.NewServer(
		http.HandlerFunc(
			func(
				res http.ResponseWriter, req *http.Request) {
				res.WriteHeader(404)
			},
		),
	)
	defer func() { testServer.Close() }()

	var request RequestFetch
	request.Host = testServer.URL
	request.AccountID = "123e4567-e89b-12d3-a456-426614174123"

	_, err := AccountFetch(&request)
	if err.Error() != "404 Not Found" {
		t.Errorf("Unexpected error: got %v", err.Error())
	}
}

func TestFetchAccount(t *testing.T) {
	body := `{
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
			  "name": ["Burek"]
		  }
		}
	  }`
	var expectedRes model.AccountData
	json.Unmarshal([]byte(body), &expectedRes)

	expectedBody := []byte(body)

	testServer := httptest.NewServer(
		http.HandlerFunc(
			func(
				res http.ResponseWriter, req *http.Request) {
				res.WriteHeader(200)
				res.Write(expectedBody)
			},
		),
	)
	defer func() { testServer.Close() }()

	var request RequestFetch
	request.Host = testServer.URL
	request.AccountID = "123e4567-e89b-12d3-a456-426614174123"

	resp, err := AccountFetch(&request)

	if err != nil {
		t.Errorf("Oh snap, we got error: got %v", err.Error())
	}

	// Test Key-Values
	if resp.Data.ID != expectedRes.Data.ID {
		t.Errorf("Wrong ID: got %v expected %v", resp.Data.ID, expectedRes.Data.ID)
	}
	if resp.Data.OrganisationID != expectedRes.Data.OrganisationID {
		t.Errorf("Wrong OrganisationID: got %v expected %v", resp.Data.OrganisationID, expectedRes.Data.OrganisationID)
	}
	if resp.Data.Type != expectedRes.Data.Type {
		t.Errorf("Wrong Type: got %v expected %v", resp.Data.Type, expectedRes.Data.Type)
	}
	if *resp.Data.Attributes.Country != *expectedRes.Data.Attributes.Country {
		t.Errorf("Wrong Country: got %v expected %v", *resp.Data.Attributes.Country, *expectedRes.Data.Attributes.Country)
	}
	if resp.Data.Attributes.BaseCurrency != expectedRes.Data.Attributes.BaseCurrency {
		t.Errorf("Wrong BaseCurrency: got %v expected %v", resp.Data.Attributes.BaseCurrency, expectedRes.Data.Attributes.BaseCurrency)
	}
	if resp.Data.Attributes.BankID != expectedRes.Data.Attributes.BankID {
		t.Errorf("Wrong BankID: got %v expected %v", resp.Data.Attributes.BankID, expectedRes.Data.Attributes.BankID)
	}
	if resp.Data.Attributes.BankIDCode != expectedRes.Data.Attributes.BankIDCode {
		t.Errorf("Wrong BankIDCode: got %v expected %v", resp.Data.Attributes.BankIDCode, expectedRes.Data.Attributes.BankIDCode)
	}
	if resp.Data.Attributes.AccountNumber != expectedRes.Data.Attributes.AccountNumber {
		t.Errorf("Wrong AccountNumber: got %v expected %v", resp.Data.Attributes.AccountNumber, expectedRes.Data.Attributes.AccountNumber)
	}
	if resp.Data.Attributes.Iban != expectedRes.Data.Attributes.Iban {
		t.Errorf("Wrong Iban: got %v expected %v", resp.Data.Attributes.Iban, expectedRes.Data.Attributes.Iban)
	}
	if resp.Data.Attributes.Bic != expectedRes.Data.Attributes.Bic {
		t.Errorf("Wrong Bic: got %v expected %v", resp.Data.Attributes.Bic, expectedRes.Data.Attributes.Bic)
	}
	if *resp.Data.Attributes.AccountClassification != *expectedRes.Data.Attributes.AccountClassification {
		t.Errorf("Wrong AccountClassification: got %v expected %v", *resp.Data.Attributes.AccountClassification, *expectedRes.Data.Attributes.AccountClassification)
	}
	if len(resp.Data.Attributes.Name) == len(expectedRes.Data.Attributes.Name) {
		for i := 0; i < len(resp.Data.Attributes.Name); i++ {
			if resp.Data.Attributes.Name[i] != expectedRes.Data.Attributes.Name[i] {
				t.Errorf("Wrong Name: got %v expected %v", resp.Data.Attributes.Name[i], expectedRes.Data.Attributes.Name[i])
			}
		}
	} else {
		t.Errorf("Name array is not the same")
	}
}
