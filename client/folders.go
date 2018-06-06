package client

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

type FoldersService service

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

func (f *FoldersService) List(ctx context.Context) (*FoldersList, error) {
	req, err := f.client.createRequest(http.MethodGet, "folders", nil)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create request")
	}
	status := &FoldersList{}
	err = f.client.doRequest(ctx, req, status)
	return status, err
}
