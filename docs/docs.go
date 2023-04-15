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
        "/auth/refresh": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Refreshing a user session",
                "parameters": [
                    {
                        "description": "User session",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/docs.RefreshSessionDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "New data for user authorization",
                        "schema": {
                            "$ref": "#/definitions/docs.RefreshResponse"
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/docs.MessageResponse"
                        }
                    }
                }
            }
        },
        "/auth/sign-in": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "User sign in",
                "parameters": [
                    {
                        "description": "User JSON",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.GetUserIDDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Data for user authorization",
                        "schema": {
                            "$ref": "#/definitions/docs.SignInResponse"
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/docs.MessageResponse"
                        }
                    },
                    "500": {
                        "description": "Server error message",
                        "schema": {
                            "$ref": "#/definitions/docs.MessageResponse"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "User sign up",
                "parameters": [
                    {
                        "description": "User JSON",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.CreateUserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User ID",
                        "schema": {
                            "$ref": "#/definitions/docs.SignUpResponse"
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/docs.MessageResponse"
                        }
                    }
                }
            }
        },
        "/notebooks": {
            "get": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notebook"
                ],
                "summary": "Getting list of user notebooks",
                "responses": {
                    "200": {
                        "description": "Notebooks",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/docs.GetAllNotebooksResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/docs.MessageResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notebook"
                ],
                "summary": "Creating notebook",
                "parameters": [
                    {
                        "description": "Notebook data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/docs.CreateUpdateNotebookDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Notebook ID",
                        "schema": {
                            "$ref": "#/definitions/docs.CreateNotebookResponse"
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/docs.MessageResponse"
                        }
                    }
                }
            }
        },
        "/notebooks/{notebook_id}": {
            "put": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notebook"
                ],
                "summary": "Updating user notebook",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Notebook ID",
                        "name": "notebook_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "New notebook data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/docs.CreateUpdateNotebookDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK status",
                        "schema": {
                            "$ref": "#/definitions/docs.OkStatusResponse"
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/docs.MessageResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notebook"
                ],
                "summary": "Deleting user notebook",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Notebook ID",
                        "name": "notebook_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK status",
                        "schema": {
                            "$ref": "#/definitions/docs.OkStatusResponse"
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/docs.MessageResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "docs.CreateNotebookResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "docs.CreateUpdateNotebookDTO": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                }
            }
        },
        "docs.GetAllNotebooksResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "docs.MessageResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "docs.OkStatusResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "ok"
                }
            }
        },
        "docs.RefreshResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "docs.RefreshSessionDTO": {
            "type": "object",
            "properties": {
                "refresh_token": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "docs.SignInResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "docs.SignUpResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "domain.CreateUserDTO": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "domain.GetUserIDDTO": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Notes API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
