// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

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
        "contact": {
            "name": "gatblau",
            "url": "http://onix.gatblau.org/",
            "email": "onix@gatblau.org"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "Checks that Artie's HTTP server is listening on the required port.\nUse a liveliness probe.\nIt does not guarantee the server is ready to accept calls.",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "General"
                ],
                "summary": "Check that Artie's HTTP API is live",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/artefact/{repository-group}/{repository-name}/id/{artefact-id}": {
            "get": {
                "description": "gets meta data about the artefact identified by its id",
                "consumes": [
                    "text/html",
                    " application/json",
                    " application/yaml",
                    " application/xml",
                    " application/xhtml+xml"
                ],
                "produces": [
                    "application/json",
                    " application/yaml",
                    " application/xml"
                ],
                "tags": [
                    "Artefacts"
                ],
                "summary": "Get information about the specified artefact",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the artefact repository group name",
                        "name": "repository-group",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "the artefact repository name",
                        "name": "repository-name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "updates meta data about the artefact identified by its id",
                "tags": [
                    "Artefacts"
                ],
                "summary": "Update information about the specified artefact",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the artefact repository group name",
                        "name": "repository-group",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "the artefact repository name",
                        "name": "repository-name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "the artefact unique identifier",
                        "name": "artefact-id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "the artefact information to be updated",
                        "name": "artefact-info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/registry.Artefact"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/artefact/{repository-group}/{repository-name}/tag/{artefact-tag}": {
            "post": {
                "description": "uploads the artefact file and its seal to the pre-configured backend (e.g. Nexus, etc)",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Artefacts"
                ],
                "summary": "Push an artefact to the configured backend",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the artefact repository group name",
                        "name": "repository-group",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "the artefact repository name",
                        "name": "repository-name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "the artefact reference name",
                        "name": "tag",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "the artefact metadata in JSON base64 encoded string format",
                        "name": "artefact-meta",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "the artefact file part of the multipart message",
                        "name": "artefact-file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "the seal file part of the multipart message",
                        "name": "artefact-seal",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "423": {
                        "description": "Locked",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/file/{repository-group}/{repository-name}/{filename}": {
            "get": {
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "Files"
                ],
                "summary": "Download a file from the registry",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the artefact repository group name",
                        "name": "repository-group",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "the artefact repository name",
                        "name": "repository-name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "the filename to download",
                        "name": "filename",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/repository/{repository-group}/{repository-name}": {
            "get": {
                "description": "gets meta data about artefacts in the specified repository",
                "consumes": [
                    "text/html",
                    " application/json",
                    " application/yaml",
                    " application/xml",
                    " application/xhtml+xml"
                ],
                "produces": [
                    "application/json",
                    " application/yaml",
                    " application/xml"
                ],
                "tags": [
                    "Repositories"
                ],
                "summary": "Get information about the artefacts in a repository",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the artefact repository group name",
                        "name": "repository-group",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "the artefact repository name",
                        "name": "repository-name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/webhook/{repository-group}/{repository-name}": {
            "get": {
                "description": "get a list of webhook configurations for the specified repository",
                "tags": [
                    "Webhooks"
                ],
                "summary": "get a list of webhooks",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the artefact repository group name",
                        "name": "repository-group",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "the artefact repository name",
                        "name": "repository-name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "create the webhook configuration for a specified repository and url",
                "tags": [
                    "Webhooks"
                ],
                "summary": "creates a webhook configuration",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the artefact repository group name",
                        "name": "repository-group",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "the artefact repository name",
                        "name": "repository-name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "the webhook configuration",
                        "name": "artefact-info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.WebHookConfig"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/webhook/{repository-group}/{repository-name}/{webhook-id}": {
            "delete": {
                "description": "delete the specified webhook configuration",
                "tags": [
                    "Webhooks"
                ],
                "summary": "delete a webhook configuration by Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the artefact repository group name",
                        "name": "repository-group",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "the artefact repository name",
                        "name": "repository-name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "the webhook unique identifier",
                        "name": "webhook-id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "registry.Artefact": {
            "type": "object",
            "properties": {
                "created": {
                    "type": "string"
                },
                "file_ref": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "size": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "server.WebHookConfig": {
            "type": "object",
            "properties": {
                "actions": {
                    "description": "Actions that should trigger the webhook",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "group": {
                    "description": "the repository Group for the webhook",
                    "type": "string"
                },
                "id": {
                    "description": "the unique webhook identifier",
                    "type": "string"
                },
                "name": {
                    "description": "the repository Name for the webhook",
                    "type": "string"
                },
                "pwd": {
                    "description": "the webhook URI password",
                    "type": "string"
                },
                "uname": {
                    "description": "the webhook URI user",
                    "type": "string"
                },
                "uri": {
                    "description": "the webhook URI",
                    "type": "string"
                }
            }
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
	Version:     "0.0.4",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Artie - Artefact Registry",
	Description: "Application Artefacts Registry that supports generic packaging, signing and tagging.\nAllows to manage application artefacts in a similar way to linux container images.",
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
