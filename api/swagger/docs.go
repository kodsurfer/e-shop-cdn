// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "/",
        "contact": {
            "name": "mail",
            "url": "/",
            "email": "kartashov_egor96@mail.ru"
        },
        "license": {
            "name": "MIT",
            "url": "http://www.apache.org/licenses/MIT.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/cdn/delete/{id}": {
            "delete": {
                "description": "delete file by name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Files Controller"
                ],
                "summary": "Allow delete file",
                "parameters": [
                    {
                        "type": "string",
                        "description": "123",
                        "name": "x-api-key",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Filenam",
                        "name": "filename",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/cdn/download/{filename}": {
            "get": {
                "description": "download file by name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Files Controller"
                ],
                "summary": "Allow \t\t\t\t\t\tdownload file",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Filenam",
                        "name": "filename",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/cdn/files": {
            "post": {
                "description": "show paginated files",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Files Controller"
                ],
                "summary": "Allow get paginated files",
                "parameters": [
                    {
                        "type": "string",
                        "description": "123",
                        "name": "x-api-key",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/cdn/metadata/{filename}": {
            "get": {
                "description": "show file metadata",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Files Controller"
                ],
                "summary": "Allow get file metadata",
                "parameters": [
                    {
                        "type": "string",
                        "description": "123",
                        "name": "x-api-key",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Filename",
                        "name": "page",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/cdn/upload": {
            "post": {
                "description": "upload files",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Upload Controller"
                ],
                "summary": "Allow upload multiple files",
                "parameters": [
                    {
                        "type": "string",
                        "description": "123",
                        "name": "x-api-key",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Files",
                        "name": "files",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/livez": {
            "get": {
                "description": "Health check service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health Controller"
                ],
                "summary": "Health check service",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/readyz": {
            "get": {
                "description": "Ready check service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health Controller"
                ],
                "summary": "Ready check service",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8888",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "eShopCDN Swagger Doc",
	Description:      "eShopCDN",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
