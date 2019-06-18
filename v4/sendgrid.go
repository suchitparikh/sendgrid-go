// Package sendgrid provides a simple interface to interact with the SendGrid API
package sendgrid

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sendgrid/sendgrid-go/v4/apiv3"
)

const (
	// Version of the API Library
	Version = "4.0.0"

	// SGAPIV3 Versions of SendGrid APIs supported
	//
	// https://sendgrid.com/docs/API_Reference/api_v3.html
	SGAPIV3 = 3
)

// NewV3Client returns a new Client for SendGrid V3 APIs
func NewV3Client(key string) *apiv3.Client {
	c := &apiv3.Client{
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	// add common headers
	c.Header.Add("Authorization", fmt.Sprintf("Bearer %s", key))
	c.Header.Add("User-Agent", fmt.Sprintf("sendgrid-go/v%s", Version))
	c.Header.Add("Accept", "application/json")

	return c
}
