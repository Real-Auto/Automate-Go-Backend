// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a simple todo list API illustrating go-swagger codegen capabilities.",
    "title": "Automate API",
    "version": "0.0.1"
  },
  "host": "0.0.0.0",
  "basePath": "/",
  "paths": {
    "/Check": {
      "get": {
        "description": "This endpoint returns a confirmation message.",
        "produces": [
          "application/json"
        ],
        "summary": "check message",
        "operationId": "Check",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/GetUser": {
      "get": {
        "description": "This endpoint returns a user object.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Get a user.",
        "operationId": "GetUser",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/changePassword": {
      "post": {
        "description": "This endpoint returns a confirmation message.",
        "produces": [
          "application/json"
        ],
        "summary": "changePassword endpoint",
        "operationId": "changePassword",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/deleteUser": {
      "post": {
        "description": "This endpoint returns a confirmation message.",
        "produces": [
          "application/json"
        ],
        "summary": "delete User endpoint",
        "operationId": "deleteUser",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/login": {
      "post": {
        "description": "This endpoint returns a confirmation message.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Login in endpoint",
        "operationId": "Login",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/signUp": {
      "post": {
        "description": "This endpoint returns a confirmation message.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Sign up endpoint",
        "operationId": "signUp",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/updateUser": {
      "post": {
        "description": "This endpoint returns a confirmation message.",
        "produces": [
          "application/json"
        ],
        "summary": "update User endpoint",
        "operationId": "updateUser",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    }
  },
  "definitions": {
    "Auth0User": {
      "type": "object",
      "properties": {
        "_id": {
          "type": "string",
          "x-go-name": "ID"
        },
        "client_id": {
          "type": "string",
          "x-go-name": "ClientId"
        },
        "connection": {
          "type": "string",
          "x-go-name": "Connection"
        },
        "email": {
          "type": "string",
          "x-go-name": "Email"
        },
        "family_name": {
          "type": "string",
          "x-go-name": "FamilyName"
        },
        "given_name": {
          "type": "string",
          "x-go-name": "GivenName"
        },
        "name": {
          "type": "string",
          "x-go-name": "Name"
        },
        "password": {
          "type": "string",
          "x-go-name": "Password"
        },
        "tenant": {
          "type": "string",
          "x-go-name": "Tenant"
        },
        "user_metadata": {
          "$ref": "#/definitions/UserMetaData"
        }
      },
      "x-go-package": "Automate-Go-Backend/databaseModels"
    },
    "Auth0UserChangePassword": {
      "type": "object",
      "properties": {
        "client_id": {
          "type": "string",
          "x-go-name": "ClientId"
        },
        "connection": {
          "type": "string",
          "x-go-name": "Connection"
        },
        "username": {
          "type": "string",
          "x-go-name": "Email"
        }
      },
      "x-go-package": "Automate-Go-Backend/databaseModels"
    },
    "Auth0UserLogin": {
      "type": "object",
      "properties": {
        "audience": {
          "type": "string",
          "x-go-name": "Audience"
        },
        "client_id": {
          "type": "string",
          "x-go-name": "ClientId"
        },
        "client_secret": {
          "type": "string",
          "x-go-name": "ClientSecret"
        },
        "grant_type": {
          "type": "string",
          "x-go-name": "GrantType"
        },
        "password": {
          "type": "string",
          "x-go-name": "Password"
        },
        "realm": {
          "type": "string",
          "x-go-name": "Realm"
        },
        "scope": {
          "type": "string",
          "x-go-name": "Scope"
        },
        "username": {
          "type": "string",
          "x-go-name": "Email"
        }
      },
      "x-go-package": "Automate-Go-Backend/databaseModels"
    },
    "ChangePasswordPayload": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string",
          "x-go-name": "Email"
        }
      },
      "x-go-package": "Automate-Go-Backend/databaseModels"
    },
    "EditUserPayload": {
      "type": "object",
      "properties": {
        "date_of_birth": {
          "type": "string",
          "x-go-name": "DateOfBirth"
        },
        "first_name": {
          "type": "string",
          "x-go-name": "FirstName"
        },
        "last_name": {
          "type": "string",
          "x-go-name": "LastName"
        },
        "name": {
          "type": "string",
          "x-go-name": "Name"
        },
        "phone": {
          "type": "string",
          "x-go-name": "Phone"
        },
        "photo_file_url": {
          "type": "string",
          "x-go-name": "PhotoFileUrl"
        },
        "services": {
          "type": "string",
          "x-go-name": "Services"
        }
      },
      "x-go-package": "Automate-Go-Backend/databaseModels"
    },
    "GetAuth0UserFieldsPayload": {
      "type": "object",
      "properties": {
        "access_token": {
          "type": "string",
          "x-go-name": "AccessToken"
        }
      },
      "x-go-package": "Automate-Go-Backend/databaseModels"
    },
    "LoginPayload": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string",
          "x-go-name": "Email"
        },
        "password": {
          "type": "string",
          "x-go-name": "Password"
        }
      },
      "x-go-package": "Automate-Go-Backend/databaseModels"
    },
    "SignUpPayload": {
      "type": "object",
      "properties": {
        "date_of_birth": {
          "type": "string",
          "x-go-name": "DateOfBirth"
        },
        "email": {
          "type": "string",
          "x-go-name": "Email"
        },
        "first_name": {
          "type": "string",
          "x-go-name": "FirstName"
        },
        "last_name": {
          "type": "string",
          "x-go-name": "LastName"
        },
        "name": {
          "type": "string",
          "x-go-name": "Name"
        },
        "password": {
          "type": "string",
          "x-go-name": "Password"
        },
        "phone": {
          "type": "string",
          "x-go-name": "Phone"
        },
        "photo_file_url": {
          "type": "string",
          "x-go-name": "PhotoFileUrl"
        },
        "services": {
          "type": "string",
          "x-go-name": "Services"
        }
      },
      "x-go-package": "Automate-Go-Backend/databaseModels"
    },
    "UserMetaData": {
      "type": "object",
      "properties": {
        "date_of_birth": {
          "type": "string",
          "x-go-name": "DateOfBirth"
        },
        "phone": {
          "type": "string",
          "x-go-name": "Phone"
        },
        "photo_file_url": {
          "type": "string",
          "x-go-name": "PhotoFileUrl"
        },
        "services": {
          "type": "string",
          "x-go-name": "Services"
        }
      },
      "x-go-package": "Automate-Go-Backend/databaseModels"
    }
  },
  "securityDefinitions": {
    "oauth2": {
      "type": "oauth2",
      "flow": "accessCode",
      "authorizationUrl": "/oauth2/auth",
      "tokenUrl": "/oauth2/token",
      "scopes": {
        "bar": "foo"
      }
    }
  },
  "x-meta-array": [
    "value1",
    "value2"
  ],
  "x-meta-array-obj": [
    {
      "name": "obj",
      "value": "field"
    }
  ],
  "x-meta-value": "value"
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a simple todo list API illustrating go-swagger codegen capabilities.",
    "title": "Automate API",
    "version": "0.0.1"
  },
  "host": "0.0.0.0",
  "basePath": "/",
  "paths": {
    "/Check": {
      "get": {
        "description": "This endpoint returns a confirmation message.",
        "produces": [
          "application/json"
        ],
        "summary": "check message",
        "operationId": "Check",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/GetUser": {
      "get": {
        "description": "This endpoint returns a user object.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Get a user.",
        "operationId": "GetUser",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/changePassword": {
      "post": {
        "description": "This endpoint returns a confirmation message.",
        "produces": [
          "application/json"
        ],
        "summary": "changePassword endpoint",
        "operationId": "changePassword",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/deleteUser": {
      "post": {
        "description": "This endpoint returns a confirmation message.",
        "produces": [
          "application/json"
        ],
        "summary": "delete User endpoint",
        "operationId": "deleteUser",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/login": {
      "post": {
        "description": "This endpoint returns a confirmation message.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Login in endpoint",
        "operationId": "Login",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/signUp": {
      "post": {
        "description": "This endpoint returns a confirmation message.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Sign up endpoint",
        "operationId": "signUp",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/updateUser": {
      "post": {
        "description": "This endpoint returns a confirmation message.",
        "produces": [
          "application/json"
        ],
        "summary": "update User endpoint",
        "operationId": "updateUser",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    }
  },
  "definitions": {
    "Auth0User": {
      "type": "object",
      "properties": {
        "_id": {
          "type": "string",
          "x-go-name": "ID"
        },
        "client_id": {
          "type": "string",
          "x-go-name": "ClientId"
        },
        "connection": {
          "type": "string",
          "x-go-name": "Connection"
        },
        "email": {
          "type": "string",
          "x-go-name": "Email"
        },
        "family_name": {
          "type": "string",
          "x-go-name": "FamilyName"
        },
        "given_name": {
          "type": "string",
          "x-go-name": "GivenName"
        },
        "name": {
          "type": "string",
          "x-go-name": "Name"
        },
        "password": {
          "type": "string",
          "x-go-name": "Password"
        },
        "tenant": {
          "type": "string",
          "x-go-name": "Tenant"
        },
        "user_metadata": {
          "$ref": "#/definitions/UserMetaData"
        }
      },
      "x-go-package": "Automate-Go-Backend/databaseModels"
    },
    "Auth0UserChangePassword": {
      "type": "object",
      "properties": {
        "client_id": {
          "type": "string",
          "x-go-name": "ClientId"
        },
        "connection": {
          "type": "string",
          "x-go-name": "Connection"
        },
        "username": {
          "type": "string",
          "x-go-name": "Email"
        }
      },
      "x-go-package": "Automate-Go-Backend/databaseModels"
    },
    "Auth0UserLogin": {
      "type": "object",
      "properties": {
        "audience": {
          "type": "string",
          "x-go-name": "Audience"
        },
        "client_id": {
          "type": "string",
          "x-go-name": "ClientId"
        },
        "client_secret": {
          "type": "string",
          "x-go-name": "ClientSecret"
        },
        "grant_type": {
          "type": "string",
          "x-go-name": "GrantType"
        },
        "password": {
          "type": "string",
          "x-go-name": "Password"
        },
        "realm": {
          "type": "string",
          "x-go-name": "Realm"
        },
        "scope": {
          "type": "string",
          "x-go-name": "Scope"
        },
        "username": {
          "type": "string",
          "x-go-name": "Email"
        }
      },
      "x-go-package": "Automate-Go-Backend/databaseModels"
    },
    "ChangePasswordPayload": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string",
          "x-go-name": "Email"
        }
      },
      "x-go-package": "Automate-Go-Backend/databaseModels"
    },
    "EditUserPayload": {
      "type": "object",
      "properties": {
        "date_of_birth": {
          "type": "string",
          "x-go-name": "DateOfBirth"
        },
        "first_name": {
          "type": "string",
          "x-go-name": "FirstName"
        },
        "last_name": {
          "type": "string",
          "x-go-name": "LastName"
        },
        "name": {
          "type": "string",
          "x-go-name": "Name"
        },
        "phone": {
          "type": "string",
          "x-go-name": "Phone"
        },
        "photo_file_url": {
          "type": "string",
          "x-go-name": "PhotoFileUrl"
        },
        "services": {
          "type": "string",
          "x-go-name": "Services"
        }
      },
      "x-go-package": "Automate-Go-Backend/databaseModels"
    },
    "GetAuth0UserFieldsPayload": {
      "type": "object",
      "properties": {
        "access_token": {
          "type": "string",
          "x-go-name": "AccessToken"
        }
      },
      "x-go-package": "Automate-Go-Backend/databaseModels"
    },
    "LoginPayload": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string",
          "x-go-name": "Email"
        },
        "password": {
          "type": "string",
          "x-go-name": "Password"
        }
      },
      "x-go-package": "Automate-Go-Backend/databaseModels"
    },
    "SignUpPayload": {
      "type": "object",
      "properties": {
        "date_of_birth": {
          "type": "string",
          "x-go-name": "DateOfBirth"
        },
        "email": {
          "type": "string",
          "x-go-name": "Email"
        },
        "first_name": {
          "type": "string",
          "x-go-name": "FirstName"
        },
        "last_name": {
          "type": "string",
          "x-go-name": "LastName"
        },
        "name": {
          "type": "string",
          "x-go-name": "Name"
        },
        "password": {
          "type": "string",
          "x-go-name": "Password"
        },
        "phone": {
          "type": "string",
          "x-go-name": "Phone"
        },
        "photo_file_url": {
          "type": "string",
          "x-go-name": "PhotoFileUrl"
        },
        "services": {
          "type": "string",
          "x-go-name": "Services"
        }
      },
      "x-go-package": "Automate-Go-Backend/databaseModels"
    },
    "UserMetaData": {
      "type": "object",
      "properties": {
        "date_of_birth": {
          "type": "string",
          "x-go-name": "DateOfBirth"
        },
        "phone": {
          "type": "string",
          "x-go-name": "Phone"
        },
        "photo_file_url": {
          "type": "string",
          "x-go-name": "PhotoFileUrl"
        },
        "services": {
          "type": "string",
          "x-go-name": "Services"
        }
      },
      "x-go-package": "Automate-Go-Backend/databaseModels"
    }
  },
  "securityDefinitions": {
    "oauth2": {
      "type": "oauth2",
      "flow": "accessCode",
      "authorizationUrl": "/oauth2/auth",
      "tokenUrl": "/oauth2/token",
      "scopes": {
        "bar": "foo"
      }
    }
  },
  "x-meta-array": [
    "value1",
    "value2"
  ],
  "x-meta-array-obj": [
    {
      "name": "obj",
      "value": "field"
    }
  ],
  "x-meta-value": "value"
}`))
}