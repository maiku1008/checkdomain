package requests

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestMakeAuthString(t *testing.T) {
    assert := assert.New(t)

    user := "foo"
    apikey := "bar"

    // echo -n foo:bar | base64
    // Zm9vOmJhcg==
    assert.Equal(MakeAuthString(user, apikey), "Basic Zm9vOmJhcg==")
}

// Fix this with Mocks!
func TestRequestOauthTokens(t *testing.T) {
    assert := assert.New(t)

    assert.True(false)
}
