package lastfm

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type requestOption func(*request)

type request struct {
	query   url.Values
	body    io.ReadCloser
	headers http.Header
}

func withQuery(q url.Values) requestOption {
	return func(r *request) {
		r.query = q
	}
}

func withBody(b io.ReadCloser) requestOption {
	return func(r *request) {
		r.body = b
	}
}

func withHeaders(h http.Header) requestOption {
	return func(r *request) {
		r.headers = h
	}
}

func (c *Client) do(ctx context.Context, obj any, method string, opts ...requestOption) error {
	r := &request{}

	for _, o := range opts {
		o(r)
	}

	req, err := http.NewRequestWithContext(ctx, method, baseURL, r.body)
	if err != nil {
		return err
	}

	if r.headers != nil {
		req.Header = r.headers.Clone()
	}

	if r.query != nil {
		r.query.Set("api_key", c.apiKey)
		r.query.Set("format", "json")

		req.URL.RawQuery = r.query.Encode()
	} else {
		req.URL.Query().Set("api_key", c.apiKey)
		req.URL.Query().Set("format", "json")
	}

	res, err := c.cli.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		var ae *APIError
		if err := json.NewDecoder(res.Body).Decode(&ae); err != nil {
			return err
		}

		ae.HTTPStatus = res.StatusCode
		return ae
	}

	if err := json.NewDecoder(res.Body).Decode(&obj); err != nil {
		return err
	}

	return nil
}
