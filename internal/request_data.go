// Package requests ...
// This file contains endpoint and body information
// about the requests the application is going to make
// against Altravia's test api.
// The request bodies and endpoints listed here are correct
// Nevertheless it is strongly suggested to verify that against the docs
// https://test.api.altravia.com/
package requests

// TokenEndPoint ...
var TokenEndPoint = "https://test.oauth.altravia.com/token/"

// ContactEndpoint ...
var ContactEndpoint = "https://test.domains.altravia.com/contact"

// CheckEndpoint needs to be interpolated with domain name
var CheckEndpoint = "https://test.domains.altravia.com/check/%s"

// DomainEndpoint ...
var DomainEndpoint = "https://test.domains.altravia.com/domain"

// OauthTokenBody defines the scopes to obtain tokens for
var OauthTokenBody = `
{"scopes":[
"POST:test.domains.altravia.com/domain",
"GET:test.domains.altravia.com/check",
"POST:test.domains.altravia.com/contact"
]}`

// CreateContactBody defines the contact body
// This data will be used as registrant information for your
// purchased domain. Edit this!
var CreateContactBody = `
{"name":"John Gopher",
"org":"John Gopher",
"street":"Via Freedom 1008",
"city":"Roma",
"province":"RM",
"postalcode":"00187",
"countrycode":"IT",
"voice":"+39.00000000000",
"email":"john@gopher.biz",
"nationalitycode":"IT",
"entitytype":1,
"regcode":"JHGPHR7238247U2A"}
`

// BuyDomainBody interpolate domain name and contact ID here
var BuyDomainBody = `
{"domain":"%s",
  "registrant":"%s",
  "admin":"%s",
  "tech":["%s"],
  "dns":["dns1.altravia.com","dns2.altravia.com","dns3.altravia.com"]}
`
