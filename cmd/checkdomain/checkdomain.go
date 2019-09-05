package main

import (
    "flag"
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

    cc := flag.Bool("cc", false, "Create Admin Contact")
    domain := flag.String("d", "example.it", "The domain you wish to check")
    _ = domain
    flag.Parse()

    user := os.Getenv("USERNAME")
    apikey := os.Getenv("APIKEY")

    authString := requests.MakeAuthString(user, apikey)

    fmt.Println("Requesting Auth Token...")
    response := requests.Request("POST", requests.TokenEndPoint, authString, requests.OauthTokenBody)

    bearerToken := fmt.Sprintf("Bearer %s", response.Token)

    // We might only need the following once
    if *cc {

        fmt.Println("Creating Contact...")
        response = requests.Request("POST", requests.ContactEndpoint, bearerToken, requests.CreateContactBody)

        fmt.Println(response.Message)
        return
    }

    // Find way to pass domain here
    fmt.Println("Checking if domain is available...")
    response = requests.Request("GET", requests.CheckEndpoint, bearerToken, "")

    if response.Success {
        fmt.Println("Acquiring domain...")
        response = requests.Request("POST", requests.DomainEndpoint, bearerToken, requests.BuyDomainBody)
    }
    fmt.Println("Finishing...")
}
