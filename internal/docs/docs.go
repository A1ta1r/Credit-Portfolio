// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-11-05 03:03:38.0023774 +0500 +05 m=+0.051862001

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "info": {
        "contact": {},
        "license": {}
    },
    "paths": {
        "/partners": {
            "get": {
                "description": "Метод возвращает список всех имеющихся в системе рекламодателей",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Advertisers"
                ],
                "summary": "Получить список всех рекламодателей",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/responses.AllAdvertisers"
                        }
                    }
                }
            },
            "post": {
                "description": "Метод добавляет в систему нового рекламодателя с заданными параметрами",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Advertisers"
                ],
                "summary": "Добавить нового рекламодателя",
                "parameters": [
                    {
                        "description": "Данные о рекламодателе",
                        "name": "advertiser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/requests.NewAdvertiser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/responses.OneAdvertiser"
                        }
                    }
                }
            }
        },
        "/partners/{id}": {
            "get": {
                "description": "Метод возвращает рекламодателя по его ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Advertisers"
                ],
                "summary": "Получить рекламодателя по ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID рекламодателя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/responses.OneAdvertiser"
                        }
                    },
                    "404": {
                        "description": "{\"message\": \"resource not found\"}"
                    },
                    "422": {}
                }
            },
            "put": {
                "description": "Метод обновляет рекламодателя заданными параметрами по ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Advertisers"
                ],
                "summary": "Обновить рекламодателя",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID рекламодателя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Новые данные о рекламодателе",
                        "name": "advertiser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/requests.UpdateAdvertiser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/responses.OneAdvertiser"
                        }
                    },
                    "404": {},
                    "422": {}
                }
            },
            "delete": {
                "description": "Метод удаляет из системы рекламодателя с заданным ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Advertisers"
                ],
                "summary": "Удалить рекламодателя",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID рекламодателя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {},
                    "422": {}
                }
            }
        },
        "/partners/{id}/promotions": {
            "get": {
                "description": "Метод возвращает данные о рекламных объявлениях рекламодателя с заданным ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Advertisements"
                ],
                "summary": "Получить рекламные объявления рекламодателя",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID рекламодателя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/responses.AllAdvertisements"
                        }
                    },
                    "422": {}
                }
            }
        },
        "/promotions": {
            "get": {
                "description": "Метод возвращает данные о всех рекламных объявлениях",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Advertisements"
                ],
                "summary": "Получить вообще все рекламные объявления",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/responses.AllAdvertisements"
                        }
                    },
                    "422": {}
                }
            },
            "post": {
                "description": "Метод добавляет новое рекламне объявление с заданными параметрами",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Advertisements"
                ],
                "summary": "Добавить рекламное объявление",
                "parameters": [
                    {
                        "description": "Данные о рекламном объявлении",
                        "name": "advertisement",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/requests.NewAdvertisement"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/responses.OneAdvertisement"
                        }
                    }
                }
            }
        },
        "/promotions/{id}": {
            "get": {
                "description": "Метод возвращает данные рекламном объявлении с данным ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Advertisements"
                ],
                "summary": "Получить рекламное объявление по ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID рекламного объявления",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/responses.OneAdvertisement"
                        }
                    },
                    "404": {},
                    "422": {}
                }
            },
            "put": {
                "description": "Метод обновляет рекламне объявление с заданным ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Advertisements"
                ],
                "summary": "Обновить рекламное объявление",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID рекламного объявления",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Новые данные о рекламном объявлении",
                        "name": "advertisement",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/requests.UpdateAdvertisement"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/responses.OneAdvertisement"
                        }
                    },
                    "404": {},
                    "422": {}
                }
            },
            "delete": {
                "description": "Метод удаляет рекламне объявление с заданным ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Advertisements"
                ],
                "summary": "Удалить рекламное объявление",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID рекламного объявления",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {},
                    "422": {}
                }
            }
        }
    },
    "definitions": {
        "entities.Advertisement": {
            "type": "object",
            "properties": {
                "advertiserId": {
                    "type": "integer"
                },
                "banners": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.Banner"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isActive": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "entities.Advertiser": {
            "type": "object",
            "properties": {
                "contactInfo": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string",
                    "example": "2018-10-30T19:43:15.251038Z"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isActive": {
                    "type": "boolean"
                },
                "notes": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string",
                    "example": "2018-10-30T19:43:15.251038Z"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "entities.Banner": {
            "type": "object",
            "properties": {
                "advertisementId": {
                    "type": "integer"
                },
                "bannerPlaceId": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isVisible": {
                    "type": "boolean"
                },
                "pictureUrl": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "uniqueViews": {
                    "type": "integer"
                },
                "uniqueViewsRequired": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "views": {
                    "type": "integer"
                }
            }
        },
        "requests.NewAdvertisement": {
            "type": "object",
            "required": [
                "advertiserId",
                "isActive",
                "title"
            ],
            "properties": {
                "advertiserId": {
                    "type": "integer"
                },
                "isActive": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "requests.NewAdvertiser": {
            "type": "object",
            "required": [
                "email",
                "isActive",
                "password",
                "username"
            ],
            "properties": {
                "contactInfo": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "isActive": {
                    "type": "boolean"
                },
                "notes": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "requests.UpdateAdvertisement": {
            "type": "object",
            "properties": {
                "isActive": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "requests.UpdateAdvertiser": {
            "type": "object",
            "properties": {
                "contactInfo": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "isActive": {
                    "type": "boolean"
                },
                "notes": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "responses.AllAdvertisements": {
            "type": "object",
            "properties": {
                "advertisements": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.Advertisement"
                    }
                },
                "count": {
                    "type": "integer",
                    "example": 0
                },
                "status": {
                    "type": "string",
                    "example": "OK"
                }
            }
        },
        "responses.AllAdvertisers": {
            "type": "object",
            "properties": {
                "advertisers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.Advertiser"
                    }
                },
                "count": {
                    "type": "integer",
                    "example": 0
                },
                "status": {
                    "type": "string",
                    "example": "OK"
                }
            }
        },
        "responses.OneAdvertisement": {
            "type": "object",
            "properties": {
                "Advertisement": {
                    "type": "object",
                    "$ref": "#/definitions/entities.Advertisement"
                }
            }
        },
        "responses.OneAdvertiser": {
            "type": "object",
            "properties": {
                "advertiser": {
                    "type": "object",
                    "$ref": "#/definitions/entities.Advertiser"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
