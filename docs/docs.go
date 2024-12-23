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
        "/terms": {
            "get": {
                "description": "Возвращает список всех терминов из глоссария.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Глоссарий"
                ],
                "summary": "Получение всех терминов",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.GlossaryTerm"
                            }
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Добавляет новый термин в глоссарий с описанием.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Глоссарий"
                ],
                "summary": "Добавление нового термина",
                "parameters": [
                    {
                        "description": "Термин",
                        "name": "term",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.GlossaryTerm"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.GlossaryTerm"
                        }
                    },
                    "400": {
                        "description": "Неверные данные",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/terms/{term}": {
            "get": {
                "description": "Возвращает информацию о термине по ключевому слову.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Глоссарий"
                ],
                "summary": "Получение информации о термине",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ключевое слово",
                        "name": "term",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.GlossaryTerm"
                        }
                    },
                    "404": {
                        "description": "Термин не найден",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Обновляет информацию о термине в глоссарии.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Глоссарий"
                ],
                "summary": "Обновление термина",
                "parameters": [
                    {
                        "description": "Термин",
                        "name": "term",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.GlossaryTerm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.GlossaryTerm"
                        }
                    },
                    "400": {
                        "description": "Неверные данные",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаляет термин из глоссария по ключевому слову.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Глоссарий"
                ],
                "summary": "Удаление термина",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ключевое слово",
                        "name": "term",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Термин успешно удален"
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.GlossaryTerm": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "term": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
