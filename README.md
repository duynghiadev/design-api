# API Design Guide with Swagger/OpenAPI

## Table of Contents

- [Introduction](#introduction)
- [What is OpenAPI/Swagger?](#what-is-openapiswagger)
- [File Structure](#file-structure)
- [Basic Example](#basic-example)
- [Best Practices](#best-practices)

## Introduction

This guide explains how to design and document RESTful APIs using OpenAPI (formerly known as Swagger). OpenAPI is a specification for machine-readable interface files that helps describe, produce, consume, and visualize RESTful web services.

## What is OpenAPI/Swagger?

- OpenAPI is the specification (current version 3.0)
- Swagger is a set of tools for implementing the OpenAPI specification
- Benefits:
  - Interactive API documentation
  - API code generation
  - Test case generation
  - API validation

## File Structure

OpenAPI documents can be written in YAML or JSON. Here's the basic structure:

```yaml
openapi: 3.0.0
info:
  title: Sample API
  description: Optional description
  version: 1.0.0
servers:
  - url: http://api.example.com/v1
paths:
  /users:
    get:
      summary: Returns a list of users
      responses:
        "200":
          description: A JSON array of user names
          content:
            application/json:
              schema:
                type: array
                items:
                  type: string
```

## Basic Example

Here's a complete example of an API specification for a simple todo application:

```yaml
openapi: 3.0.0
info:
  title: Todo API
  description: A simple Todo API
  version: 1.0.0
  contact:
    name: API Support
    email: support@example.com

servers:
  - url: http://localhost:3000/api/v1
    description: Development server

paths:
  /todos:
    get:
      summary: Get all todos
      responses:
        "200":
          description: List of todos
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: "#/components/schemas/Todo"

    post:
      summary: Create a new todo
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/TodoInput"
      responses:
        "201":
          description: Todo created successfully
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Todo"

components:
  schemas:
    Todo:
      type: object
      properties:
        id:
          type: integer
          format: int64
        title:
          type: string
        completed:
          type: boolean
        createdAt:
          type: string
          format: date-time
      required:
        - title

    TodoInput:
      type: object
      properties:
        title:
          type: string
        completed:
          type: boolean
      required:
        - title
```

## Best Practices

### 1. Use Proper HTTP Methods

- GET: Retrieve resources
- POST: Create new resources
- PUT: Update entire resources
- PATCH: Partial updates
- DELETE: Remove resources

### 2. Schema Definitions

Define reusable components in the `components/schemas` section:

```yaml
components:
  schemas:
    Error:
      type: object
      properties:
        code:
          type: integer
        message:
          type: string
      required:
        - code
        - message
```

### 3. Security Definitions

Include authentication methods:

```yaml
components:
  securitySchemes:
    bearerAuth:
      type: http
      scheme: bearer
      bearerFormat: JWT

security:
  - bearerAuth: []
```

### 4. Response Examples

Include examples in your documentation:

```yaml
responses:
  "200":
    description: Successful response
    content:
      application/json:
        example:
          id: 1
          title: "Complete the project"
          completed: false
```

## Tools and Resources

1. **Swagger UI**: Interactive documentation viewer
2. **Swagger Editor**: Online editor for OpenAPI specs
3. **Swagger Codegen**: Generate server stubs and client SDKs
4. **Postman**: API testing and documentation

## Getting Started

1. Create a new file named `openapi.yaml` or `swagger.json`
2. Define your API specification using the OpenAPI format
3. Use Swagger UI to visualize and test your API
4. Generate server code or documentation as needed

## Validation

Always validate your OpenAPI specification using:

- [Swagger Editor](https://editor.swagger.io/)
- [OpenAPI Validator](https://github.com/swagger-api/validator-badge)

## Additional Tips

- Keep your API documentation up-to-date
- Use clear and consistent naming conventions
- Include authentication details
- Document error responses
- Provide examples for request/response bodies
