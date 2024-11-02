// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/appointments": {
            "get": {
                "description": "Get appointments by date range",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Appointment"
                ],
                "summary": "Get appointments by date range",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Start Date with format YYYY-MM-DD",
                        "name": "start_date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "End Date with format YYYY-MM-DD",
                        "name": "end_date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Doctor ID",
                        "name": "doctor_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Patient ID",
                        "name": "patient_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Authorization Token",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Appointment"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    }
                }
            },
            "post": {
                "description": "Create an appointment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Appointment"
                ],
                "summary": "Create an appointment",
                "parameters": [
                    {
                        "description": "Appointment Information",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Appointment"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization Token",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Appointment"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    }
                }
            }
        },
        "/appointments/{id}": {
            "patch": {
                "description": "Cancel an appointment by an id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Appointment"
                ],
                "summary": "Cancel an appointment by an id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Appointment id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization Token",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    }
                }
            }
        },
        "/recover-password": {
            "post": {
                "description": "Recover user password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Recover user password",
                "parameters": [
                    {
                        "description": "Recover Passsword information",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.RecoverPassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    }
                }
            }
        },
        "/reset-password": {
            "post": {
                "description": "Reset user password with verification token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Reset user password with verification token",
                "parameters": [
                    {
                        "description": "Reset User Password",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ResetPassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    }
                }
            }
        },
        "/sign-in": {
            "post": {
                "description": "Authenticate user by DNI and Password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Authenticate user by DNI and Password",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Auth"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    }
                }
            }
        },
        "/users": {
            "post": {
                "description": "Create an regular or admin user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create an regular or admin user",
                "parameters": [
                    {
                        "description": "User Information",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization Token",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    }
                }
            }
        },
        "/users/": {
            "get": {
                "description": "Get appointments by role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get users by role",
                "parameters": [
                    {
                        "enum": [
                            0,
                            1,
                            2
                        ],
                        "type": "integer",
                        "description": "Role ID",
                        "name": "role",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization Token",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.User"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    }
                }
            }
        },
        "/users/assign-role": {
            "patch": {
                "description": "Assign user role by an admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Assign user role by an admin",
                "parameters": [
                    {
                        "description": "Role to update",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UpdateRole"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization Token",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    }
                }
            }
        },
        "/users/load-by-csv": {
            "post": {
                "description": "Load user information by CSV",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Load user information by CSV",
                "parameters": [
                    {
                        "type": "file",
                        "description": "CSV file with user information",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization Token",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    }
                }
            }
        },
        "/users/me": {
            "get": {
                "description": "Get authenticated user information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get authenticated user information",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization Token",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    }
                }
            },
            "patch": {
                "description": "Update authenticated user information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update authenticated user information",
                "parameters": [
                    {
                        "description": "User information to update",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UpdateUser"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization Token",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    }
                }
            }
        },
        "/users/reset-password": {
            "post": {
                "description": "Reset the password of an user by DNI",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Reset the password of an user by DNI",
                "parameters": [
                    {
                        "description": "User password",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UpdatePassword"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization Token",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    }
                }
            }
        },
        "/users/{dni}": {
            "get": {
                "description": "Get user by DNI",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user by DNI",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User DNI",
                        "name": "dni",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization Token",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Appointment": {
            "type": "object",
            "properties": {
                "doctor_id": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "patient_id": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/domain.AppointmentStatus"
                }
            }
        },
        "domain.AppointmentStatus": {
            "type": "integer",
            "enum": [
                0,
                1,
                2
            ],
            "x-enum-varnames": [
                "AppointmentStatusPending",
                "AppointmentStatusCancelled",
                "AppointmentStatusDone"
            ]
        },
        "domain.Auth": {
            "type": "object",
            "properties": {
                "dni": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "domain.RecoverPassword": {
            "type": "object",
            "properties": {
                "dni": {
                    "type": "string"
                }
            }
        },
        "domain.ResetPassword": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "domain.UpdatePassword": {
            "type": "object",
            "properties": {
                "new_password": {
                    "type": "string"
                }
            }
        },
        "domain.UpdateRole": {
            "type": "object",
            "properties": {
                "dni": {
                    "type": "string"
                },
                "new_role": {
                    "$ref": "#/definitions/domain.UserRole"
                }
            }
        },
        "domain.UpdateUser": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "domain.User": {
            "type": "object",
            "required": [
                "dni",
                "email",
                "first_name"
            ],
            "properties": {
                "address": {
                    "type": "string"
                },
                "dni": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean",
                    "default": true
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/domain.UserRole"
                },
                "type_dni": {
                    "$ref": "#/definitions/domain.UserTypeDNI"
                }
            }
        },
        "domain.UserRole": {
            "type": "integer",
            "enum": [
                0,
                1,
                2
            ],
            "x-enum-varnames": [
                "AdminRole",
                "MedicRole",
                "PatientRole"
            ]
        },
        "domain.UserTypeDNI": {
            "type": "integer",
            "enum": [
                0,
                1,
                2
            ],
            "x-enum-comments": {
                "TypeDniTP": "passport"
            },
            "x-enum-varnames": [
                "TypeDniCC",
                "TypeDniTI",
                "TypeDniTP"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{"http", "https"},
	Title:            "Software2Backend",
	Description:      "API para el backend de Software2",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
