package client

import (
	"context"
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
	list := &FoldersList{}
	response, err := s.client.Get(ctx, "folders", nil, list)
	return list, response, err
}
