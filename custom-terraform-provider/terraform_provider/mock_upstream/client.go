package mock_upstream

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type (
	MockUpstreamClient struct {
		bseUrl     string
		httpClient *http.Client
	}

	CustomError struct {
		Code    int
		Message string
	}
)

func (customError *CustomError) Error() string { return customError.Message }

func (mockUpstreamClient *MockUpstreamClient) setHttpClient() {
	client := &http.Client{}

	mockUpstreamClient.httpClient = client
}

func (mockUpstreamClient *MockUpstreamClient) put(ctx context.Context, path string, requestBody interface{}) (*http.Response, error) {
	resourceUrl := mockUpstreamClient.bseUrl + path

	request, err := http.NewRequestWithContext(ctx, http.MethodPut, resourceUrl, nil)
	if err != nil {
		return nil, err
	}

	response, err := mockUpstreamClient.doRequest(ctx, request, requestBody)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (mockUpstreamClient *MockUpstreamClient) post(ctx context.Context, path string, requestBody interface{}) (*http.Response, error) {
	resourceUrl := mockUpstreamClient.bseUrl + path

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, resourceUrl, nil)
	if err != nil {
		return nil, err
	}

	response, err := mockUpstreamClient.doRequest(ctx, request, requestBody)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (mockUpstreamClient *MockUpstreamClient) get(ctx context.Context, path string) (*http.Response, error) {
	resourceUrl := mockUpstreamClient.bseUrl + path

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, resourceUrl, nil)
	if err != nil {
		return nil, err
	}

	response, err := mockUpstreamClient.doRequest(ctx, request, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (mockUpstreamClient *MockUpstreamClient) delete(ctx context.Context, path string) (*http.Response, error) {
	resourceUrl := mockUpstreamClient.bseUrl + path

	request, err := http.NewRequestWithContext(ctx, http.MethodDelete, resourceUrl, nil)
	if err != nil {
		return nil, err
	}

	response, err := mockUpstreamClient.doRequest(ctx, request, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (mockUpstreamClient *MockUpstreamClient) doRequest(ctx context.Context, request *http.Request, requestBody interface{}) (*http.Response, error) {
	payload, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	if payload != nil {
		request.Body = io.NopCloser(bytes.NewReader(payload))
	}

	response, err := mockUpstreamClient.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode >= 400 {
		return nil, &CustomError{
			Code:    response.StatusCode,
			Message: fmt.Sprintf("Error sending %s to %s. Request failed with %d:%s", request.Method, request.URL.Path, response.StatusCode, response.Status),
		}
	}

	return response, nil
}

func BuildClient(baseUrl string) *MockUpstreamClient {
	client := &MockUpstreamClient{bseUrl: baseUrl}
	client.setHttpClient()

	return client
}
