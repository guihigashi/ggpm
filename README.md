# ggpm

Declaratively manage and install Go binaries with custom build tags using a simple configuration file.

## Features

- **Declarative Configuration**: Define all your global Go packages and their corresponding tags in a single YAML file.
- **Tag Support**: Pass custom build tags (`-tags`) to `go install` per package.
- **Flexible Path Resolution**: Automatically detects configuration files in the current working directory or your home
  directory (`$HOME` / `~`).
- **Visual Progress**: Clean CLI feedback with real-time logging and spinners.

## Configuration

Create a `.ggpm` (or `.ggpm.yml`) file in your home directory or project root:

```yaml
packages:
  - module: github.com/guihigashi/ggpm/cmd/ggpm@latest
  - module: github.com/air-verse/air@latest
  - module: github.com/auth0/auth0-cli/cmd/auth0@latest
  - module: github.com/golang-migrate/migrate/v4/cmd/migrate@latest
    tags: 'pgx5,clickhouse'
  - module: google.golang.org/protobuf/cmd/protoc-gen-go@latest
  - module: google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```