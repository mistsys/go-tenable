package tenable

import (
	"context"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestResponse_BodyJson(t *testing.T) {
	tests := []struct {
		name string
		r    *Response
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.BodyJson(); got != tt.want {
				t.Errorf("Response.BodyJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		accessKey string
		secretKey string
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.accessKey, tt.args.secretKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_NewRequest(t *testing.T) {
	type args struct {
		method      string
		relativeUrl string
		body        io.Reader
	}
	tests := []struct {
		name    string
		t       *Client
		args    args
		want    *http.Request
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.t.NewRequest(tt.args.method, tt.args.relativeUrl, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.NewRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.NewRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Do(t *testing.T) {
	type args struct {
		ctx  context.Context
		req  *http.Request
		dest interface{}
	}
	tests := []struct {
		name    string
		t       *Client
		args    args
		want    *Response
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.t.Do(tt.args.ctx, tt.args.req, tt.args.dest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Get(t *testing.T) {
	type args struct {
		ctx  context.Context
		u    string
		opts *QueryOpts
		dest interface{}
	}
	tests := []struct {
		name    string
		t       *Client
		args    args
		want    *Response
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.t.Get(tt.args.ctx, tt.args.u, tt.args.opts, tt.args.dest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Post(t *testing.T) {
	type args struct {
		ctx  context.Context
		u    string
		opts *QueryOpts
		body io.Reader
		dest interface{}
	}
	tests := []struct {
		name    string
		t       *Client
		args    args
		want    *Response
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.t.Post(tt.args.ctx, tt.args.u, tt.args.opts, tt.args.body, tt.args.dest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Post() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Post() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_SetHttpClient(t *testing.T) {
	type args struct {
		client *http.Client
	}
	tests := []struct {
		name string
		t    *Client
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.t.SetHttpClient(tt.args.client)
		})
	}
}

func TestClient_SetBaseUrl(t *testing.T) {
	type args struct {
		baseUrl string
	}
	tests := []struct {
		name string
		t    *Client
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.t.SetBaseUrl(tt.args.baseUrl)
		})
	}
}

func TestClient_ImpersonateAs(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		t    *Client
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.t.ImpersonateAs(tt.args.username)
		})
	}
}
