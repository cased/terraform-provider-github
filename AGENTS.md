# AGENTS.md

## Environment

### Setup script

```bash
go install golang.org/dl/go1.22.0@latest
go1.22.0 download
export GOTOOLCHAIN=go1.22.0
export PATH="$PATH:$(go env GOPATH)/bin"
make tools
```

### Maintenance script

```bash
export GOTOOLCHAIN=go1.22.0
export PATH="$PATH:$(go env GOPATH)/bin"
make tools
```

## Build

- Build all packages: `go build ./...`

## Testing

- Run unit tests: `make test`
- Run acceptance tests (require GitHub credentials):
  ```bash
  export GITHUB_TOKEN=<token>
  export GITHUB_ORGANIZATION=<org>
  make testacc
  ```

## Lint

- Run linter: `make lint`
  - The linter may report existing issues.

## Code style

- Format code: `make fmt`
- Verify formatting: `make fmtcheck`
