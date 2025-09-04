# AGENTS.md

## Setup commands
- Install Go tools: `make tools`
- Install dependencies: `go mod download`
- Set up PATH: `export PATH=$PATH:~/go/bin` (for linting tools)
- Upgrade golangci-lint (if compatibility issues): `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`

## Build commands
- Build provider: `make build`
- Format code: `make fmt`
- Check formatting: `make fmtcheck`

## Testing commands
- Run unit tests: `make test`
- Run acceptance tests: `make testacc` (requires GitHub token)
- Run specific test: `TF_ACC=1 go test -v ./github -run TestAccGithubRepository`
- Run with debug logging: `TF_LOG=DEBUG TF_ACC=1 go test -v ./github -run TestAccGithubRepository`

## Code quality
- Format check: `make fmtcheck`
- Format fix: `make fmt`
- Go vet: `make vet`
- Lint (requires PATH setup): `export PATH=$PATH:~/go/bin && make lint`
- Website lint: `make website-lint`

## Environment variables for testing
```bash
# Required for acceptance tests
export GITHUB_TOKEN=<your_github_token>
export GITHUB_ORGANIZATION=<test_organization>

# Optional for debugging
export TF_LOG=DEBUG
export TF_ACC=1

# For local provider development
export TF_CLI_CONFIG_FILE=path/to/project/examples/dev.tfrc
```

## Development workflow
- Always run `make fmt` before committing
- Use `make build` to verify compilation
- Run `make test` for unit tests before push
- Set up test organization with `terraform-template-module` repository
- Use examples/ directory for manual testing with local builds
- Example terraform files available in `examples/*/`

## Local development setup
1. Build provider: `go build -gcflags="all=-N -l" -o ~/go/bin/`
2. Configure dev override: Use `examples/dev.tfrc`
3. Set config file: `export TF_CLI_CONFIG_FILE=path/to/project/examples/dev.tfrc`
4. Run terraform commands in examples/ directory

## Troubleshooting
- **Formatting errors**: Run `make fmt` to auto-fix
- **Linting issues**: 
  - Ensure `export PATH=$PATH:~/go/bin` is set
  - If golangci-lint version issues: `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`
- **Test failures**: Check GitHub token and organization setup
- **Build errors**: Run `go mod tidy` to clean dependencies
- **Unit test example**: `go test -v ./github -run TestRetryTransport`

## Useful commands for exploration
- List all make targets: `make help` (if available) or check `GNUmakefile`
- Find test functions: `grep -r "func TestAcc" github/ | head -5`
- List example directories: `ls examples/`
- Check Go version: `go version`
- Clean build cache: `go clean -cache`
- Update all dependencies: `go get -u ./...`

## Code style
- Go 1.21+ required
- Uses golangci-lint for linting
- Standard Go formatting with gofmt
- Acceptance tests require real GitHub resources
- Use `log.Printf("[DEBUG] message")` for debug output when TF_LOG=DEBUG