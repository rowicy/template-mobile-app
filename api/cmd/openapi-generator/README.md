# OpenAPI Generator

This directory contains a pure Go implementation for converting Swagger 2.0 specifications to OpenAPI 3.0.3 format.

## Overview

The `main.go` file in this directory provides a command-line tool that:

1. Reads Swagger 2.0 JSON files
2. Converts them to OpenAPI 3.0.3 format
3. Outputs both JSON and YAML formats

## Usage

```bash
go run cmd/openapi-generator/main.go docs/swagger.json
```

This will generate:
- `docs/openapi3.json` - OpenAPI 3.0.3 in JSON format
- `docs/openapi3.yaml` - OpenAPI 3.0.3 in YAML format

## Features

- **Pure Go implementation** - No npm or Node.js dependencies
- **Full conversion** - Handles all common Swagger 2.0 constructs
- **Multiple formats** - Outputs both JSON and YAML
- **Fast and reliable** - Native Go performance

## Integration

This converter is automatically used by the `generate-openapi-go.sh` script in the parent directory to provide a complete Go-only solution for OpenAPI 3.0.3 generation.