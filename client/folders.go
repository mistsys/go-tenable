package client

import (
	"context"
	"net/http"
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
	req, err := s.client.NewRequest(http.MethodGet, "folders", nil)
	if err != nil {
		return nil, nil, err
	}
	status := &FoldersList{}
	response, err := s.client.Do(ctx, req, status)
	return status, response, err
}
