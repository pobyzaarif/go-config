# go-config
Golang config loader to parse environment variables into structs
Using [https://github.com/joho/godotenv](joho/godotenv) and [https://github.com/caarlos0/env](caarlos0/env) as base.

## Implementation
Just take a look at the test file ( *_test.go )

## Run test
```
go test -count=1 -coverprofile=coverage.out -coverpkg=./... && go tool cover -html=coverage.out -o cover.html
```
