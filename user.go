package lastfm

import (
	"context"
	"net/http"
)

type getUserInfo struct {
	User *UserInfo `json:"user"`
}

// UserInfo is a user's info sent by the API.
type UserInfo struct {
	Name        string   `json:"name"`
	Age         int      `json:"age,string"`
	Subscriber  int      `json:"subscriber,string"`
	RealName    string   `json:"realname"`
	Bootstrap   string   `json:"bootstrap"`
	PlayCount   int      `json:"playcount,string"`
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

// UserInfo fetches a user's info.
func (c *Client) UserInfo(ctx context.Context, query *UserQuery) (*UserInfo, error) {
	var uf *getUserInfo

	err := c.do(ctx, &uf, http.MethodGet, "user.getinfo", withQuery(query))
	if err != nil {
		return nil, err
	}

	return uf.User, nil
}

// UserMeta is a meta information about a request sent with some user endpoints.
type UserMeta struct {
	User       string `json:"user"`
	TotalPages int    `json:"totalPages,string"`
	Page       int    `json:"page,string"`
	PerPage    int    `json:"perPage,string"`
	Total      int    `json:"total,string"`
}

type getUserRecentTracks struct {
	RecentTracks *UserRecentTracks `json:"recenttracks"`
}

type UserRecentTracks struct {
	Tracks []*UserRecentTrack `json:"track"`
	Meta   *UserMeta          `json:"@attr"`
}

// UserRecentTrack represents a track recently listened to by a user.
type UserRecentTrack struct {
	Artist struct {
		MBID string `json:"mbid"`
		Name string `json:"#text"`
	} `json:"artist"`
	Album struct {
		MBID string `json:"mbid"`
		Name string `json:"#text"`
	} `json:"album"`
	Streamable Bool     `json:"streamable"`
	Image      []*Image `json:"image"`
	MBID       string   `json:"mbid"`
	Name       string   `json:"name"`
	URL        string   `json:"url"`
	Date       struct {
		Time Time `json:"uts"`
	}
	Meta struct {
		NowPlaying Bool `json:"nowplaying"`
	} `json:"@attr,omitempty"`
}

// UserRecentTracks fetches the tracks a user has most recently listened to.
func (c *Client) UserRecentTracks(ctx context.Context, query *UserQuery) (*UserRecentTracks, error) {
	var ut *getUserRecentTracks

	err := c.do(ctx, &ut, http.MethodGet, "user.getrecenttracks", withQuery(query))
	if err != nil {
		return nil, err
	}

	return ut.RecentTracks, nil
}

type getUserTopAlbums struct {
	Albums *UserTopAlbums `json:"topalbums"`
}

// UserTopAlbums is a user's top albums.
type UserTopAlbums struct {
	Albums []*UserTopAlbum `json:"album"`
	Meta   *UserMeta       `json:"@attr"`
}

// UserTopAlbum represents a top album for a user.
type UserTopAlbum struct {
	Artist struct {
		URL  string `json:"url"`
		Name string `json:"name"`
		MBID string `json:"mbid"`
	} `json:"artist"`
	Image     []*Image `json:"image"`
	MBID      string   `json:"mbid"`
	URL       string   `json:"url"`
	PlayCount int      `json:"playcount,string"`
	Name      string   `json:"name"`
	Meta      struct {
		Rank int `json:"rank,string"`
	} `json:"@attr"`
}

// UserTopAlbums fetches a user's top albums for a period.
func (c *Client) UserTopAlbums(ctx context.Context, query *UserQuery) (*UserTopAlbums, error) {
	var ut *getUserTopAlbums

	err := c.do(ctx, &ut, http.MethodGet, "user.gettopalbums", withQuery(query))
	if err != nil {
		return nil, err
	}

	return ut.Albums, nil
}
