package requests

import (
    "github.com/stretchr/testify/mock"
    "net/http"
)

// ClientInterface defines a commond interface to
// *http.Client and our *MockClient
type ClientInterface interface {
    Do(req *http.Request) (*http.Response, error)
}

// Validate satisfaction of interface
var _ ClientInterface = (*MockClient)(nil)
var _ ClientInterface = (*http.Client)(nil)

// MockClient embeds mock.Mock to reuse its methods
type MockClient struct {
    mock.Mock
}

// Do is a mock of http.Client's method Do
func (mockClient *MockClient) Do(req *http.Request) (*http.Response, error) {
    args := mockClient.Called(req)
    return args.Get(0).(*http.Response), args.Error(1)
}
