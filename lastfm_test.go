package lastfm

import (
	"context"
	"os"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestLastFM(t *testing.T) {
	apiKey := os.Getenv("LASTFM_API_KEY")
	assert.Assert(t, apiKey != "")

	c, err := New(apiKey)
	assert.NilError(t, err, "creating client")

	ctx := context.Background()

	t.Run("TestUser", func(t *testing.T) {
		user, err := c.UserInfo(ctx, "holedaemon")
		assert.NilError(t, err, "getting user info")

		assert.Assert(t, user.Name == "holedaemon", "username is wrong")

		t.Logf("user name is %s\n", user.Name)
		t.Logf("user was created at %s\n", user.Registered.UnixTime.Time().Format(time.RFC3339))

		topTracks, err := c.TopUserTracks(ctx, user.Name, UserPeriodOverall, 0, 50)
		assert.NilError(t, err, "getting top tracks for user")

		assert.Assert(t, len(topTracks.Tracks) > 0, "top user tracks empty")

		t.Logf("top track for %s: %s\n", user.Name, topTracks.Tracks[0].Name)
		t.Logf("meta for top track: user: %s page: %d perpage: %d total: %d totalpages: %d\n",
			topTracks.Meta.User,
			topTracks.Meta.Page,
			topTracks.Meta.PerPage,
			topTracks.Meta.Total,
			topTracks.Meta.TotalPages,
		)
		t.Logf("streamable?: %v\b", topTracks.Tracks[0].Streamable.Fulltrack)
	})
}
