package lastfm

import (
	"net/url"
	"time"

	querystring "github.com/google/go-querystring/query"
)

type requestQuery interface {
	Values() (url.Values, error)
}

// UserQuery is used to configure a request to the users endpoint.
type UserQuery struct {
	// Every user request
	User string `url:"user"`

	// getRecentTracks / getTopAlbums
	Limit  int    `url:"limit,omitempty"`
	Page   int    `url:"page,omitempty"`
	Period string `url:"period,omitempty"`

	// getRecentTracks
	From     time.Time `url:"from,unix,omitempty"`
	To       time.Time `url:"to,unix,omitempty"`
	Extended bool      `url:"extended,int,omitempty"`
}

// Values returns a query's values as a querystring.
func (q *UserQuery) Values() (url.Values, error) {
	return querystring.Values(q)
}
