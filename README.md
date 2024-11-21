# API Documentation

This is a Go API built with the Gin framework that provides document generation functionality.

## Base URL

**http://localhost:8080**

## Endpoints

### Health Check

**GET** /ping

Simple health check endpoint to verify the API is running.

### Generate Document

**POST** /generate-doc

Generates documentation based on provided parameters.

#### Request Body

**{**

**    **"**mrLink**"**:** **"string"**,**      **// Merge Request link

**    **"**gitlabToken**"**:** **"string"**,** **// GitLab authentication token

**    **"**model**"**:** **"string"**        **// Model identifier**

**}**

#### Response

**{**

**    **"**doc**"**:** **"string"** **// Generated documentation content**

**}**

#### Status Codes

* 200 OK: Document generated successfully
* 400 Bad Request: Invalid request parameters
* 500 Internal Server Error: Server-side error during processing

## Quick Start

* Make sure you have Go installed on your system
* Install dependencies:

  Bash

  Ask

  Copy

  Run

  **go** **mod** **download**
* Run the server:

  Bash

  Ask

  Copy

  Run

  **go** **run** **main.go**

## Example Usage

Bash

Ask

Copy

Run

**# Health check**

**curl** **http://localhost:8080/ping**

**# Generate documentation**

**curl** **-X** **POST** **http://localhost:8080/generate-doc** **\**

**  **-H** **"Content-Type: application/json"** **\

**  **-d** **'{

**    "mrLink": "https://gitlab.com/your-project/mer**ge_requests/1",

**    "gitlabToken": "your-gitlab-token",**

**    "model": "your-model-identifier"**

**  }'**

Note: Replace the example values in the request body with your actual dat
