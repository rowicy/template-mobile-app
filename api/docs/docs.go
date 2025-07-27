// Package docs OpenAPI 3.0.3 documentation
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "openapi": "3.0.3",
    "info": {
        "description": "This is a template API server using Go Gin framework",
        "title": "Template Mobile App API",
        "contact": {},
        "version": "1.0"
    },
    "paths": {
        "/health": {
            "get": {
                "description": "Returns the health status of the API",
                "tags": [
                    "health"
                ],
                "summary": "Health check endpoint",
                "responses": {
                    "200": {
                        "description": "OK",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "type": "object",
                                    "additionalProperties": true
                                }
                            }
                        }
                    }
                }
            }
        },
        "/posts": {
            "get": {
                "description": "Fetch all posts from JSONPlaceholder API and return formatted response",
                "tags": [
                    "posts"
                ],
                "summary": "Get all posts from JSONPlaceholder",
                "responses": {
                    "200": {
                        "description": "OK",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "type": "array",
                                    "items": {
                                        "$ref": "#/components/schemas/main.PostResponse"
                                    }
                                }
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/main.ErrorResponse"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/posts/{id}": {
            "get": {
                "description": "Fetch a specific post by ID from JSONPlaceholder API and return formatted response",
                "tags": [
                    "posts"
                ],
                "summary": "Get a post by ID from JSONPlaceholder",
                "parameters": [
                    {
                        "description": "Post ID",
                        "name": "id",
                        "in": "path",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/main.PostResponse"
                                }
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/main.ErrorResponse"
                                }
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/main.ErrorResponse"
                                }
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/main.ErrorResponse"
                                }
                            }
                        }
                    }
                }
            }
        }
    },
    "servers": [
        {
            "url": "http://localhost:8080/api/v1"
        }
    ],
    "components": {
        "schemas": {
            "main.ErrorResponse": {
                "type": "object",
                "properties": {
                    "error": {
                        "type": "string",
                        "example": "Internal server error"
                    },
                    "message": {
                        "type": "string",
                        "example": "Failed to fetch data"
                    }
                }
            },
            "main.PostResponse": {
                "type": "object",
                "properties": {
                    "body": {
                        "type": "string",
                        "example": "Sample post body content"
                    },
                    "formattedAt": {
                        "type": "string",
                        "example": "2024-01-01T12:00:00Z"
                    },
                    "id": {
                        "type": "integer",
                        "example": 1
                    },
                    "title": {
                        "type": "string",
                        "example": "Sample Post Title"
                    },
                    "userId": {
                        "type": "integer",
                        "example": 1
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Template Mobile App API",
	Description:      "This is a template API server using Go Gin framework",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}