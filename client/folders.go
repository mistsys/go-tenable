package client

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

type FoldersList struct {
	Folders []Folder `json:"folders"`
}

type Folder struct {
	Custom      int    `json:"custom"`
	DefaultTag  int    `json:"default_tag"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	UnreadCount int    `json:"unread_count"`
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
