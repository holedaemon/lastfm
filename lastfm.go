package lastfm

const (
	root    = "https://ws.audioscrobbler.com"
	version = "2.0"

	baseURL = root + "/" + version + "/"
)

type Image struct {
	Size string `json:"size"`
	Text string `json:"#text"`
}
