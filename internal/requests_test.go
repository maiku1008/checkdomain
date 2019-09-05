package requests

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestMakeAuthString(t *testing.T) {
    assert := assert.New(t)

    user := "foo"
    apikey := "bar"

    assert.Equal(MakeAuthString(user, apikey), "Basic Zm9vOmJhcg==")
}

// Fix this with Mocks!
func TestRequest(t *testing.T) {
    assert := assert.New(t)

    //     {
    //     "scopes": [
    //         "POST:test.domains.altravia.com\/domain",
    //         "POST:test.domains.altravia.com\/contact",
    //         "GET:test.domains.altravia.com\/check"
    //     ],
    //     "expire": 1599073035,
    //     "token": "5d6eb78bed92ef09cb60699a",
    //     "success": true,
    //     "message": "",
    //     "error": null
    // }
    assert.True(false)
}
