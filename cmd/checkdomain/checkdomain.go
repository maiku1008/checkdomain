package main

import (
    "flag"
    "fmt"
    "github.com/joho/godotenv"
    "github.com/sk1u/checkdomain/internal"
    "net/http"
    "os"
    "strings"
)

func main() {
    // Initialize
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
    }

    domain := flag.String("d", "example.it", "The domain you wish to check")
    flag.Parse()

    user := os.Getenv("USERNAME")
    apikey := os.Getenv("APIKEY")
    client := &http.Client{}
    authString := requests.MakeAuthString(user, apikey)

    // Request auth tokens with authString
    fmt.Println("Requesting Auth Token...")
    req := requests.SetRequest("POST", requests.TokenEndPoint, authString, requests.OauthTokenBody)
    response := requests.GetResponse(req, client)
    fmt.Println("response Message:", response.Message)

    // Create Bearer Token, used from now for authentication
    bearerToken := fmt.Sprintf("Bearer %s", response.Token)

    // Check if domain is available
    checkEndpoint := fmt.Sprintf(requests.CheckEndpoint, *domain)
    fmt.Println(fmt.Sprintf("Checking if %s is available...", *domain))
    req = requests.SetRequest("GET", checkEndpoint, bearerToken, "")
    response = requests.GetResponse(req, client)
    fmt.Println("response Message:", response.Message)

    // If the domain is available
    if response.Success {

        // Create a contact to be associated with the domain
        fmt.Println("Creating Contact...")
        req = requests.SetRequest("POST", requests.ContactEndpoint, bearerToken, requests.CreateContactBody)
        response = requests.GetResponse(req, client)
        fmt.Println("response Message:", response.Message)

        // Save Contact ID, to use in next step
        cID := strings.Split(response.Message, "'")[1]

        // Buy the domain
        buyDomainBody := fmt.Sprintf(requests.BuyDomainBody, *domain, cID, cID, cID)
        fmt.Println("Acquiring domain...")
        req = requests.SetRequest("POST", requests.DomainEndpoint, bearerToken, buyDomainBody)
        response = requests.GetResponse(req, client)
        fmt.Println("response Message:", response.Message)
        fmt.Println("Credit:", response.Credit)
    }
    fmt.Println("Finishing...")
}
