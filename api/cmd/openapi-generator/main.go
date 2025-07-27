package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// OpenAPI 3.0.3 structure
type OpenAPISpec struct {
	OpenAPI    string                 `json:"openapi" yaml:"openapi"`
	Info       Info                   `json:"info" yaml:"info"`
	Servers    []Server               `json:"servers,omitempty" yaml:"servers,omitempty"`
	Paths      map[string]PathItem    `json:"paths" yaml:"paths"`
	Components *Components            `json:"components,omitempty" yaml:"components,omitempty"`
	Tags       []Tag                  `json:"tags,omitempty" yaml:"tags,omitempty"`
}

type Info struct {
	Title       string   `json:"title" yaml:"title"`
	Description string   `json:"description,omitempty" yaml:"description,omitempty"`
	Version     string   `json:"version" yaml:"version"`
	Contact     *Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
}

type Contact struct {
	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	URL   string `json:"url,omitempty" yaml:"url,omitempty"`
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
}

type Server struct {
	URL         string `json:"url" yaml:"url"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

type PathItem struct {
	Get    *Operation `json:"get,omitempty" yaml:"get,omitempty"`
	Post   *Operation `json:"post,omitempty" yaml:"post,omitempty"`
	Put    *Operation `json:"put,omitempty" yaml:"put,omitempty"`
	Delete *Operation `json:"delete,omitempty" yaml:"delete,omitempty"`
	Patch  *Operation `json:"patch,omitempty" yaml:"patch,omitempty"`
}

type Operation struct {
	Tags        []string             `json:"tags,omitempty" yaml:"tags,omitempty"`
	Summary     string               `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string               `json:"description,omitempty" yaml:"description,omitempty"`
	OperationID string               `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Parameters  []Parameter          `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody *RequestBody         `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Responses   map[string]Response  `json:"responses" yaml:"responses"`
}

type Parameter struct {
	Name        string  `json:"name" yaml:"name"`
	In          string  `json:"in" yaml:"in"`
	Description string  `json:"description,omitempty" yaml:"description,omitempty"`
	Required    bool    `json:"required,omitempty" yaml:"required,omitempty"`
	Schema      *Schema `json:"schema,omitempty" yaml:"schema,omitempty"`
}

type RequestBody struct {
	Description string               `json:"description,omitempty" yaml:"description,omitempty"`
	Content     map[string]MediaType `json:"content" yaml:"content"`
	Required    bool                 `json:"required,omitempty" yaml:"required,omitempty"`
}

type Response struct {
	Description string               `json:"description" yaml:"description"`
	Content     map[string]MediaType `json:"content,omitempty" yaml:"content,omitempty"`
}

type MediaType struct {
	Schema   *Schema                `json:"schema,omitempty" yaml:"schema,omitempty"`
	Examples map[string]interface{} `json:"examples,omitempty" yaml:"examples,omitempty"`
}

type Schema struct {
	Type                 string             `json:"type,omitempty" yaml:"type,omitempty"`
	Format               string             `json:"format,omitempty" yaml:"format,omitempty"`
	Items                *Schema            `json:"items,omitempty" yaml:"items,omitempty"`
	Properties           map[string]*Schema `json:"properties,omitempty" yaml:"properties,omitempty"`
	Required             []string           `json:"required,omitempty" yaml:"required,omitempty"`
	Description          string             `json:"description,omitempty" yaml:"description,omitempty"`
	Example              interface{}        `json:"example,omitempty" yaml:"example,omitempty"`
	Ref                  string             `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	AdditionalProperties interface{}        `json:"additionalProperties,omitempty" yaml:"additionalProperties,omitempty"`
}

type Components struct {
	Schemas map[string]*Schema `json:"schemas,omitempty" yaml:"schemas,omitempty"`
}

