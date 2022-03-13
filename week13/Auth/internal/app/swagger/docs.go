// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package swagger

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/external/login/authenticate": {
            "get": {
                "tags": [
                    "登录管理"
                ],
                "summary": "调用鉴权",
                "parameters": [
                    {
                        "type": "string",
                        "description": "资源特征值",
                        "name": "feature",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "请求类型",
                        "name": "method",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "资源类型（api|...）",
                        "name": "resourceType",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.AuthenticateInfo"
                        }
                    },
                    "401": {
                        "description": "{error:{code:401, message:登录信息无效}}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResult"
                        }
                    },
                    "403": {
                        "description": "{error:{code:600, message:该应用尚未接入}}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResult"
                        }
                    },
                    "500": {
                        "description": "{error:{code:500,message:服务器错误}}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResult"
                        }
                    }
                }
            }
        },
        "/api/pub/login": {
            "post": {
                "tags": [
                    "登录管理"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.LoginParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.LoginTokenInfo"
                        }
                    },
                    "403": {
                        "description": "{error:{code:600, message:该应用尚未接入}}\u003cbr\u003e{error:{code:1000, message:登录失败，账号或密码错误}}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResult"
                        }
                    },
                    "500": {
                        "description": "{error:{code:500, message:服务器错误}}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResult"
                        }
                    }
                }
            }
        },
        "/api/pub/users": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "账号管理"
                ],
                "summary": "创建账号",
                "parameters": [
                    {
                        "description": "创建账号",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.AccountCreateParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.NilResult"
                        }
                    },
                    "400": {
                        "description": "{error:{code:400,message:无效的请求参数。（账号标识不符合格式/账号名不符合格式）}}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResult"
                        }
                    },
                    "403": {
                        "description": "{error:{code:1200,message:账号标识已存在/账号名已存在}}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResult"
                        }
                    },
                    "500": {
                        "description": "{error:{code:500,message:服务器错误}}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResult"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "schema.AccountCreateParam": {
            "type": "object",
            "required": [
                "accountType",
                "password",
                "username"
            ],
            "properties": {
                "accountType": {
                    "description": "账号类型",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "schema.AuthenticateInfo": {
            "type": "object",
            "properties": {
                "grant": {
                    "description": "AccountKey \t\tstring \t` + "`" + `json:\"accountKey\"` + "`" + ` \t// 账号标识\nAccountType   \tstring \t` + "`" + `json:\"accountType\"` + "`" + `   \t// 账号类型",
                    "type": "integer"
                }
            }
        },
        "schema.ErrorItem": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "错误码",
                    "type": "integer"
                },
                "message": {
                    "description": "错误信息",
                    "type": "string"
                }
            }
        },
        "schema.ErrorResult": {
            "type": "object",
            "properties": {
                "error": {
                    "description": "错误项",
                    "$ref": "#/definitions/schema.ErrorItem"
                }
            }
        },
        "schema.LoginParam": {
            "type": "object",
            "required": [
                "appKey",
                "password",
                "username"
            ],
            "properties": {
                "appKey": {
                    "description": "应用标识",
                    "type": "string"
                },
                "password": {
                    "description": "密码(md5加密)",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "schema.LoginTokenInfo": {
            "type": "object",
            "properties": {
                "expiresAt": {
                    "description": "令牌过期时间戳",
                    "type": "integer"
                },
                "token": {
                    "description": "访问令牌",
                    "type": "string"
                }
            }
        },
        "schema.NilResult": {
            "type": "object"
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "RJ-Authorization RJ-AppKey",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.1.0",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{"http", "https"},
	Title:       "Auth",
	Description: "RBAC scaffolding based on GIN + GORM + CASBIN + WIRE.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
