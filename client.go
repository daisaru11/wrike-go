package wrike

import (
	"net/http"
)

type ClientOptions struct {
	BaseURL     string
	OAuth2Token string
}

type Client struct {
	baseURL     string
	oauth2Token string
	httpClient  *http.Client
}

func NewClient(options *ClientOptions) *Client {
	baseURL := "https://www.wrike.com/api/v4/"
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	oauth2Token := options.OAuth2Token

	return &Client{
		baseURL:     baseURL,
		oauth2Token: oauth2Token,
		httpClient:  http.DefaultClient,
	}
}
