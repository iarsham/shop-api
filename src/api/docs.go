// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "email": "arshamdev2001@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user/register": {
            "post": {
                "description": "Create user with firstname / lastname / phone",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Register By Phone",
                "parameters": [
                    {
                        "description": "register body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/responses.RegisterOKResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/responses.RegisterConflictResponse"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/responses.InterServerErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/send-otp": {
            "post": {
                "description": "This endpoint receives the user's phone in request body and generates an otp. it then sends the otp to the user's phone via sms.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Send OTP",
                "parameters": [
                    {
                        "description": "send otp body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SendOTPRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/responses.SendOtpOkResponse"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/responses.UserNotFoundResponse"
                        }
                    }
                }
            }
        },
        "/user/verify-otp": {
            "post": {
                "description": "this endpoint receives the user's phone and otp code in request body.if code match, the verification is successfully.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Verify OTP",
                "parameters": [
                    {
                        "description": "verify otp body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.VerifyOTPRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/responses.VerifyOTPResponse"
                        }
                    },
                    "401": {
                        "description": "incorrect",
                        "schema": {
                            "$ref": "#/definitions/responses.OtpIncorrectResponse"
                        }
                    },
                    "410": {
                        "description": "Expired",
                        "schema": {
                            "$ref": "#/definitions/responses.OtpExpiredResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.RegisterRequest": {
            "type": "object",
            "required": [
                "first_name",
                "last_name",
                "phone"
            ],
            "properties": {
                "first_name": {
                    "type": "string",
                    "maxLength": 75,
                    "minLength": 1,
                    "example": "James"
                },
                "last_name": {
                    "type": "string",
                    "maxLength": 75,
                    "minLength": 1,
                    "example": "Rodriguez"
                },
                "phone": {
                    "type": "string",
                    "maxLength": 13,
                    "minLength": 11,
                    "example": "+989021112299"
                }
            }
        },
        "dto.SendOTPRequest": {
            "type": "object",
            "required": [
                "phone"
            ],
            "properties": {
                "phone": {
                    "type": "string",
                    "maxLength": 13,
                    "minLength": 11,
                    "example": "+989021112299"
                }
            }
        },
        "dto.VerifyOTPRequest": {
            "type": "object",
            "required": [
                "code",
                "phone"
            ],
            "properties": {
                "code": {
                    "type": "string",
                    "maxLength": 6,
                    "minLength": 6,
                    "example": "241960"
                },
                "phone": {
                    "type": "string",
                    "maxLength": 13,
                    "minLength": 11,
                    "example": "+989021112299"
                }
            }
        },
        "responses.InterServerErrorResponse": {
            "type": "object",
            "properties": {
                "response": {
                    "type": "string",
                    "example": "Internal server error"
                }
            }
        },
        "responses.OtpExpiredResponse": {
            "type": "object",
            "properties": {
                "response": {
                    "type": "string",
                    "example": "otp expired"
                }
            }
        },
        "responses.OtpIncorrectResponse": {
            "type": "object",
            "properties": {
                "response": {
                    "type": "string",
                    "example": "code is incorrect"
                }
            }
        },
        "responses.RegisterConflictResponse": {
            "type": "object",
            "properties": {
                "response": {
                    "type": "string",
                    "example": "user with this phone already exists"
                }
            }
        },
        "responses.RegisterOKResponse": {
            "type": "object",
            "properties": {
                "response": {
                    "type": "string",
                    "example": "user created"
                }
            }
        },
        "responses.SendOtpOkResponse": {
            "type": "object",
            "properties": {
                "response": {
                    "type": "string",
                    "example": "otp was sent"
                }
            }
        },
        "responses.UserNotFoundResponse": {
            "type": "object",
            "properties": {
                "response": {
                    "type": "string",
                    "example": "user not found"
                }
            }
        },
        "responses.VerifyOTPResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTYzMzg4MTcsInBob25lIjoiKzk4OTAyMTMxMjIyNCIsInVzZXJfaWQiOiI1In0.DAxOeyiWpPZWXyVnnyLajMQ9SGsBKw65qOAurhjlFy0"
                },
                "refresh_token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTY5NDE4MTcsInBob25lIjoiKzk4OTAyMTMxMjIyNCIsInVzZXJfaWQiOiI1In0.hzmZdfltaMDWaiTwO8IG1uPEyXOsu3JBs6giU2BDeMI"
                }
            }
        }
    },
    "securityDefinitions": {
        "Authorization": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:8000",
	BasePath:         "/api/v1/",
	Schemes:          []string{"http", "https"},
	Title:            "Shop API Document",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
