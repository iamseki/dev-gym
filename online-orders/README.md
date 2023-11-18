# Sistema desacoplado de pedidos online

> Challenge description: https://app.devgym.com.br/challenges/b7e7157e-1e85-40f3-9c24-032cd6a45aec

## Running the challenge

Since there is no interface required, we can validate the behavior of business rules just using unit tests:

- `go test ./... -coverprofile=coverage.out`
- `go tool cover -html=coverage.out`
