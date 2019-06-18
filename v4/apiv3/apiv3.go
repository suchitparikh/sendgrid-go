package apiv3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// APIv3BaseURL is the base URL for SendGrid v3 APIs
	APIv3BaseURL = "https://api.sendgrid.com/v3"
)

// ErrorDetails corresponds to individual error item in API v3 error response
type ErrorDetails struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Response represents a client response for SendGrid v3 API client
type Response struct {
	RawBody    []byte
	StatusCode int
	Header     http.Header

	Errors []ErrorDetails `json:"errors"`
}

// Client is a client for SendGrid v3 APIs
type Client struct {
	HTTPClient *http.Client
	Header     http.Header
}

// Get provides a generic method to make a API v3 GET call for any given endpoint and returns
// a comprehensive response that includes the original response body, status code
// and a slice of error messages returned by the API, if any
//
// If response.Errors is an empty slice, the body should be unmarshaled to respective
// response struct
//
// This method can be used to make an API call for which the client method is not yet
// available
func (c *Client) Get(endpoint string, params url.Values) (resp *Response, err error) {
	url := fmt.Sprintf("%s/%s?%s", APIv3BaseURL, endpoint, params.Encode())
	resp = &Response{}

	// create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating new API v3 GET request: %s", err)
	}

	// add request headers
	req.Header = c.Header

	// make HTTP GET request
	r, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making API v3 GET request: %s", err)
	}
	defer r.Body.Close()
	resp.StatusCode = r.StatusCode
	resp.Header = r.Header

	// read HTTP response body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading API v3 response body: %s", err)
	}
	resp.RawBody = body

	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling API v3 response: %s", err)
	}

	return resp, nil
}

// Post provides a generic method to make a API v3 POST call for any given endpoint and returns
// a comprehensive response that includes the original response body, status code
// and a slice of error messages returned by the API, if any
//
// If response.Errors is an empty slice, the body should be unmarshaled to respective
// response struct
//
// This method can be used to make an API call for which the client method is not yet
// available
func (c *Client) Post(endpoint string, body []byte) (resp *Response, err error) {
	url := fmt.Sprintf("%s/%s", APIv3BaseURL, endpoint)
	resp = &Response{}

	// create POST request
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("error creating new API v3 POST request: %s", err)
	}

	// add request headers
	req.Header = c.Header

	// make HTTP POST request
	r, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making API v3 POST request: %s", err)
	}
	resp.StatusCode = r.StatusCode
	resp.Header = r.Header

	// read HTTP response body
	resp.RawBody, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading API v3 response body: %s", err)
	}
	r.Body.Close()

	return resp, nil
}
