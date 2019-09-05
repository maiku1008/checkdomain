package requests

import (
    "bytes"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

// MakeAuthString creates authentication string
func MakeAuthString(user, apikey string) string {
    credentials := fmt.Sprintf("%s:%s", user, apikey)
    encoded := base64.StdEncoding.EncodeToString([]byte(credentials))
    return fmt.Sprintf("Basic %s", encoded)
}

// Response defines the structure of the response
type Response struct {
    Expire  int    `json:"expire"`
    Token   string `json:"token"`
    Success bool   `json:"success"`
    Credit  string `json:"credit"`
    Message string `json:"message"`
    Error   int    `json:"error"`
}

// Request makes a request with the necessary params
func Request(method, endpoint, authstring, jsonBody string) Response {

    var jsonStr = []byte(jsonBody)

    req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(jsonStr))
    req.Header.Set("Authorization", authstring)
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("cache-control", "no-cache")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)

    var response Response
    err = json.Unmarshal(body, &response)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Message:", response.Message)

    return response
}
