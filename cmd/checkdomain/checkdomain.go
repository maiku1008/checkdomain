package main

import (
    "flag"
    "fmt"
    "github.com/joho/godotenv"
    "github.com/sk1u/checkdomain/internal"
    "os"
    "strings"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
    }

    domain := flag.String("d", "example.it", "The domain you wish to check")
    flag.Parse()

    user := os.Getenv("USERNAME")
    apikey := os.Getenv("APIKEY")

    authString := requests.MakeAuthString(user, apikey)

    fmt.Println("Requesting Auth Token...")
    response := requests.Request("POST", requests.TokenEndPoint, authString, requests.OauthTokenBody)

    bearerToken := fmt.Sprintf("Bearer %s", response.Token)

    checkEndpoint := fmt.Sprintf(requests.CheckEndpoint, *domain)
    fmt.Println(fmt.Sprintf("Checking if %s is available...", *domain))
    response = requests.Request("GET", checkEndpoint, bearerToken, "")

    if response.Success {
        fmt.Println("Creating Contact...")
        response = requests.Request("POST", requests.ContactEndpoint, bearerToken, requests.CreateContactBody)
        cID := strings.Split(response.Message, "'")[1]

        buyDomainBody := fmt.Sprintf(requests.BuyDomainBody, *domain, cID, cID, cID)
        fmt.Println("Acquiring domain...")
        response = requests.Request("POST", requests.DomainEndpoint, bearerToken, buyDomainBody)
        fmt.Println(response.Credit)
    }
    fmt.Println("Finishing...")
}
