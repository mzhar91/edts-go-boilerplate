// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Harry K",
            "email": "k.harry791@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth": {
            "post": {
                "description": "Add a new credential for user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Add Credential",
                "parameters": [
                    {
                        "description": "Add Credential Request",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddCredentialRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/healthcheck": {
            "get": {
                "description": "Health check endpoint",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health Check"
                ],
                "summary": "Health Check",
                "responses": {}
            }
        }
    },
    "definitions": {
        "model.AddCredentialRequest": {
            "type": "object",
            "required": [
                "claim",
                "password",
                "scope",
                "username"
            ],
            "properties": {
                "claim": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "requestInfo": {
                    "$ref": "#/definitions/model.RequestInfo"
                },
                "scope": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.RequestInfo": {
            "type": "object",
            "properties": {
                "host": {
                    "type": "string"
                },
                "ipNumber": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3010",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger EDTS go boilerplate API",
	Description:      "This is a sample server Petstore server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
