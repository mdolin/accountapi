package client

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	HTTPClient HttpClient
}

func CreateClient() *Client {
	return &Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *Client) Get(url string) ([]byte, error) {
	// HTTP method GET to make request
	req, err := http.NewRequest(http.MethodGet, url, nil)

	// Handle Error
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body, err
}

func (c *Client) Post(url string, data []byte) ([]byte, error) {
	// HTTP method POST to make request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))

	// Handle Error
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body, err
}

func (c *Client) Delete(url string) (int, error) {
	// HTTP method DELETE to make request
	req, err := http.NewRequest(http.MethodDelete, url, nil)

	// Handle Error
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	return resp.StatusCode, err
}
