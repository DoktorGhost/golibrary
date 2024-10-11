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
        "/author/add": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Добавляет нового автора в систему.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Library"
                ],
                "summary": "Добавить автора",
                "parameters": [
                    {
                        "description": "ФИО Автора",
                        "name": "author",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.AuthorRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Автор успешно добавлен, ID: {id}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Ошибка декодирования JSON или чтения тела запроса",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка при добавлении автора",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/authors": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Возвращает список всех авторов с информацией о их книгах.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Library"
                ],
                "summary": "Получить всех авторов",
                "responses": {
                    "200": {
                        "description": "Список авторов",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.Author"
                            }
                        }
                    },
                    "500": {
                        "description": "Ошибка получения авторов",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/books": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Возвращает список всех книг с информацией об авторах.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Library"
                ],
                "summary": "Получить все книги",
                "responses": {
                    "200": {
                        "description": "Список книг",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.Book"
                            }
                        }
                    },
                    "500": {
                        "description": "Ошибка получения книг",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/books/add": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Добавляет новую книгу в систему.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Library"
                ],
                "summary": "Добавить книгу",
                "parameters": [
                    {
                        "description": "Информация о книге: название и ID автора",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.BookRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Книга успешно добавлена, ID: {id}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Ошибка декодирования JSON или чтения тела запроса",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка при добавлении книги",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/debug/pprof/": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Возвращает информацию о профилировании для приложения",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pprof"
                ],
                "summary": "Профилирование приложения",
                "responses": {}
            }
        },
        "/login": {
            "post": {
                "description": "Аутентификация пользователя по имени пользователя и паролю, возвращает JWT-токен.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Логин пользователя",
                "parameters": [
                    {
                        "description": "Данные для входа",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JWT-токен",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Ошибка декодирования данных или ошибка аутентификации",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "405": {
                        "description": "Неправильный метод",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/rental/add/{user_id}/{book_id}": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Позволяет выдать книгу пользователю по его идентификатору и идентификатору книги.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rentals"
                ],
                "summary": "Выдать книгу пользователю",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор пользователя",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Идентификатор книги",
                        "name": "book_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Книга успешно выдана, RentalID: {rentalID}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Неправильный UserID или BookID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка при выдаче книги",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/rental/back/{book_id}": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Позволяет вернуть книгу по её идентификатору.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rentals"
                ],
                "summary": "Вернуть книгу",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор книги",
                        "name": "book_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Книга успешно возвращена",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Неправильный BookID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка при возврате книги",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/rentals": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Возвращает список всех записей аренды для пользователя.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rentals"
                ],
                "summary": "Получить все аренды",
                "responses": {
                    "200": {
                        "description": "Список записей аренды",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.UserWithRentedBooks"
                            }
                        }
                    },
                    "500": {
                        "description": "Ошибка чтения аренды",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/top/{period}/{limit}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Возвращает список топовых авторов за указанный период с заданным лимитом.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Library"
                ],
                "summary": "Получить топ авторов",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Период (количество дней)",
                        "name": "period",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Максимальное количество авторов для возврата",
                        "name": "limit",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Список топовых авторов",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.Author"
                            }
                        }
                    },
                    "400": {
                        "description": "Неправильный период или лимит",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка получения топа авторов",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Возвращает информацию о пользователе по его ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Получить пользователя",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID пользователя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Информация о пользователе",
                        "schema": {
                            "$ref": "#/definitions/dao.UserTable"
                        }
                    },
                    "400": {
                        "description": "Неверный ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка при получении пользователя",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dao.AuthorTable": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dao.BookTable": {
            "type": "object",
            "properties": {
                "authorID": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dao.UserTable": {
            "type": "object",
            "properties": {
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "entities.Author": {
            "type": "object",
            "properties": {
                "books": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dao.BookTable"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entities.AuthorRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "entities.Book": {
            "type": "object",
            "properties": {
                "author": {
                    "$ref": "#/definitions/dao.AuthorTable"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "entities.BookRequest": {
            "type": "object",
            "properties": {
                "author_id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "entities.Login": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "entities.UserWithRentedBooks": {
            "type": "object",
            "properties": {
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "rented_books": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.Book"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Library",
	Description:      "Библиотека",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
