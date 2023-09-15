package lastfm

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
)

type UserPeriod string

const (
	UserPeriodOverall = "overall"
	UserPeriod7Day    = "7day"
	UserPeriod1Month  = "1month"
	UserPeriod3Month  = "3month"
	UserPeriod6Month  = "6month"
	UserPeriod12onth  = "12month"
)

type getUserInfo struct {
	User *UserInfo `json:"user"`
}

// UserInfo is a last.fm user's info.
type UserInfo struct {
	Name        string   `json:"name"`
	Age         int      `json:"age,string"`
	Subscriber  int      `json:"subscriber,string"`
	RealName    string   `json:"realname"`
	Bootstrap   string   `json:"bootstrap"`
	Playcount   int      `json:"playcount,string"`
	ArtistCount int      `json:"artist_count,string"`
	Playlists   int      `json:"playlists,string"`
	TrackCount  int      `json:"track_count,string"`
	AlbumCount  int      `json:"album_count,string"`
	Image       []*Image `json:"image"`
	Registered  struct {
		UnixTime Time   `json:"unixtime"`
		Text     string `json:"text"`
	} `json:"registered"`
	Country string `json:"country"`
	Gender  string `json:"gender"`
	URL     string `json:"url"`
	Type    string `json:"type"`
}

// UserInfo retrieves a user's info.
func (c *Client) UserInfo(ctx context.Context, user string) (*UserInfo, error) {
	var uf *getUserInfo

	q := url.Values{}
	q.Set("method", "user.getinfo")
	q.Set("user", user)

	err := c.do(ctx, &uf, http.MethodGet, withQuery(q))
	if err != nil {
		return nil, err
	}

	return uf.User, nil
}

type topUserTracks struct {
	TopTracks *TopUserTracks `json:"toptracks"`
}

// TopUserTracks is a user's top tracks for a period.
type TopUserTracks struct {
	Tracks []*UserTrack       `json:"track"`
	Meta   *TopUserTracksMeta `json:"@attr"`
}

// TopUserTracksMeta is the meta information returned with a user.getTopTracks query.
type TopUserTracksMeta struct {
	User       string `json:"user"`
	TotalPages int    `json:"totalPages,string"`
	Page       int    `json:"page,string"`
	PerPage    int    `json:"perPage,string"`
	Total      int    `json:"total,string"`
}

// UserTrack represents a single track a user has listened to.
type UserTrack struct {
	Streamable struct {
		Fulltrack Bool   `json:"fulltrack"`
		Text      string `json:"text"`
	} `json:"streamable"`
	MBID   string   `json:"mbid"`
	Name   string   `json:"name"`
	Image  []*Image `json:"image"`
	Artist struct {
		URL  string `json:"url"`
		Name string `json:"name"`
		MBID string `json:"mbid"`
	} `json:"artist"`
	URL       string `json:"url"`
	Duration  int    `json:"duration,string"`
	Playcount int    `json:"playcount,string"`
	Meta      struct {
		Rank int `json:"rank,string"`
	} `json:"@attr"`
}

// TopUserTracks retrieves a user's top tracks for a period.
func (c *Client) TopUserTracks(ctx context.Context, user string, period UserPeriod, page, limit int) (*TopUserTracks, error) {
	var ut *topUserTracks

	q := url.Values{}
	q.Set("method", "user.gettoptracks")
	q.Set("user", user)

	if page > 0 {
		pageStr := strconv.FormatInt(int64(page), 10)

		q.Set("page", pageStr)
	}

	if limit > 0 {
		limitStr := strconv.FormatInt(int64(limit), 10)

		q.Set("limit", limitStr)
	}

	if period != "" {
		q.Set("period", string(period))
	}

	err := c.do(ctx, &ut, http.MethodGet, withQuery(q))
	if err != nil {
		return nil, err
	}

	return ut.TopTracks, nil
}
