# placelists-back

project-root/
    ├── cmd/
    │   ├── your-app-name/
    │   │   ├── main.go         # Application entry point
    │   │   └── ...             # Other application-specific files
    │   └── another-app/
    │       ├── main.go         # Another application entry point
    │       └── ...
    ├── internal/                # Private application and package code
    │   ├── config/
    │   │   ├── config.go       # Configuration logic
    │   │   └── ...
    │   ├── database/
    │   │   ├── database.go     # Database setup and access
    │   │   └── ...
    │   └── ...
    ├── pkg/                     # Public, reusable packages
    │   ├── mypackage/
    │   │   ├── mypackage.go    # Public package code
    │   │   └── ...
    │   └── ...
    ├── api/                     # API-related code (e.g., REST or gRPC)
    │   ├── handler/
    │   │   ├── handler.go      # HTTP request handlers
    │   │   └── ...
    │   ├── middleware/
    │   │   ├── middleware.go  # Middleware for HTTP requests
    │   │   └── ...
    │   └── ...
    ├── web/                     # Front-end web application assets
    │   ├── static/
    │   │   ├── css/
    │   │   ├── js/
    │   │   └── ...
    │   └── templates/
    │       ├── index.html
    │       └── ...
    ├── scripts/                 # Build, deployment, and maintenance scripts
    │   ├── build.sh
    │   ├── deploy.sh
    │   └── ...
    ├── configs/                 # Configuration files for different environments
    │   ├── development.yaml
    │   ├── production.yaml
    │   └── ...
    ├── tests/                   # Unit and integration tests
    │   ├── unit/
    │   │   ├── ...
    │   └── integration/
    │       ├── ...
    ├── docs/                    # Project documentation
    ├── .gitignore               # Gitignore file
    ├── go.mod                   # Go module file
    ├── go.sum                   # Go module dependencies file
    └── README.md                # Project README


golang-hexagon-example/
├── cmd/
│   └── main.go
├── configs/
│   └── config.json
├── internal/
│   ├── app/
│   │   ├── adapters/
│   │   │   └── cache/
│   │   │       └── redis_client
│   │   ├── handlers/
│   │   │   └── kafka/
│   │   │       ├── module.go
│   │   │       └── service.go
│   │   │   └── api/
│   │   │       ├── module.go
│   │   │       └── controller.go
│   │   ├── repositories/
│   │   │   └── mongodb/
│   │   │       ├── connection.go
│   │   │       ├── module.go
│   │   │       └── repository.go
│   │   └── module.go
│   ├── core/
│   │   ├── models/
│   │   │   └── models.go
│   │   ├── ports/
│   │   │   └── ports.go
│   │   └── services/
│   │       ├── domain_service.go
│   │       └── module.go
│   ├── infrastructure/
│   │   ├── config/
│   │   │   └── config.go
│   │   └── server/
│   │       └── echo_server.go
│   └── app_module.go
├── go.mod
├── makefile
└── README.md

https://medium.com/@shershnev/layered-architecture-implementation-in-golang-6318a72c1e10