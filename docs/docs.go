// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/custom_type": {
            "get": {
                "description": "returns a response with custom struct",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "custom_struct"
                ],
                "summary": "Sample response",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.sampleResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.CustomType": {
            "type": "object",
            "properties": {
                "customField": {
                    "type": "object"
                }
            }
        },
        "main.sampleResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/main.CustomType"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Swaggo API",
	Description:      "Sample API App to reproduce the bug",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}