package tenable

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
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

// Take key1=val1,key2=val2,... to key1=val1&key2=val2&...
func kvToQuery(s string) string {
	opts := make(url.Values)
	split := strings.Split(s, ",")
	for i := 0; i < len(split); i++ {
		// every key=val gets split; in case of degenerate key=val=something, the extra = is ignored
		opt := strings.SplitN(split[i], "=", 2)
		if len(opt)%2 == 0 {
			opts.Add(opt[0], opt[1])
		}
	}
	return opts.Encode()
}

// GROSS should not be here
func ReadScanConfig(configFile string) (*ScansCreateConfig, error) {
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to read %s", configFile)
	}
	config := &ScansCreateConfig{}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to parse %s", configFile)
	}
	return config, nil
}
