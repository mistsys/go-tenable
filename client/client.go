package tenable

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"

	"golang.org/x/net/context/ctxhttp"

	"github.com/pkg/errors"
)

// TODO this doesn't add any utility, and takes up more space
// just use http.Response and nopcloser to get raw body reuse
// or just don't even pass the response around; there's really no need
type Response struct {
	RawResponse *http.Response
	RawBody     []byte
}

// TODO error handling
func (r *Response) BodyJson() string {
	var buf bytes.Buffer
	if err := json.Indent(&buf, r.RawBody, "", "  "); err != nil {
		panic(errors.Wrapf(err, "Failed to format JSON body"))
	}
	return string(buf.Bytes())
}

type QueryOpts struct {
	// your normal key=value,key=value params
	Params string // TODO map
	// for Tenable's query param array format
	// not implemented. the scheme takes args like []string{"filter=foo,quality=bar",...}
	// to filter.0.name=foo,filter.0.quality=bar...
	// Filters []string
}

const tenableAPI = "https://cloud.tenable.com"

type Client struct {
	baseURL string
	client  *http.Client
	common  service
	// AccessKey for service
	accessKey string
	secretKey string
	// turn this on if you want to dump request/response
	Debug bool
	// username to impersonate as
	impersonate string

	// all the service objects defined in lowercaseservicename.go
	Editor      *EditorService
	Folders     *FoldersService
	Server      *ServerService
	Scans       *ScansService
	Scanners    *ScannersService
	Workbenches *WorkbenchesService

	// Query parameters struct
	QueryOpts *QueryOpts
}

type service struct {
	client *Client
}

func NewClient(accessKey string, secretKey string) *Client {
	c := &Client{
		baseURL:   tenableAPI,
		accessKey: accessKey,
		secretKey: secretKey,
		client:    http.DefaultClient,
	}
	c.common.client = c
	c.Editor = (*EditorService)(&c.common)
	c.Folders = (*FoldersService)(&c.common)
	c.Server = (*ServerService)(&c.common)
	c.Scans = (*ScansService)(&c.common)
	c.Scanners = (*ScannersService)(&c.common)
	c.Workbenches = (*WorkbenchesService)(&c.common)
	return c
}

func (t *Client) NewRequest(method string, relativeUrl string, body io.Reader) (*http.Request, error) {
	u, err := url.Parse(t.baseURL)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, relativeUrl)
	rawQuery := kvToQuery(t.QueryOpts.Params)
	u.RawQuery = rawQuery

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create request")

	}

	if body != nil {
		// take note!
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("X-ApiKeys", fmt.Sprintf("accessKey=%s; secretKey=%s", t.accessKey, t.secretKey))
	if t.impersonate != "" {
		req.Header.Set("X-Impersonate", fmt.Sprintf("username=%s", t.impersonate))
	}
	if t.Debug {
		requestDump, err := httputil.DumpRequest(req, true)
		if err != nil {
			log.Println(err)
		}
		log.Println(string(requestDump))
	}
	return req, nil
}

func (t *Client) Do(ctx context.Context, req *http.Request, dest interface{}) (*Response, error) {
	res, err := ctxhttp.Do(ctx, t.client, req)
	response := &Response{RawResponse: res}
	if err != nil { // hm
		return response, errors.Wrapf(err, "Failed to do request")
	}
	if res.StatusCode >= 400 {
		return response, errors.New(fmt.Sprintf("Error response from server: %d", res.StatusCode))
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response, errors.Wrapf(err, "Failed to read response.")
	}
	response.RawBody = buf

	if t.Debug {
		log.Printf("DEBUG body: %q", buf)
	}

	defer res.Body.Close()

	err = json.Unmarshal(buf, dest)
	if err != nil {
		return response, errors.Wrapf(err, "Failed to unmarshal")
	}

	return response, err
}

func (t *Client) Get(ctx context.Context, u string, opts *QueryOpts, dest interface{}) (*Response, error) {
	// t.QueryOpts = opts // TODO
	// nil body because it's a GET request
	req, err := t.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := t.Do(ctx, req, dest)
	return resp, err
}

func (t *Client) Post(ctx context.Context, u string, opts *QueryOpts, body io.Reader, dest interface{}) (*Response, error) {
	// t.QueryOpts = opts // TODO
	req, err := t.NewRequest(http.MethodPost, u, body)
	if err != nil {
		return nil, err
	}
	resp, err := t.Do(ctx, req, dest)
	return resp, err
}

func (t *Client) SetHttpClient(client *http.Client) {
	t.client = client
}

func (t *Client) SetBaseUrl(baseUrl string) {
	t.baseURL = baseUrl
}

func (t *Client) ImpersonateAs(username string) {
	t.impersonate = username
}
