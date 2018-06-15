package client

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// NumericBool type because Tenable sometimes returns
// 1 for what should be boolean
type NumericBool bool

func (n NumericBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(n)
}

func (n NumericBool) UnmarshalJSON(b []byte) error {
	strB := string(b)
	switch strB {
	case "true":
		n = true
	case "false":
		n = false
	default:
		i, err := strconv.Atoi(string(b))
		if err != nil {
			return errors.Wrapf(err, "Failed to parse as int")
		}
		if i > 0 {
			n = true
		} else {
			n = false
		}
	}
	return nil
}

func makeUrl(base string, opts interface{}) (string, error) {
	u, err := url.Parse(base)
	if err != nil {
		return "", err
	}
	if opts != nil {
		optsVals, err := query.Values(opts)
		if err != nil {
			return "", errors.Wrapf(err, "Failed to create query string")
		}
		u.RawQuery = optsVals.Encode()
	}
	return u.String(), nil
}

// tiny thing to key1=val1,key2=val2,... and returns key1=val1&key2=val2&...
func toQueryString(s string) string {
	opts := new(url.Values)
	split := strings.Split(s, ",")
	for i := 0; i < len(split); i++ {
		// every key=val gets split; if you input a degenerate key=val=something, it just silently ignores the extra =
		opt := strings.SplitN(split[i], "=", 2)
		opts.Add(opt[0], opt[1])
	}
	return opts.Encode()
}

// maybe it'll be more someday but this is just a json unmarshal. imo, there's not much reason to verify the parameters clientside
func checkPayload(payload []byte) error {
	return json.Unmarshal(payload, nil)
}
