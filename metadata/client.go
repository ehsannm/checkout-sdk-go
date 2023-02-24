package metadata

import (
	"encoding/json"

	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/httpclient"
)

const path = "metadata/card"

// Client ...
type Client struct {
	API checkout.HTTPClient
}

// NewClient ...
func NewClient(config checkout.Config) *Client {
	config.BearerAuthentication = true
	return &Client{
		API: httpclient.NewClient(config),
	}
}

// Request ...
func (c *Client) Request(request *Request) (*Response, error) {
	response, err := c.API.Post("/"+path, request, nil)
	if err != nil {
		return nil, err
	}

	resp := &Response{}
	err = json.Unmarshal(response.ResponseBody, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
