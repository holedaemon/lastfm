package lastfm

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

// ErrorClientOption is returned when a Client is misconfigured.
var ErrorClientOption = errors.New("lastfm: client misconfigured")

// Client is an HTTP client for the Last.fm JSON API.
type Client struct {
	cli       http.Client
	apiKey    string
	userAgent string
}

// New creates a new Client.
func New(apiKey string, opts ...Option) (*Client, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("%w: missing api key", ErrorClientOption)
	}

	c := &Client{
		cli: http.Client{
			Timeout: time.Second * 10,
		},
		apiKey: apiKey,
	}

	for _, o := range opts {
		o(c)
	}

	if c.userAgent == "" {
		c.userAgent = libUserAgent
	}

	return c, nil
}
