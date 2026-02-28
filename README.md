# Gin Product Service

A backend product management service built with Go (Gin) implementing a clean, pragmatic architecture.  
This service provides CRUD operations with pagination support and is designed as a solid foundation before scaling into authentication, payment, or notification services.

The project intentionally avoids premature over-engineering while keeping clear separation of concerns so the codebase remains readable, testable, and easy to extend.

---

## 🧩 Architecture Mapping

| Layer            | Folder / File Path             | Responsibility                   |
|------------------|--------------------------------|----------------------------------|
| Main             | cmd/main.go                    | App bootstrap, dependency wiring |
| Router           | internal/delivery/http/router  | HTTP route definitions           |
| Handler          | internal/delivery/http/handler | HTTP request/response handling   |
| Usecase          | internal/usecase               | Business logic                   |
| Domain Interface | internal/domain                | Business contracts (interfaces)  |
| Repository       | internal/repository/impl       | Database access                  |
| Database         | config/database.go             | Database initialization          |

---

## Overview

![GitHub last commit](https://img.shields.io/github/last-commit/ArifRosandika/gin_product_service?color=blue)
![GitHub repo size](https://img.shields.io/github/repo-size/ArifRosandika/gin_product_service)
![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)
![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-00ADD8?logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?logo=postgresql&logoColor=white)

---

## 📘 Table of Contents

- [Architecture Mapping](#architecture-mapping)
- [Overview](#overview)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Product Flow](#product-flow)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Run Locally](#run-locally)
- [Environment Variables](#environment-variables)
- [API Endpoints](#api-endpoints)
- [API Testing](#api-testing)
- [License](#license)

---

## ✨ Features

- Create new product
- Get product list with pagination
- Get product detail by ID
- Update product data
- Delete product
- PostgreSQL persistence with GORM
- Clean Architecture (Handler → Usecase → Repository)
- Request validation layer
- Environment-based configuration
- Dockerized for local development

---

## 🧱 Tech Stack

### Backend
- Go
- Gin (HTTP framework)
- GORM (PostgreSQL ORM)
- Viper

### Infrastructure
- PostgreSQL

---

## 📦 Product Flow

### Create Product
1. Client sends product payload
2. Request is validated
3. Usecase applies business rules
4. Product is stored in database

### Get Products
1. Client requests product list
2. Repository retrieves data from database
3. Response is returned to client

### Update Product
1. Client sends updated product data
2. Product existence is validated
3. Data is updated in database

### Delete Product
1. Client sends delete request
2. Product is removed from database

---

## 📁 Project Structure

```text
.
├── cmd/
│   └── main.go                 # Application entry point
├── config/
│   └── database.go             # Database initialization
├── internal/
│   ├── delivery/
│   │   └── http/
│   │       ├── dto/             # HTTP request DTOs
│   │       ├── handler/         # HTTP handlers
│   │       └── router/          # Route definitions
│   ├── domain/                  # Domain interfaces
│   ├── repository/              # Repository layer
│   ├── usecase/                 # Business logic
│   └── helper/                  # Shared helpers
├── .env.example                 # Environment variables template
├── go.mod
├── go.sum
├── README.md
└── test.rest                    # API testing with REST client

---

🚀 Getting Started
Prerequisites

Go 1.22+

PostgreSQL

Run Locally
go mod tidy
go run cmd/main.go

---

🌱 Environment Variables

Copy the example file and adjust values as needed:

cp .env.example .env

---

🔀 API Endpoints
List Products (Pagination)

GET /products?limit=10&page=1

Query Parameters
Name	Default	Max
limit	10	50
page	1	-
Response Example
{
  "status": "success",
  "message": "products fetched",
  "data": {
    "items": [
      {
        "id": 1,
        "name": "Product 1",
        "description": "Nice product",
        "price": 100,
        "image": "http://image.png"
      }
    ],
    "meta": {
      "page": 1,
      "limit": 10,
      "total": 5
    }
  }
}

---

🧪 API Testing

A test.rest file is included for testing:

Create product

Get product list

Get product detail

Update product

Delete product

Compatible with VS Code REST Client extension.

---

📜 License

This project is licensed under the MIT License.