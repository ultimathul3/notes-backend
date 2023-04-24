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
        "/auth/logout": {
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
                    "Auth"
                ],
                "summary": "User logout",
                "parameters": [
                    {
                        "description": "User refresh token",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.LogoutDTO"
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
            }
        },
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
                            "$ref": "#/definitions/domain.GetUserDTO"
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
                "summary": "Getting a list of user notebooks",
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
        },
        "/notebooks/{notebook_id}/notes": {
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
                    "Note"
                ],
                "summary": "Getting a list of user notes in a notebook",
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
                        "description": "Notes",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/docs.GetAllNotesResponse"
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
                    "Note"
                ],
                "summary": "Creating a note in a notebook",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Notebook ID",
                        "name": "notebook_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Note data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.CreateNoteDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Note ID",
                        "schema": {
                            "$ref": "#/definitions/docs.CreateNoteResponse"
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
        "/notebooks/{notebook_id}/notes/{note_id}": {
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
                    "Note"
                ],
                "summary": "Deleting a note from a notebook",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Notebook ID",
                        "name": "notebook_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Note ID",
                        "name": "note_id",
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
            },
            "patch": {
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
                    "Note"
                ],
                "summary": "Updating a note in a notebook",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Notebook ID",
                        "name": "notebook_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Note ID",
                        "name": "note_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "New note data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.PatchNoteDTO"
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
            }
        },
        "/shared_notes": {
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
                    "Shared note"
                ],
                "summary": "Getting a list of incoming shared notes",
                "responses": {
                    "200": {
                        "description": "Incoming shared notes",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.GetAllIncomingSharedNotesResponse"
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
                    "Shared note"
                ],
                "summary": "Creating a shared note",
                "parameters": [
                    {
                        "description": "Shared note data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.CreateSharedNoteDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Shared note ID",
                        "schema": {
                            "$ref": "#/definitions/docs.CreateSharedNoteResponse"
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
        "/shared_notes/{shared_note_id}": {
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
                    "Shared note"
                ],
                "summary": "Deleting a shared note",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Shared note ID",
                        "name": "shared_note_id",
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
        "docs.CreateNoteResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "docs.CreateNotebookResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "docs.CreateSharedNoteResponse": {
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
                "count": {
                    "type": "integer"
                },
                "notebooks": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "description": {
                                "type": "string"
                            },
                            "id": {
                                "type": "integer"
                            }
                        }
                    }
                }
            }
        },
        "docs.GetAllNotesResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "notes": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "body": {
                                "type": "string"
                            },
                            "created_at": {
                                "type": "string"
                            },
                            "id": {
                                "type": "integer"
                            },
                            "title": {
                                "type": "string"
                            },
                            "updated_at": {
                                "type": "string"
                            }
                        }
                    }
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
                "name": {
                    "type": "string"
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
        "domain.CreateNoteDTO": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "domain.CreateSharedNoteDTO": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "note_id": {
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
        "domain.GetAllIncomingSharedNotesResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "incoming_shared_notes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.IncomingSharedNote"
                    }
                }
            }
        },
        "domain.GetUserDTO": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "domain.IncomingSharedNote": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "owner_login": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "domain.LogoutDTO": {
            "type": "object",
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "domain.PatchNoteDTO": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "title": {
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
