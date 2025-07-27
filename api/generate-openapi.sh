#!/bin/bash

# OpenAPI 3.0.3 Generation Script
# This script converts Swagger 2.0 to OpenAPI 3.0.3 using swagger2openapi

set -e

echo "🔄 Converting Swagger 2.0 to OpenAPI 3.0.3..."

# Ensure we're in the API directory
cd "$(dirname "$0")"

# Install swagger2openapi if not available
if ! command -v swagger2openapi &> /dev/null; then
    echo "📦 Installing swagger2openapi..."
    npm install -g swagger2openapi
fi

# Ensure swag is in PATH
export PATH=$PATH:$(go env GOPATH)/bin

# Install swag if not available
if ! command -v swag &> /dev/null; then
    echo "📦 Installing swag..."
    go install github.com/swaggo/swag/cmd/swag@latest
fi

# Generate initial Swagger 2.0 docs
echo "📖 Generating Swagger 2.0 documentation..."
swag init

# Convert to OpenAPI 3.0.3
echo "🔄 Converting to OpenAPI 3.0.3..."
swagger2openapi docs/swagger.json -o docs/openapi3.json
swagger2openapi docs/swagger.yaml -o docs/openapi3.yaml

# Update version to 3.0.3
sed -i 's/"openapi": "3.0.0"/"openapi": "3.0.3"/g' docs/openapi3.json
sed -i 's/openapi: 3.0.0/openapi: 3.0.3/g' docs/openapi3.yaml

# Replace original files
cp docs/swagger.json docs/swagger.json.backup
cp docs/swagger.yaml docs/swagger.yaml.backup
cp docs/openapi3.json docs/swagger.json
cp docs/openapi3.yaml docs/swagger.yaml

# Update openapi-specifications directory
cp docs/openapi3.json ../openapi-specifications/api.swagger.json

echo "✅ OpenAPI 3.0.3 documentation generated successfully!"
echo "📁 Files updated:"
echo "   - docs/swagger.json (OpenAPI 3.0.3)"
echo "   - docs/swagger.yaml (OpenAPI 3.0.3)"
echo "   - openapi-specifications/api.swagger.json (OpenAPI 3.0.3)"
echo ""
echo "💡 Swagger UI available at: http://localhost:8080/swagger/index.html"