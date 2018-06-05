package client

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestTenableClient_createRequest(t *testing.T) {
	type fields struct {
		baseURL     string
		client      *http.Client
		accessKey   string
		secretKey   string
		Debug       bool
		impersonate string
	}
	type args struct {
		method      string
		relativeUrl string
		data        url.Values
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Request
		wantErr bool
	}{
		{ // X-ApiKeys: accessKey={accessKey}; secretKey={secretKey};
			name: "Requests include the X-ApiKeys headers in required format",
			fields: fields{
				baseURL:   "https://cloud.tenable.com",
				client:    http.DefaultClient,
				accessKey: "fakeKey",
				secretKey: "fakeSecret",
			},
			args: args{
				method:      "GET",
				relativeUrl: "/foo",
				data:        nil,
			},
			want: &http.Request{
				Header: map[string][]string{
					// apparently Go downcases the keys
					"X-Apikeys": []string{"accessKey=fakeKey; secretKey=fakeSecret"},
				},
			},
		},
		{ // X-ApiKeys: accessKey={accessKey}; secretKey={secretKey};
			name: "Requests include the X-Impersonate headers in required if Impersonate is set",
			fields: fields{
				baseURL:     "https://cloud.tenable.com",
				client:      http.DefaultClient,
				accessKey:   "fakeKey",
				secretKey:   "fakeSecret",
				impersonate: "fakeUser",
			},
			args: args{
				method:      "GET",
				relativeUrl: "/foo",
				data:        nil,
			},
			want: &http.Request{
				Header: map[string][]string{
					"X-Apikeys":     []string{"accessKey=fakeKey; secretKey=fakeSecret"},
					"X-Impersonate": []string{"username=fakeUser"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tc := &TenableClient{
				baseURL:     tt.fields.baseURL,
				client:      tt.fields.client,
				accessKey:   tt.fields.accessKey,
				secretKey:   tt.fields.secretKey,
				Debug:       tt.fields.Debug,
				impersonate: tt.fields.impersonate,
			}
			got, err := tc.createRequest(tt.args.method, tt.args.relativeUrl, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("TenableClient.createRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Header, tt.want.Header) {
				t.Errorf("TenableClient.createRequest() = %v, want %v", got.Header, tt.want.Header)
			}
		})
	}
}
