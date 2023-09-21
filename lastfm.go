package lastfm

const (
	root    = "https://ws.audioscrobbler.com"
	version = "2.0"

	baseURL = root + "/" + version + "/"

	libVersion   = "0"
	libUserAgent = "lastfm/v" + libVersion + " (https://github.com/holedaemon/lastfm)"
)

// Image represents an image sent by the API.
type Image struct {
	Size string `json:"size"`
	Text string `json:"#text"`
}
