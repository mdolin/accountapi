package account

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAccountDeleteNoAccountID(t *testing.T) {
	var request RequestDelete
	request.Host = "http://api/accounts"
	request.AccountID = ""
	request.Version = "?version=0"

	err := AccountDelete(&request)
	if err == nil {
		t.Errorf("Unexpected error: %v", err.Error())
	}
}

func TestAccountDeleteNotFound(t *testing.T) {
	testServer := httptest.NewServer(
		http.HandlerFunc(
			func(
				res http.ResponseWriter, req *http.Request) {
				res.WriteHeader(404)
			},
		),
	)
	defer func() { testServer.Close() }()

	var request RequestDelete
	request.Host = testServer.URL
	request.AccountID = "123e4567-e89b-12d3-a456-426614174123"
	request.Version = "?version=0"

	err := AccountDelete(&request)
	if err.Error() != "404 Not Found" {
		t.Errorf("Unexpected error: %v", err.Error())
	}
}

func TestAccountDelete(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if len(req.URL.Query()) == 1 {
			res.WriteHeader(204)
		}
	}))
	defer func() { testServer.Close() }()

	var request RequestDelete
	request.Host = testServer.URL
	request.AccountID = "123e4567-e89b-12d3-a456-426614174123"
	request.Version = "?version=0"

	err := AccountDelete(&request)
	if err != nil {
		t.Errorf("Unexpected error: %v", err.Error())
	}
}
