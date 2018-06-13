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
	// "strings"

	"golang.org/x/net/context/ctxhttp"

	// "github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

type Response struct {
	RawResponse *http.Response
	RawBody     []byte
	// so eventually, this will have stuff for like, pagination, or something
	// some of the endpoints return a "see more"? though that might be per Resource, not per response
}

// TODO there can be errors (maybe?); they need to be handled (maybe?)
func (r *Response) BodyJson() string {
	var buf bytes.Buffer
	_ = json.Indent(&buf, r.RawBody, "", "  ")
	// if err != nil {
	// 	return "", errors.Wrap(err, "Failed to format response body JSON")
	// }
	return string(buf.Bytes()) // , err
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
	u, _ := url.Parse(t.baseURL)
	u.Path = path.Join(u.Path, relativeUrl)

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

func (t *TenableClient) Do(ctx context.Context, req *http.Request, obj interface{}) (*Response, error) {
	// TODO we'll need to handle actual http errors too, like 40x. maybe have some sort of CheckResponse for all the error checking
	res, err := ctxhttp.Do(ctx, t.client, req)
	response := &Response{RawResponse: res}
	if err != nil {
		return response, errors.Wrapf(err, "Failed to do request")
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

	err = json.Unmarshal(buf, obj)
	if err != nil {
		return response, errors.Wrapf(err, "Failed to unmarshal")
	}

	return response, err
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
