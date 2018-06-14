package client

import (
	"encoding/json"
	"net/url"
	"strconv"

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
