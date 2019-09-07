package requests

var tokenEndpointResponse = `
{
    "scopes": [
        "POST:test.domains.altravia.com\/domain",
        "POST:test.domains.altravia.com\/contact",
        "GET:test.domains.altravia.com\/check"
    ],
    "expire": 1599073035,
    "token": "5d6eb78bed92ef09cb60699a",
    "success": true,
    "message": "",
    "error": null
}`

var contactEndpointResponse = `
{
    "message": "Greeting OK. Login OK. Create contact 'AV-517176925' created. Logout OK.",
    "credit": "50.00",
    "success": true,
    "error": null
}`

var checkEndpointResponse = `
{
    "message": "Greeting OK. Login OK. Domain 'test-domain.it' is available. Logout OK.",
    "credit": "50.00",
    "success": true,
    "error": null
}`

var buyDomainEndpointResponse = `
{
    "message": "Greeting OK. Login OK. Domain 'test-domain.it' is available. Domain 'test-domain.it' created. Logout OK.",
    "credit": "40.00",
    "success": true,
    "error": null
}`
