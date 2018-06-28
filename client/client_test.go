package client

import (
	"context"
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
		want *TenableClient
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

func TestTenableClient_NewRequest(t *testing.T) {
	type args struct {
		method      string
		relativeUrl string
		body        interface{}
	}
	tests := []struct {
		name    string
		t       *TenableClient
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
				t.Errorf("TenableClient.NewRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TenableClient.NewRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTenableClient_Do(t *testing.T) {
	type args struct {
		ctx  context.Context
		req  *http.Request
		dest interface{}
	}
	tests := []struct {
		name    string
		t       *TenableClient
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
				t.Errorf("TenableClient.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TenableClient.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTenableClient_Get(t *testing.T) {
	type args struct {
		ctx  context.Context
		u    string
		opts *TenableQueryOpts
		dest interface{}
	}
	tests := []struct {
		name    string
		t       *TenableClient
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
				t.Errorf("TenableClient.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TenableClient.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTenableClient_Post(t *testing.T) {
	type args struct {
		ctx  context.Context
		u    string
		opts *TenableQueryOpts
		body interface{}
		dest interface{}
	}
	tests := []struct {
		name    string
		t       *TenableClient
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
				t.Errorf("TenableClient.Post() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TenableClient.Post() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTenableClient_SetHttpClient(t *testing.T) {
	type args struct {
		client *http.Client
	}
	tests := []struct {
		name string
		t    *TenableClient
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

func TestTenableClient_SetBaseUrl(t *testing.T) {
	type args struct {
		baseUrl string
	}
	tests := []struct {
		name string
		t    *TenableClient
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

func TestTenableClient_ImpersonateAs(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		t    *TenableClient
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
