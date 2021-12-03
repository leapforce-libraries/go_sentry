package sentry

import (
	"fmt"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	s_types "github.com/leapforce-libraries/go_sentry/types"
	go_types "github.com/leapforce-libraries/go_types"
)

type Event struct {
	ID          string                 `json:"id"`
	EventType   string                 `json:"event.type"`
	GroupID     go_types.Int64String   `json:"groupID"`
	EventID     string                 `json:"eventID"`
	ProjectID   go_types.Int64String   `json:"projectID"`
	Message     string                 `json:"message"`
	Title       string                 `json:"title"`
	Location    string                 `json:"location"`
	Culprit     string                 `json:"culprit"`
	User        *User                  `json:"user"`
	Tags        []Tag                  `json:"tags"`
	Platform    string                 `json:"platform"`
	DateCreated s_types.DateTimeString `json:"dateCreated"`
	Context     map[string]string      `json:"context"`
}

type GetEventsConfig struct {
	OrganizationSlug string
	ProjectSlug      string
	Full             bool
}

// GetEvents returns all events
//
func (service *Service) GetEvents(config *GetEventsConfig) (*[]Event, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetEventsConfig must not be a nil pointer")
	}

	values := url.Values{}
	values.Set("full", fmt.Sprintf("%v", config.Full))

	endpoint := fmt.Sprintf("projects/%s/%s/events/?%s", config.OrganizationSlug, config.ProjectSlug, values.Encode())
	events := []Event{}

	url := service.url(endpoint)

	for url != "" {
		_events := []Event{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			URL:           url,
			ResponseModel: &_events,
		}
		_, response, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		events = append(events, _events...)

		url = nextURL(response)
	}

	return &events, nil
}
