// XXX naming this package client (and calling the client TenableClient) is poor style, or at least idiosyncratic
// usually, the package would be named tenable, and the client would just be Client. TODO make that change...
package client

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

// TODO this doesn't add any utility; just use http.Response and nopcloser to get raw body reuse
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

const tenableAPI = "https://cloud.tenable.com"

type TenableClient struct {
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
	Scans       *ScansService
	Folders     *FoldersService
	Server      *ServerService
	Workbenches *WorkbenchesService

	// Query parameters struct
	QueryOpts *TenableQueryOpts
}

type service struct {
	client *TenableClient
}

func NewClient(accessKey string, secretKey string) *TenableClient {
	c := &TenableClient{
		baseURL:   tenableAPI,
		accessKey: accessKey,
		secretKey: secretKey,
		client:    http.DefaultClient,
	}
	c.common.client = c
	c.Scans = (*ScansService)(&c.common)
	c.Folders = (*FoldersService)(&c.common)
	c.Server = (*ServerService)(&c.common)
	c.Workbenches = (*WorkbenchesService)(&c.common)
	return c
}

func (t *TenableClient) NewRequest(method string, relativeUrl string, body interface{}) (*http.Request, error) {
	u, err := url.Parse(t.baseURL)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, relativeUrl)
	rawQuery := kvToQuery(t.QueryOpts.Params)
	u.RawQuery = rawQuery

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create request")

	}

	if body != nil {
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

func (t *TenableClient) Do(ctx context.Context, req *http.Request, dest interface{}) (*Response, error) {
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

// so I've been using these like opts is set *on the client struct* rather than getting passed
// in because it reduces repetition in cmd. That means the opts arg here and in Post is unused
// It's a pretty idiosyncratic interface, so TODO switch to using the arg that's passed...
func (t *TenableClient) Get(ctx context.Context, u string, opts *TenableQueryOpts, dest interface{}) (*Response, error) {
	// nil body because it's a GET request
	req, err := t.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := t.Do(ctx, req, dest)
	return resp, err
}

func (t *TenableClient) Post(ctx context.Context, u string, opts *TenableQueryOpts, body interface{}, dest interface{}) (*Response, error) {
	req, err := t.NewRequest(http.MethodPost, u, body)
	if err != nil {
		return nil, err
	}
	resp, err := t.Do(ctx, req, dest)
	return resp, err
}

func (t *TenableClient) SetHttpClient(client *http.Client) {
	t.client = client
}

func (t *TenableClient) SetBaseUrl(baseUrl string) {
	t.baseURL = baseUrl
}

func (t *TenableClient) ImpersonateAs(username string) {
	t.impersonate = username
}
