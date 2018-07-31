// TODO consider renaming editor to template
package tenable

import (
	"context"
	"fmt"
)

type EditorService service

type Template struct {
	Unsupported      bool        `json:"unsupported"`
	CloudOnly        bool        `json:"cloud_only"`
	Desc             string      `json:"desc"`
	Order            interface{} `json:"order"`
	SubscriptionOnly bool        `json:"subscription_only"`
	IsWas            interface{} `json:"is_was"`
	Title            string      `json:"title"`
	IsAgent          interface{} `json:"is_agent"`
	UUID             string      `json:"uuid"`
	ManagerOnly      bool        `json:"manager_only"`
	Name             string      `json:"name"`
}

type Templates struct {
	Templates []Template `json:"templates"`
}

// List templates (API supports "scan" and "policy" template types)
func (s *EditorService) List(ctx context.Context, templateType string) (*Templates, *Response, error) {
	u := fmt.Sprintf("editor/%s/templates", templateType)
	templates := &Templates{}
	response, err := s.client.Get(ctx, u, nil, templates)
	return templates, response, err
}
