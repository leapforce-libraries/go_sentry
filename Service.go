package sentry

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"github.com/tomnomnom/linkheader"
)

const (
	apiName string = "Sentry"
	apiUrl  string = "https://sentry.io/api/0"
	/*dateTimeFormat            string = "2006-01-02T15:04:05Z"
	dateTimeFormatCustomField string = "2006-01-02 15:04:05"
	dateFormat                string = "2006-01-02"
	defaultMaxRowCount        uint64 = ^uint64(0)
	defaultTop                uint64 = 500 //max 500, see: https://api.insightly.com/v3.1/Help#!/Overview/Introduction*/
)

type Service struct {
	apiKey      string
	httpService *go_http.Service
}

type ServiceConfig struct {
	ApiKey string
}

func NewService(serviceConfig *ServiceConfig) (*Service, *errortools.Error) {
	if serviceConfig == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	if serviceConfig.ApiKey == "" {
		return nil, errortools.ErrorMessage("Service Api Key not provided")
	}

	httpService, e := go_http.NewService(&go_http.ServiceConfig{})
	if e != nil {
		return nil, e
	}

	return &Service{
		apiKey:      serviceConfig.ApiKey,
		httpService: httpService,
	}, nil
}

func (service *Service) httpRequest(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	// add authentication header
	header := http.Header{}
	header.Set("Authorization", fmt.Sprintf("Bearer %s", service.apiKey))
	(*requestConfig).NonDefaultHeaders = &header

	// add error model
	errorResponse := ErrorResponse{}
	(*requestConfig).ErrorModel = &errorResponse

	request, response, e := service.httpService.HttpRequest(requestConfig)
	if errorResponse.Message != "" {
		e.SetMessage(errorResponse.Message)
	}

	return request, response, e
}

func nextUrl(response *http.Response) string {
	linkHeader := response.Header.Get("link")
	if linkHeader == "" {
		return ""
	}
	links := linkheader.Parse(linkHeader)
	for _, link := range links {
		if link.Rel == "next" {
			if link.Params["results"] == "true" {
				return link.URL
			}
			return ""
		}
	}

	return ""
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", apiUrl, path)
}

func (service *Service) ApiName() string {
	return apiName
}

func (service *Service) ApiKey() string {
	return service.apiKey
}

func (service *Service) ApiCallCount() int64 {
	return service.httpService.RequestCount()
}

func (service *Service) ApiReset() {
	service.httpService.ResetRequestCount()
}
