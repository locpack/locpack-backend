# locpack-backend

## Swagger

1. Generate OpenAPI schema.
    ```bash
   swag init -output ./docs/swagger -g ./cmd/locpack-backend/main.go --parseDependency
    ```

## Testing

1. Initialize mocks.
    ```bash
    mockery
    ```
2. Run tests.
    ```bash
    go test -cover ./...
    ```
