# Check Domain
Play around with Altravia srl Test API for fun but also for profit

## The Problem
We want to have a small, fast program that:
- Checks for purchase availability of an italian internet domain.
- If the domain is available, purchase it.

We want to run this program as a scheduled task, for example as a cronjob on a Linux server,
where it will be run every hour or so.

## The Constraints
- The program should be small, lightweight and require minimal dependencies to run.
- We should use Altravia's API. Altravia is a web registrar whom has given us permission to use their api.
Credentials are available, and also [documentation](https://test.api.altravia.com/docs/).
- We should run it responsibly, without spamming Altravia srl's api.
- The application should be written in Go, with TDD approach.
- Deadline for a working product is 10th of September 2019.

## Requirements

Obtain API credentials, and store them in an `internal/.env` file:

```
APIKEY=759827398275345345FDGSRTHSRTH5345345345DGDFGDFGDFG
USERNAME=john@gopher.com
```

Check `internal/request_data.go` for validity of information.

## Install

```
go install github.com/sk1u/checkdomain/cmd/checkdomain
```

## Run
```
checkdomain
checkdomain -d example.it
```

## Test
```
go test -v ./internal
```
