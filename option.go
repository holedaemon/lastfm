package lastfm

// Option configures a client.
type Option func(*Client)

// UserAgent sets a client's User-Agent header.
func UserAgent(ua string) Option {
	return func(c *Client) {
		c.userAgent = ua
	}
}
