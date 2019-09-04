package requests

import (
    "bytes"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

// MakeAuthString encodes user and apikey into base64
func MakeAuthString(user, apikey string) string {
    credentials := fmt.Sprintf("%s:%s", user, apikey)
    encoded := base64.StdEncoding.EncodeToString([]byte(credentials))
    return fmt.Sprintf("Basic %s", encoded)
}

// RequestOauthTokens requests oauth tokens for scopes
func RequestOauthTokens(tokenEndPoint, authString string) string {

    // Let's not hardcode these
    var jsonStr = []byte(`{"scopes":[
        "POST:test.domains.altravia.com/domain",
        "GET:test.domains.altravia.com/check",
        "POST:test.domains.altravia.com/contact"
        ]}`)

    req, err := http.NewRequest("POST", tokenEndPoint, bytes.NewBuffer(jsonStr))
    req.Header.Set("Authorization", authString)
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("cache-control", "no-cache")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)

    var response RequestOauthTokenResponse
    err = json.Unmarshal(body, &response)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("response Status:", resp.Status)
    // fmt.Println("response Headers:", resp.Header)
    // fmt.Println("response Body:", string(body))
    fmt.Println("Token:", response.Token)

    return response.Token
}

// RequestOauthTokenResponse defines the structure of the response
type RequestOauthTokenResponse struct {
    Expire  int    `json:"expire"`
    Token   string `json:"token"`
    Success bool   `json:"success"`
}
