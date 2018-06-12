package client

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

type FoldersService service

type Folder struct {
	Custom      int    `json:"custom"`
	DefaultTag  int    `json:"default_tag"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	UnreadCount int    `json:"unread_count"`
}

type FoldersList struct {
	Folders []Folder `json:"folders"`
}

func (s *FoldersService) List(ctx context.Context) (*FoldersList, *Response, error) {
	req, err := s.client.createRequest(http.MethodGet, "folders", nil)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "Failed to create request")
	}
	status := &FoldersList{}
	response, err := s.client.doRequest(ctx, req, status)
	return status, response, err
}