type Tag struct {
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

type SwaggerResponse struct {
	Description string  `json:"description"`
	Schema      *Schema `json:"schema,omitempty"`
}

type SwaggerSpec struct {
	Swagger     string                        `json:"swagger"`
	Info        Info                          `json:"info"`
	Host        string                        `json:"host,omitempty"`
	BasePath    string                        `json:"basePath,omitempty"`
	Schemes     []string                      `json:"schemes,omitempty"`
	Consumes    []string                      `json:"consumes,omitempty"`
	Produces    []string                      `json:"produces,omitempty"`
	Paths       map[string]SwaggerPathItem    `json:"paths"`
	Definitions map[string]*Schema            `json:"definitions,omitempty"`
	Tags        []Tag                         `json:"tags,omitempty"`
}

type SwaggerPathItem struct {
	Get    *SwaggerOperation `json:"get,omitempty"`
	Post   *SwaggerOperation `json:"post,omitempty"`
	Put    *SwaggerOperation `json:"put,omitempty"`
	Delete *SwaggerOperation `json:"delete,omitempty"`
	Patch  *SwaggerOperation `json:"patch,omitempty"`
}

type SwaggerOperation struct {
	Tags        []string                     `json:"tags,omitempty"`
	Summary     string                       `json:"summary,omitempty"`
	Description string                       `json:"description,omitempty"`
	OperationID string                       `json:"operationId,omitempty"`
	Parameters  []Parameter                  `json:"parameters,omitempty"`
	Responses   map[string]SwaggerResponse   `json:"responses"`
	Consumes    []string                     `json:"consumes,omitempty"`
	Produces    []string                     `json:"produces,omitempty"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <swagger.json>")
	}

	inputFile := os.Args[1]
	
	// Read Swagger 2.0 file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var swagger SwaggerSpec
	if err := json.Unmarshal(data, &swagger); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	// Convert to OpenAPI 3.0.3
	openapi := convertToOpenAPI(swagger)

	// Write JSON output
	jsonData, err := json.MarshalIndent(openapi, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	outputDir := filepath.Dir(inputFile)
	jsonOutputFile := filepath.Join(outputDir, "openapi3.json")
	if err := os.WriteFile(jsonOutputFile, jsonData, 0644); err != nil {
		log.Fatalf("Error writing JSON file: %v", err)
	}

	// Write YAML output
	yamlData, err := yaml.Marshal(openapi)
	if err != nil {
		log.Fatalf("Error marshaling YAML: %v", err)
	}

	yamlOutputFile := filepath.Join(outputDir, "openapi3.yaml")
	if err := os.WriteFile(yamlOutputFile, yamlData, 0644); err != nil {
		log.Fatalf("Error writing YAML file: %v", err)
	}

	fmt.Printf("✅ Successfully converted Swagger 2.0 to OpenAPI 3.0.3\n")
	fmt.Printf("📁 JSON output: %s\n", jsonOutputFile)
	fmt.Printf("📁 YAML output: %s\n", yamlOutputFile)
}

func convertToOpenAPI(swagger SwaggerSpec) OpenAPISpec {
	openapi := OpenAPISpec{
		OpenAPI: "3.0.3",
		Info:    swagger.Info,
		Paths:   convertPaths(swagger.Paths),
		Tags:    swagger.Tags,
	}

	// Convert host and basePath to servers
	if swagger.Host != "" {
		servers := []Server{}
		schemes := swagger.Schemes
		if len(schemes) == 0 {
			schemes = []string{"http"}
		}
		
		for _, scheme := range schemes {
			url := fmt.Sprintf("%s://%s", scheme, swagger.Host)
			if swagger.BasePath != "" && swagger.BasePath != "/" {
				url += swagger.BasePath
			}
			servers = append(servers, Server{
				URL:         url,
				Description: fmt.Sprintf("%s server", strings.Title(scheme)),
			})
		}
		openapi.Servers = servers
	}

	// Convert definitions to components/schemas
	if len(swagger.Definitions) > 0 {
		openapi.Components = &Components{
			Schemas: swagger.Definitions,
		}
	}

	return openapi
}

func convertPaths(paths map[string]SwaggerPathItem) map[string]PathItem {
	converted := make(map[string]PathItem)
	
	for path, pathItem := range paths {
		convertedPathItem := PathItem{}
		
		if pathItem.Get != nil {
			convertedPathItem.Get = convertSwaggerOperation(*pathItem.Get)
		}
		if pathItem.Post != nil {
			convertedPathItem.Post = convertSwaggerOperation(*pathItem.Post)
		}
		if pathItem.Put != nil {
			convertedPathItem.Put = convertSwaggerOperation(*pathItem.Put)
		}
		if pathItem.Delete != nil {
			convertedPathItem.Delete = convertSwaggerOperation(*pathItem.Delete)
		}
		if pathItem.Patch != nil {
			convertedPathItem.Patch = convertSwaggerOperation(*pathItem.Patch)
		}
		
		converted[path] = convertedPathItem
	}
	
	return converted
}

func convertSwaggerOperation(op SwaggerOperation) *Operation {
	converted := &Operation{
		Tags:        op.Tags,
		Summary:     op.Summary,
		Description: op.Description,
		OperationID: op.OperationID,
		Parameters:  convertParameters(op.Parameters),
		Responses:   convertSwaggerResponses(op.Responses),
	}
	
	return converted
}

func convertSwaggerResponses(responses map[string]SwaggerResponse) map[string]Response {
	converted := make(map[string]Response)
	
	for code, response := range responses {
		convertedResponse := Response{
			Description: response.Description,
		}
		
		// Convert Swagger 2.0 schema to OpenAPI 3.0.3 content
		if response.Schema != nil {
			mediaType := MediaType{
				Schema: response.Schema,
			}
			
			// Add examples based on schema references
			if response.Schema.Ref != "" {
				examples := createExampleFromRef(response.Schema.Ref)
				if examples != nil {
					mediaType.Examples = examples
				}
			} else if response.Schema.Type == "array" && response.Schema.Items != nil && response.Schema.Items.Ref != "" {
				// Handle array responses
				examples := createArrayExampleFromRef(response.Schema.Items.Ref)
				if examples != nil {
					mediaType.Examples = examples
				}
			} else if response.Schema.Type == "object" && response.Schema.AdditionalProperties != nil {
				// Handle generic map objects (like health endpoint)
				examples := map[string]interface{}{
					"health_response": map[string]interface{}{
						"summary": "Health check response",
						"value": map[string]interface{}{
							"status":    "healthy",
							"timestamp": "2024-01-01T12:00:00Z",
							"service":   "template-mobile-app-api",
							"version":   "1.0.0",
						},
					},
				}
				mediaType.Examples = examples
			}
			
			convertedResponse.Content = map[string]MediaType{
				"application/json": mediaType,
			}
		}
		
		converted[code] = convertedResponse
	}
	
	return converted
}

func convertParameters(params []Parameter) []Parameter {
	converted := make([]Parameter, len(params))
	copy(converted, params)
	return converted
}

func createExampleFromRef(ref string) map[string]interface{} {
	switch ref {
	case "#/definitions/main.PostResponse":
		return map[string]interface{}{
			"single_post": map[string]interface{}{
				"summary": "Single post response",
				"value": map[string]interface{}{
					"id":          1,
					"title":       "Sample Post Title",
					"body":        "Sample post body content",
					"userId":      1,
					"formattedAt": "2024-01-01T12:00:00Z",
				},
			},
		}
	case "#/definitions/main.ErrorResponse":
		return map[string]interface{}{
			"error_response": map[string]interface{}{
				"summary": "Error response",
				"value": map[string]interface{}{
					"error":   "Internal server error",
					"message": "Failed to fetch data",
				},
			},
		}
	}
	return nil
}

func createArrayExampleFromRef(ref string) map[string]interface{} {
	switch ref {
	case "#/definitions/main.PostResponse":
		return map[string]interface{}{
			"posts_list": map[string]interface{}{
				"summary": "List of posts",
				"value": []interface{}{
					map[string]interface{}{
						"id":          1,
						"title":       "First Post Title",
						"body":        "This is the content of the first post",
						"userId":      1,
						"formattedAt": "2024-01-01T12:00:00Z",
					},
					map[string]interface{}{
						"id":          2,
						"title":       "Second Post Title",
						"body":        "This is the content of the second post",
						"userId":      2,
						"formattedAt": "2024-01-01T12:00:00Z",
					},
				},
			},
		}
	}
	return nil
}