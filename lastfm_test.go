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
		user, err := c.UserInfo(ctx, &UserQuery{
			User: "holedaemon",
		})
		assert.NilError(t, err, "getting user info")

		assert.Assert(t, user.Name == "holedaemon", "username is wrong")

		t.Logf("user name is %s\n", user.Name)
		t.Logf("user was created at %s\n", user.Registered.UnixTime.Time().Format(time.RFC3339))

		recentTracks, err := c.UserRecentTracks(ctx, &UserQuery{
			User: "holedaemon",
		})
		assert.NilError(t, err, "getting recent tracks")
		assert.Assert(t, len(recentTracks.Tracks) > 0, "recent tracks empty")

		t.Logf("last track: %s", recentTracks.Tracks[0].Name)

		topAlbums, err := c.UserTopAlbums(ctx, &UserQuery{
			User:   "holedaemon",
			Period: "7day",
		})
		assert.NilError(t, err, "getting top albums")

		assert.Assert(t, len(topAlbums.Albums) > 0, "top albums is empty")

		t.Logf("top album: %s", topAlbums.Albums[0].Name)
	})
}
