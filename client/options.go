package tenable

import (
	"fmt"
	"net/url"
)

// TODO later work to actually use this is array flags, ie, flags you can use multiple times to produce an array
// of values in the binding, eg
// tenable wb assets ls vulnerable --filter filter=host.hostname,value=something --filter filter=somethingelse,value=something

// For use in query strings
type FilterOpts struct {
	// the filter name; get from the filters endpoint or the Tenable API docs
	Filter  string `json:"filter"`
	Quality string `json:"quality"`
	Value   string `json:"value"`
}

type Filters struct {
	Opts []FilterOpts
}

// TODO test
func (f *Filters) ToQueryString() string {
	opts := new(url.Values)
	for i, filterOpts := range f.Opts {
		opts.Add(fmt.Sprintf("filter.%d.filter", i), filterOpts.Filter)
		opts.Add(fmt.Sprintf("filter.%d.quality", i), filterOpts.Quality)
		opts.Add(fmt.Sprintf("filter.%d.value", i), filterOpts.Value)
	}
	return opts.Encode()
}
