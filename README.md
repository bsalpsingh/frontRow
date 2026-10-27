# frontRow
  myapp/
  ├── cmd/
  │   └── api/
  │       └── main.go          # entrypoint, wiring only
  ├── internal/                # private code, can't be imported by other modules
  │   ├── config/              # env/config loading
  │   ├── handler/             # HTTP handlers (transport layer)
  │   ├── service/             # business logic
  │   ├── repository/          # DB access
  │   ├── model/               # domain types
  │   └── middleware/
  ├── pkg/                     # optional: code safe for others to import
  ├── migrations/              # SQL migrations
  ├── api/                     # OpenAPI specs / proto files
  ├── scripts/                 # build/dev helper scripts
  ├── Dockerfile
  ├── Makefile
  ├── go.mod
  └── go.sum