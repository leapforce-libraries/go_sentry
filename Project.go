package sentry

import (
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	s_types "github.com/leapforce-libraries/go_sentry/types"
	go_types "github.com/leapforce-libraries/go_types"
)

type Project struct {
	ID          go_types.Int64String    `json:"id"`
	Slug        string                  `json:"slug"`
	Name        string                  `json:"name"`
	DateCreated *s_types.DateTimeString `json:"dateCreated"`
	Status      string                  `json:"status"`
	Platform    string                  `json:"platform"`
}

// GetProjects returns all projects
//
func (service *Service) GetProjects() (*[]Project, *errortools.Error) {
	endpoint := "projects"
	projects := []Project{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		URL:           service.url(endpoint),
		ResponseModel: &projects,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &projects, nil
}
