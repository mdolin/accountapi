package client

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var Client HTTPClient

func init() {
	Client = &http.Client{Timeout: 10 * time.Second}
}

func Get(url string) ([]byte, error) {
	// HTTP method GET to make request
	req, err := http.NewRequest(http.MethodGet, url, nil)

	// Handle Error
	if err != nil {
		return nil, err
	}

	resp, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 299 {
		return nil, errors.New(resp.Status)
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}

func Post(url string, data []byte) ([]byte, error) {
	// HTTP method POST to make request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))

	// Handle Error
	if err != nil {
		return nil, err
	}

	resp, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 299 {
		return nil, errors.New(resp.Status)
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)

	return body, err
}

func Delete(url string) error {
	// HTTP method DELETE to make request
	req, err := http.NewRequest(http.MethodDelete, url, nil)

	// Handle Error
	if err != nil {
		return err
	}

	resp, err := Client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 204 {
		return errors.New(resp.Status)
	}

	defer resp.Body.Close()

	return nil
}
