package main

import (
    "fmt"
    "github.com/joho/godotenv"
    "github.com/sk1u/checkdomain/internal"
    "os"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
    }
    // Find way not to hardcode endpoints
    var tokenEndPoint = "https://test.oauth.altravia.com/token/"

    user := os.Getenv("USERNAME")
    apikey := os.Getenv("APIKEY")

    authString := requests.MakeAuthString(user, apikey)

    requests.RequestOauthTokens(tokenEndPoint, authString)
}
