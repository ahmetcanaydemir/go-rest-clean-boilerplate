// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "produces": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/Stats": {
            "get": {
                "description": "This endpoint returns the status of the Go Rest Clean Boilerplate",
                "tags": [
                    "Stats"
                ],
                "summary": "Get Stats",
                "responses": {
                    "200": {
                        "description": "Stats",
                        "schema": {
                            "$ref": "#/definitions/model.Stats"
                        }
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/nginx-health": {
            "get": {
                "description": "This function returns the nginx health",
                "tags": [
                    "Status"
                ],
                "summary": "Get Nginx Health",
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/status": {
            "get": {
                "description": "This endpoint returns the status of the Go Rest Clean Boilerplate",
                "tags": [
                    "Status"
                ],
                "summary": "Get Status",
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Stats": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "number"
                },
                "date": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Go Rest Clean Boilerplate",
	Description:      "Go Rest Clean Boilerplate",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
