package requests

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "io/ioutil"
    "net/http"
    "strings"
    "testing"
)

const (
    authString  = "Basic Zm9vOmJhcg=="
    bearerToken = "Bearer 5d6eb78bed92ef09cb60699a"
    domain      = "test-domain.it"
    cID         = "AV-126457471"
)

func TestMakeAuthString(t *testing.T) {
    assert := assert.New(t)

    user := "foo"
    apikey := "bar"

    assert.Equal(MakeAuthString(user, apikey), authString)
}

// Prepare check domain endpoint
var checkEndpoint = fmt.Sprintf(CheckEndpoint, domain)

// Prepare buy domain body
var buyDomainBody = fmt.Sprintf(BuyDomainBody, domain, cID, cID, cID)

// A struct for testing the different calls
var responseTests = []struct {
    method       string
    endpoint     string
    auth         string
    requestBody  string
    responseBody string
}{
    {"POST", TokenEndPoint, authString, OauthTokenBody, tokenEndpointResponse},
    {"POST", ContactEndpoint, bearerToken, CreateContactBody, contactEndpointResponse},
    {"GET", checkEndpoint, bearerToken, "", checkEndpointResponse},
    {"POST", DomainEndpoint, bearerToken, buyDomainBody, buyDomainEndpointResponse},
}

func TestGetResponse(t *testing.T) {
    assert := assert.New(t)

    // Initialize mock object
    mockClient := &MockClient{}

    for _, rt := range responseTests {
        // Setup request
        req := SetRequest(rt.method, rt.endpoint, rt.auth, rt.requestBody)

        // Setup expected response
        body := ioutil.NopCloser(strings.NewReader(rt.responseBody))
        returnValue := &http.Response{Status: "200 OK", Body: body}

        // Set expectations
        mockClient.On("Do", req).Once().Return(returnValue, nil)

        // Run the thing!
        response := GetResponse(req, mockClient)

        // Assert expectations
        mockClient.AssertExpectations(t)
        assert.Equal(response.Success, true)
    }
}
