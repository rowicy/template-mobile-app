#!/bin/bash

# OpenAPI 3.0.3 Generation Script - Pure Go Implementation
# This script generates OpenAPI 3.0.3 documentation using only Go packages

set -e

echo "🔄 Generating OpenAPI 3.0.3 documentation using Go packages..."

# Ensure we're in the API directory
cd "$(dirname "$0")"

# Ensure swag is in PATH
export PATH=$PATH:$(go env GOPATH)/bin

# Install swag if not available
if ! command -v swag &> /dev/null; then
    echo "📦 Installing swag..."
    go install github.com/swaggo/swag/cmd/swag@latest
fi

# Update dependencies (ensure yaml.v3 is available)
echo "📦 Updating dependencies..."
go mod tidy

# Generate initial Swagger 2.0 docs
echo "📖 Generating Swagger 2.0 documentation..."
swag init

# Convert to OpenAPI 3.0.3 using our Go converter
echo "🔄 Converting to OpenAPI 3.0.3 using Go converter..."
go run cmd/openapi-generator/main.go docs/swagger.json

# Replace original files with OpenAPI 3.0.3 versions
echo "🔄 Updating files..."
cp docs/openapi3.json docs/swagger.json
cp docs/openapi3.yaml docs/swagger.yaml

# Update docs.go to reflect OpenAPI 3.0.3
echo "🔄 Updating docs.go for OpenAPI 3.0.3..."
sed -i 's/"swagger": "2.0"/"openapi": "3.0.3"/g' docs/docs.go
sed -i 's/Swagger: "2.0"/OpenAPI: "3.0.3"/g' docs/docs.go

# Fix swag.Spec structure for newer versions
sed -i '/LeftDelim:/d' docs/docs.go
sed -i '/RightDelim:/d' docs/docs.go

# Update openapi-specifications directory
cp docs/openapi3.json ../openapi-specifications/api.swagger.json

# Clean up temporary and backup files
echo "🧹 Cleaning up temporary files..."
rm -f docs/openapi3.json docs/openapi3.yaml
rm -f docs/*.backup
rm -f ../openapi-specifications/*.backup

echo "✅ OpenAPI 3.0.3 documentation generated successfully using Go packages!"
echo "📁 Files updated:"
echo "   - docs/swagger.json (OpenAPI 3.0.3)"
echo "   - docs/swagger.yaml (OpenAPI 3.0.3)"
echo "   - docs/docs.go (OpenAPI 3.0.3)"
echo "   - openapi-specifications/api.swagger.json (OpenAPI 3.0.3)"
echo ""
echo "💡 Swagger UI available at: http://localhost:8080/swagger/index.html"
echo "🎯 100% Go-based solution - no npm dependencies required!"