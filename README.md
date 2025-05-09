# Setup and Running Guide

This guide explains how to set up and run a Go project using the Gin framework, configure Atlas for database schema management, and integrate Swagger for API documentation.

## Prerequisites
- **Go**: Ensure Go (version 1.24.2 or higher) is installed. Download from [golang.org](https://golang.org/dl/).
- **Git**: Required to clone the repository.
- **IDE**: Optional (e.g., VSCode, GoLand) for a better development experience.
- **Atlas**: Required for database schema migrations. Install the Atlas CLI from [ariga.io](https://atlasgo.io/getting-started).
- **PostgreSQL**: Ensure a PostgreSQL database is running locally or remotely.
- **Swagger**: Required for API documentation. Install the Swag CLI:
  ```bash
  go install github.com/swaggo/swag/cmd/swag@latest
  ```

## Setup Instructions

1. **Clone the Repository**
   ```bash
   git clone https://github.com/PhongndhTeamwork/mdm.git
   cd mdm
   ```

2. **Install Dependencies**
   Run the following command to install the required Go modules, including Gin, Atlas, and Swagger dependencies:
   ```bash
   go mod tidy
   ```
   Install Swagger dependencies:
   ```bash
   go get github.com/swaggo/swag
   go get github.com/swaggo/gin-swagger
   go get github.com/swaggo/files
   ```

3. **Environment Configuration**
   - Create a `.env` file in the root directory for environment variables.
   - Example `.env` file:
     ```
     PORT=8080
     DATABASE_URL=postgresql://postgres:Phongsql123@localhost:5432/golang_example?sslmode=disable
     JWT_SECRET=your_jwt_secret
     JWT_EXPIRE=24h
     ```
   - Load environment variables using [Viper](https://github.com/spf13/viper). Install if needed:
     ```bash
     go get github.com/spf13/viper
     ```

4. **Verify Go Installation**
   Ensure Go is installed and configured:
   ```bash
   go version
   ```

5. **Install Atlas CLI**
   - On macOS/Linux:
     ```bash
     curl -sSf https://atlasgo.sh | sh
     ```
   - On Windows: Follow the official installation guide at [atlasgo.io](https://atlasgo.io/getting-started#install).
   - Verify installation:
     ```bash
     atlas version
     ```

6. **Set Up Atlas**
   - Refer to [atlasgo.io](https://atlasgo.io/guides/orms/gorm) for detailed instructions.
   - Create an `atlas.hcl` file in the project root to manage database schemas using GORM models.
   - Example `atlas.hcl`:
     ```hcl
     data "external_schema" "gorm" {
       program = [
         "go",
         "run",
         "-mod=mod",
         "ariga.io/atlas-provider-gorm",
         "load",
         "--path", "./model",        # Use the relative path to your GORM models directory
         "--dialect", "postgres",    # Use correct dialect
       ]
     }

     env "gorm" {
       src = data.external_schema.gorm.url
       dev = "postgresql://postgres:Phongsql123@localhost:5432/golang_example?sslmode=disable"  # Use your local development DB URL
       url = "postgresql://postgres:Phongsql123@localhost:5432/golang_example?sslmode=disable"  # Target DB for applying migrations
       migration {
         dir = "file://migrations"   # Path to migration folder
       }
       format {
         migrate {
           diff = "{{ sql . \"  \" }}"
         }
       }
     }
     ```
   - Install the Atlas GORM provider:
     ```bash
     go get ariga.io/atlas-provider-gorm
     ```

7. **Set Up Swagger**
   - Go to [swaggo](https://github.com/swaggo/gin-swagger) for complete guide.
   - Generate Swagger documentation by running:
     ```bash
     swag init -g cmd/main.go
     ```
     This generates a `docs` folder with Swagger JSON files.
   - Add Swagger comments (decorators) to your Go code to describe API endpoints. Example:
     ```go
     package main

     import (
         "github.com/gin-gonic/gin"
         "net/http"
     )

     // @Summary Get user by ID
     // @Description Retrieves a user based on the provided ID
     // @Tags users
     // @Accept json
     // @Produce json
     // @Param id path int true "User ID"
     // @Success 200 {object} map[string]interface{} "User data"
     // @Failure 400 {object} map[string]interface{} "Invalid ID"
     // @Failure 404 {object} map[string]interface{} "User not found"
     // @Router /users/{id} [get]
     func GetUser(c *gin.Context) {
         id := c.Param("id")
         // Example response
         c.JSON(http.StatusOK, gin.H{"id": id, "name": "John Doe"})
     }
     ```
   - In `main.go`, integrate Swagger with Gin:
     ```go
     import (
         "github.com/gin-gonic/gin"
         "github.com/swaggo/files"
         "github.com/swaggo/gin-swagger"
         _ "mdm/docs" // Import generated Swagger docs
     )

     func main() {
         r := gin.Default()
         // Add Swagger endpoint
         r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
         // Other routes
         r.GET("/users/:id", GetUser)
         r.Run(":8080")
     }
     ```

## Running the Project

1. **Start the Server**
   Run the main application file:
   ```bash
   go run cmd/main.go
   ```
   The server starts on the specified port (default: 8080) as configured in the `.env` file.

2. **Access the Application**
   - API: Use a browser or `curl` to access endpoints (e.g., `http://localhost:8080/users/1`).
     ```bash
     curl http://localhost:8080/users/1
     ```
   - Swagger UI: Open `http://localhost:8080/swagger/index.html` to view the API documentation.

3. **Hot Reload (Optional)**
   For live reloading during development, use `nodemon`:
   ```bash
   npm install -g nodemon
   nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run cmd/main.go
   ```

## Managing Database Migrations with Atlas

1. **Generate Migration Files**
   Compare GORM models with the database schema to generate migration files:
   ```bash
   atlas migrate diff --env gorm
   ```
   This creates migration files in the `migrations` directory based on the `./model` directory.

2. **Apply Migrations**
   Apply migrations to update the database schema:
   ```bash
   atlas migrate apply --env gorm
   ```

## Building the Project
To create a binary for deployment:
```bash
go build -o <binary_name> cmd/main.go
```
Run the binary:
```bash
./<binary_name>
```

## Common Issues
- **Module not found**: Run `go mod tidy` to resolve missing dependencies.
- **Port already in use**: Change the `PORT` in the `.env` file or terminate the process.
- **Database connection error**: Verify database credentials in `.env` and ensure PostgreSQL is running.
- **Atlas migration errors**: Ensure the `migrations` directory exists and `atlas.hcl` is correctly configured.
- **GORM provider errors**: Verify the `./model` directory contains valid GORM models and the Atlas GORM provider is installed.
- **Swagger not loading**: Ensure `swag init -g cmd/main.go` was run and the `docs` folder exists. Check that the Swagger endpoint is correctly set up in `main.go`.

## Additional Notes
- Ensure PostgreSQL is running and accessible.
- Check `main.go` or project documentation for custom routes or configurations.
- For production, use a process manager like `systemd` or a container like Docker.
- Regularly back up your database before applying migrations.
- Regenerate Swagger docs after updating API endpoints:
  ```bash
  swag init -g cmd/main.go
  ```
