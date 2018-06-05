package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strings"

	"golang.org/x/net/context/ctxhttp"

	"github.com/pkg/errors"
)

type TenableClient struct {
	baseURL string
	client  *http.Client
	// AccessKey for service
	accessKey string
	secretKey string
	// turn this on if you want to dump request/response
	Debug bool
	//username to impersonate as
	impersonate string
}

func NewClient(accessKey string, secretKey string) *TenableClient {
	return &TenableClient{
		baseURL:   "https://cloud.tenable.com",
		client:    http.DefaultClient,
		accessKey: accessKey,
		secretKey: secretKey,
	}
}

func (t *TenableClient) createRequest(method string, relativeUrl string, data url.Values) (*http.Request, error) {
	u, _ := url.Parse(t.baseURL)
	u.Path = path.Join(u.Path, relativeUrl)
	req, err := http.NewRequest(method, u.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create request")

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

func (t *TenableClient) doRequest(ctx context.Context, req *http.Request, obj interface{}) error {
	res, err := ctxhttp.Do(ctx, t.client, req)
	if err != nil {
		return errors.Wrapf(err, "Failed to do request")
	}
	if t.Debug {
		buf, bodyErr := ioutil.ReadAll(res.Body)
		if bodyErr != nil {
			return errors.Wrapf(bodyErr, "Failed to read response.")
		}
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
		log.Printf("DEBUG body: %q", buf)
		res.Body = rdr2
	}
	defer res.Body.Close()
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(obj)
	if err != nil {
		return errors.Wrapf(err, "Failed to unmarshal")
	}
	return err
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

func (t *TenableClient) ScansList(ctx context.Context) (*ScansList, error) {
	req, err := t.createRequest(http.MethodGet, "scans", nil)
	if err != nil {
		log.Printf("Failed to create request %s", err)
		return nil, errors.Wrapf(err, "Failed to create request")
	}
	list := &ScansList{}
	err = t.doRequest(ctx, req, list)
	return list, err
}

func (t *TenableClient) ServerProperties(ctx context.Context) (*ServerProperties, error) {
	req, err := t.createRequest(http.MethodGet, "server/properties", nil)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create request")
	}
	props := &ServerProperties{}
	err = t.doRequest(ctx, req, props)
	return props, err
}

func (t *TenableClient) ServerStatus(ctx context.Context) (*ServerStatus, error) {
	req, err := t.createRequest(http.MethodGet, "server/status", nil)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create request")
	}
	status := &ServerStatus{}
	err = t.doRequest(ctx, req, status)
	return status, err
}

func (t *TenableClient) ScanDetail(ctx context.Context, id string) (*ScanDetail, error) {
	req, err := t.createRequest(http.MethodGet, fmt.Sprintf("scans/%s", id), nil)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create request")
	}
	status := &ScanDetail{}
	err = t.doRequest(ctx, req, status)
	return status, err
}

func (t *TenableClient) FoldersList(ctx context.Context) (*FoldersList, error) {
	req, err := t.createRequest(http.MethodGet, "folders", nil)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create request")
	}
	status := &FoldersList{}
	err = t.doRequest(ctx, req, status)
	return status, err
}
