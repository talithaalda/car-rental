# Car Rental API

Car Rental API is a Go-based project using the Gin Framework that provides services for managing drivers, incentives, and customers.

## 📌 Prerequisites

Make sure you have installed the following software:

- [Go](https://go.dev/doc/install) (latest version recommended)
- [Gin Framework](https://github.com/gin-gonic/gin)
- [Swag CLI](https://github.com/swaggo/swag) for API documentation
- Database (PostgreSQL/MySQL depending on project configuration)
- Git

## 🚀 Installation Guide

1. **Clone the repository**

   ```sh
   git clone https://github.com/talithaalda/car-rental.git
   cd repository-name
   ```

2. **Initialize and install dependencies:**

   ```sh
    go mod tidy
   ```

3. **Configure environment variables:**

   ```sh
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=your_db_user
    DB_PASSWORD=your_db_password
    DB_NAME=your_db_name
   ```

4. **Run database migrations:**

   ```sh
    go run main.go migrate
   ```

5. **Generate Swagger API documentation:**

   ```sh
    swag init
   ```

6. **Generate Swagger API documentation:**

   ```sh
    go run main.go
   ```

## 📖 API Documentation

Once the server is running, open the Swagger documentation at:

```sh
 http://localhost:3000/swagger/index.html
```
