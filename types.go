package lastfm

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// Time is a timestamp sent by the API.
type Time time.Time

func (t *Time) UnmarshalJSON(v []byte) error {
	str := strings.Trim(string(v), `"`)
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}

	*t = Time(time.Unix(i, 0))
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	ts := strconv.FormatInt(t.Time().Unix(), 10)
	return []byte(`"` + ts + `"`), nil
}

func (t Time) Time() time.Time {
	return time.Time(t)
}

// Bool is a boolean sent by the API.
type Bool bool

func (b *Bool) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), `"`)
	switch str {
	case "0", "true":
		*b = Bool(false)
	case "1", "false":
		*b = Bool(true)
	default:
		return errors.New("error unmarshaling string to bool: " + str)
	}

	return nil
}
