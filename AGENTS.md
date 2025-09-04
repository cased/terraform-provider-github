# AGENTS.md

## Project Overview
This is the Terraform Provider for GitHub (Cased fork) - a Go-based Terraform provider for managing GitHub resources. The project uses Go modules, Make for build automation, and golangci-lint for code quality.

## Setup Commands
- Install Go tools: `make tools`
- Build the project: `make build`
- Run tests: `make test`
- Run acceptance tests: `make testacc` (requires GitHub token and organization)
- Format code: `make fmt`
- Lint code: `make lint` (requires PATH to include `$GOPATH/bin`)

## Environment Setup
```bash
# Add Go binaries to PATH (required for linting tools)
export PATH=$PATH:$(go env GOPATH)/bin

# For testing (required for acceptance tests)
export GITHUB_TOKEN=<your-github-token>
export GITHUB_ORGANIZATION=<your-test-organization>
```

## Code Style & Standards
- **Language**: Go 1.21+ with toolchain go1.22.0
- **Formatting**: Uses `gofmt` with `-s` flag (simplify code)
- **Linting**: golangci-lint with specific rules (see `.golangci.yml`)
- **Testing**: Uses `testify` for assertions
- **Module**: `github.com/integrations/terraform-provider-github/v6`

## Key Linters Enabled
- `errcheck`: Check for unchecked errors
- `gofmt`: Enforce formatting standards
- `gosimple`: Suggest code simplifications
- `ineffassign`: Detect ineffective assignments
- `staticcheck`: Advanced static analysis
- `vet`: Standard Go vet checks
- `misspell`: Catch spelling mistakes

## Development Workflow
1. **Format before building**: Always run `make fmt` before `make build`
2. **Build before testing**: `make build` runs format check first
3. **Use acceptance tests**: Most meaningful tests require GitHub API access
4. **Dependencies**: Managed with `go mod` - run `go mod tidy` to clean up

## Testing Instructions
- **Unit tests**: `make test` - Fast, no external dependencies
- **Acceptance tests**: `make testacc` - Requires GitHub token and test organization
- **Individual tests**: `go test -run TestName ./package`
- **Debug tests**: `TF_LOG=DEBUG TF_ACC=1 go test -v ./... -run ^TestName`
- **Test requirements**: Test organization must have a `terraform-template-module` repository marked as template

## Project Structure
- `main.go` - Entry point for the Terraform provider
- `github/` - Main source code with resources and data sources
- `scripts/` - Utility scripts (format checking, etc.)
- `examples/` - Example Terraform configurations
- `website/` - Documentation (if present)
- `.github/workflows/` - CI/CD pipeline definitions

## Common Commands
```bash
# Full development cycle
make tools      # Install required tools
make fmt        # Format code
make build      # Build the provider
make test       # Run unit tests
go vet ./...    # Additional static analysis

# Dependency management
go mod tidy     # Clean up dependencies
go mod download # Download dependencies

# Alternative linting (if golangci-lint has issues)
go vet ./...    # Built-in Go static analysis
```

## CI/CD Pipeline
The project uses GitHub Actions with the following steps:
1. Install Go tools (`make tools`)
2. Lint code (`make lint`)
3. Lint website (`make website-lint`) 
4. Build project (`make build`)
5. Run tests (`make test`)

## Troubleshooting
- **Lint fails**: Ensure `$GOPATH/bin` is in your PATH
- **Format check fails**: Run `make fmt` before building
- **Tests fail**: Ensure `GITHUB_TOKEN` and `GITHUB_ORGANIZATION` are set for acceptance tests
- **Build fails**: Check Go version (requires 1.21+) and run `go mod tidy`

## Development Tips
- Most resources are in the `github/` directory with pattern `resource_github_*.go`
- Test files follow pattern `*_test.go` 
- Use `make fmt` frequently to maintain code style
- The project uses Terraform Plugin SDK v2
- Focus on acceptance tests for meaningful validation
- Check existing issues in [PRIORITY_ISSUES.md](PRIORITY_ISSUES.md) for contribution ideas

## PR Guidelines
- Run full test suite before submitting: `make tools && make fmt && make build && make test`
- Include tests for new functionality
- Follow existing code patterns in the `github/` directory
- Update documentation if adding new resources or data sources