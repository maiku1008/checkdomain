package requests

import (
    "github.com/stretchr/testify/mock"
    "net/http"
)

// ClientInterface ...
type ClientInterface interface {
    Do(req *http.Request) (*http.Response, error)
}

// Validate satisfaction of interface
var _ ClientInterface = (*MockClient)(nil)
var _ ClientInterface = (*http.Client)(nil)

// MockClient ...
type MockClient struct {
    mock.Mock
}

// Do is a mock of client's method Do
func (mockClient *MockClient) Do(req *http.Request) (*http.Response, error) {
    args := mockClient.Called(req)
    return args.Get(0).(*http.Response), args.Error(1)
}
