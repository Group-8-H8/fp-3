// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/categories": {
            "get": {
                "description": "Get all categories",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Get All Categories",
                "operationId": "get-all-category",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.NewGetCategoriesResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Create Category",
                "operationId": "create-new-category",
                "parameters": [
                    {
                        "description": "request body json",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.NewCreateCategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.NewCreateCategoryResponse"
                        }
                    }
                }
            }
        },
        "/categories/{categoryId}": {
            "get": {
                "description": "Get category by categories ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Get Category",
                "operationId": "get-category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of the category",
                        "name": "categoryId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.NewGetCategoriesResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Category",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Delete Category",
                "operationId": "delete-category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of the category",
                        "name": "categoryId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.NewDeleteCategoryResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update category name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Update Category",
                "operationId": "update-category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of the category",
                        "name": "categoryId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request body json",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.NewUpdateCategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.NewUpdateCategoryResponse"
                        }
                    }
                }
            }
        },
        "/tasks": {
            "get": {
                "description": "Get all tasks",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "Get All Tasks",
                "operationId": "get-all-task",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.NewGetTaskResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "Create Task",
                "operationId": "create-new-task",
                "parameters": [
                    {
                        "description": "request body json",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.NewCreateTaskRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.NewCreateTaskResponse"
                        }
                    }
                }
            }
        },
        "/tasks/update-category/{taskId}": {
            "patch": {
                "description": "Update task's category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "Update Task's Category",
                "operationId": "update-tasks-category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of the task",
                        "name": "taskId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request body json",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.NewUpdateTasksCategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.NewUpdateTaskResponse"
                        }
                    }
                }
            }
        },
        "/tasks/update-status/{taskId}": {
            "patch": {
                "description": "Update task's status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "Update Task's Status",
                "operationId": "update-tasks-status",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of the task",
                        "name": "taskId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request body json",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.NewUpdateTasksStatusRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.NewUpdateTaskResponse"
                        }
                    }
                }
            }
        },
        "/tasks/{taskId}": {
            "get": {
                "description": "Get task by task's ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "Get Task By ID",
                "operationId": "get-task-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of the task",
                        "name": "taskId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.NewGetTaskResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update task's detail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "Update Task",
                "operationId": "update-task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of the task",
                        "name": "taskId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request body json",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.NewUpdateTaskRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.NewUpdateTaskResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete task by task's ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "Delete Task",
                "operationId": "delete-task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of the task",
                        "name": "taskId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.NewDeleteTaskResponse"
                        }
                    }
                }
            }
        },
        "/users/delete-account": {
            "delete": {
                "description": "Delete an account",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete Account",
                "operationId": "delete-account",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.NewDeleteAccountResponse"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "Login registered account to get the token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login Registered Account",
                "operationId": "login-account",
                "parameters": [
                    {
                        "description": "request body json",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.NewLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.NewLoginResponse"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "Create a new account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create Account",
                "operationId": "register-new-account",
                "parameters": [
                    {
                        "description": "request body json",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.NewRegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.NewRegisterResponse"
                        }
                    }
                }
            }
        },
        "/users/update-account": {
            "put": {
                "description": "Update account's fullname and email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update Account",
                "operationId": "update-account",
                "parameters": [
                    {
                        "description": "request body json",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.NewUpdateAccountRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.NewUpdateAccountResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.NewCreateCategoryRequest": {
            "type": "object",
            "required": [
                "type"
            ],
            "properties": {
                "type": {
                    "type": "string"
                }
            }
        },
        "dto.NewCreateCategoryResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "dto.NewCreateTaskRequest": {
            "type": "object",
            "required": [
                "category_id",
                "description",
                "title"
            ],
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dto.NewCreateTaskResponse": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "dto.NewDeleteAccountResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.NewDeleteCategoryResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.NewDeleteTaskResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.NewGetCategoriesResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "tasks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.NewGetTaskOnCategoriesEndpoint"
                    }
                },
                "type": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "dto.NewGetTaskOnCategoriesEndpoint": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
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
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "dto.NewGetTaskResponse": {
            "type": "object",
            "properties": {
                "User": {
                    "$ref": "#/definitions/dto.NewUserOnTaskResponse"
                },
                "category_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "dto.NewLoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "dto.NewLoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "dto.NewRegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "full_name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "dto.NewRegisterResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "dto.NewUpdateAccountRequest": {
            "type": "object",
            "required": [
                "email",
                "full_name"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                }
            }
        },
        "dto.NewUpdateAccountResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "dto.NewUpdateCategoryRequest": {
            "type": "object",
            "required": [
                "type"
            ],
            "properties": {
                "type": {
                    "type": "string"
                }
            }
        },
        "dto.NewUpdateCategoryResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "dto.NewUpdateTaskRequest": {
            "type": "object",
            "required": [
                "description",
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dto.NewUpdateTaskResponse": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "dto.NewUpdateTasksCategoryRequest": {
            "type": "object",
            "required": [
                "category_id"
            ],
            "properties": {
                "category_id": {
                    "type": "integer"
                }
            }
        },
        "dto.NewUpdateTasksStatusRequest": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "boolean"
                }
            }
        },
        "dto.NewUserOnTaskResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
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
	Title:            "Final Project 3 - Group 8 Hacktiv8",
	Description:      "This is a documentation for kanban board API from final project 3 - Group 8 Hacktiv8",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
