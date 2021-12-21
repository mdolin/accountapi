package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

// MockClient is the mock client
type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

var (
	// GetDoFunc fetches the mock client's `Do` func
	DoFunc func(req *http.Request) (*http.Response, error)
)

// Do is the mock client's `Do` func
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return DoFunc(req)
}

func init() {
	Client = &MockClient{}
}

func TestGetOK(t *testing.T) {
	url := "http://api/accounts"

	expectedResponse := []byte("body")
	r := ioutil.NopCloser(bytes.NewReader([]byte("body")))

	DoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Status:     "",
			Body:       r,
		}, nil
	}

	res, err := Get(url)

	if err != nil {
		t.Errorf("A none 200 response: %v", err)
	}

	comparingResult := bytes.Compare(res, expectedResponse)
	if comparingResult != 0 {
		t.Errorf("Request differ from response, got %v expected %v", res, expectedResponse)
	}
}

func TestGetError(t *testing.T) {
	url := "http://api/accounts"

	expectedResponse := "404 Not Found"

	DoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Body:       nil,
		}, nil
	}

	res, err := Get(url)

	if res != nil {
		t.Errorf("A none 404 response")
	}

	if err.Error() != expectedResponse {
		t.Errorf("Status messages differ, got %v expected %v", err.Error(), expectedResponse)
	}
}

func TestPostOK(t *testing.T) {
	url := "http://api/accounts"
	data := []byte("body")

	expectedResponse := []byte("body")
	r := ioutil.NopCloser(bytes.NewReader([]byte("body")))

	DoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Status:     "",
			Body:       r,
		}, nil
	}

	res, err := Post(url, data)

	if err != nil {
		t.Errorf("A none 200 response: %v", err)
	}

	comparingResult := bytes.Compare(res, expectedResponse)
	if comparingResult != 0 {
		t.Errorf("Request differ from response, got %v expected %v", res, expectedResponse)
	}
}

func TestPostError(t *testing.T) {
	url := "http://api/accounts"
	data := []byte("body")

	expectedResponse := "500 Internal Server Error"

	DoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusInternalServerError,
			Status:     "500 Internal Server Error",
			Body:       nil,
		}, nil
	}

	res, err := Post(url, data)

	if res != nil {
		t.Errorf("A none 500 response")
	}

	if err.Error() != expectedResponse {
		t.Errorf("Status messages differ, got %v expected %v", err.Error(), expectedResponse)
	}
}

func TestDeleteOK(t *testing.T) {
	url := "http://api/accounts/ID"

	DoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusNoContent,
			Status:     "",
			Body:       nil,
		}, nil
	}

	err := Delete(url)

	if err != nil {
		t.Errorf("A none 204 response: %v", err)
	}
}

func TestDeleteError(t *testing.T) {
	url := "http://api/accounts"

	expectedResponse := "500 Internal Server Error"

	DoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusInternalServerError,
			Status:     "500 Internal Server Error",
			Body:       nil,
		}, nil
	}

	err := Delete(url)

	if err.Error() != expectedResponse {
		t.Errorf("Status messages differ, got %v expected %v", err.Error(), expectedResponse)
	}
}
