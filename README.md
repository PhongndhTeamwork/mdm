# Setup and Running Guide

This guide explains how to set up and run a Go project using the Gin framework, along with configuring Atlas for database schema management.

## Prerequisites
- **Go**: Ensure Go (version 1.24.2 or higher) is installed. Download from [golang.org](https://golang.org/dl/).
- **Git**: Required to clone the repository.
- **IDE**: Optional (e.g., VSCode, GoLand) for a better development experience.
- **Atlas**: Required for database schema migrations. Install the Atlas CLI from [ariga.io](https://atlasgo.io/getting-started).
- **PostgreSQL**: Ensure a PostgreSQL database is running locally or remotely.

## Setup Instructions

1. **Clone the Repository**
   ```bash
   git clone https://github.com/PhongndhTeamwork/mdm.git
   cd mdm
   ```

2. **Install Dependencies**
   Run the following command to install the required Go modules, including Gin and Atlas:
   ```bash
   go mod tidy
   ```

3. **Environment Configuration**
   - If the project uses environment variables, create a `.env` file in the root directory.
   - Example `.env` file:
     ```
     PORT=8080
     DATABASE_URL=postgresql://postgres:Phongsql123@localhost:5432/golang_example?sslmode=disable
     JWT_SECRET=your_jwt_secret
     JWT_EXPIRE=24h
     ```
   - Ensure the project loads `.env` or environment variables using [Viper](https://github.com/spf13/viper).  
     If needed, install it with:
     ```bash
     go get github.com/spf13/viper

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
   - On Window:
   Visit the official installation guide: [atlas](https://atlasgo.io/getting-started#install)

   Verify installation:
   ```bash
   atlas version
   ```

7. **Set Up Atlas**
   - Create an Atlas configuration file (e.g., `atlas.hcl`) in the project root to manage database schemas using GORM models.
   - Example `atlas.hcl`:
     ```hcl
     data "external_schema" "gorm" {
       program = [
         "go",
         "run",
         "-mod=mod",
         "ariga.io/atlas-provider-gorm",
         "load",
         "--path", "./model",  
         "--dialect", "postgres",          
       ]
     }

     env "gorm" {
       src = data.external_schema.gorm.url
       dev = "postgresql://postgres:Phongsql123@localhost:5432/golang_example?sslmode=disable" # Use your local dev DB URL 
       url = "postgresql://postgres:Phongsql123@localhost:5432/golang_example?sslmode=disable" # Use your DB URL  

       migration {
         dir = "file://migrations"    
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

## Running the Project

1. **Start the Server**
   Run the main application file (e.g., `main.go`):
   ```bash
   go run main.go
   ```
   The server will start on the specified port (default: 8080) or as configured in the `.env` file.

2. **Access the Application**
   Open a browser or use a tool like `curl` to access the server:
   - Browser: `http://localhost:8080`
   - Curl: `curl http://localhost:8080`

3. **Hot Reload (Optional)**
   For development with live reloading, install and use `nodemon`:
   ```bash
   npm install -g nodemon
   nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run cmd/main.go
   ```

## Managing Database Migrations with Atlas

1. **Generate Migration Files**
   Use Atlas to compare the GORM models with the database schema and generate migration files:
   ```bash
   atlas migrate diff --env gorm
   ```
   This creates migration files in the `migrations` directory based on the schema defined in the `./model` directory.

2. **Apply Migrations**
   Apply the generated migrations to the database:
   ```bash
   atlas migrate apply --env gorm
   ```
   This updates the database schema to match the GORM models.

## Building the Project
To create a binary for deployment:
```bash
go build -o <binary-name> cmd/main.go
```
Run the binary:
```bash
./<binary-name>
```

## Common Issues
- **Module not found**: Run `go mod tidy` to resolve missing dependencies.
- **Port already in use**: Change the `PORT` in the `.env` file or terminate the process using the port.
- **Database connection error**: Verify database credentials in the `.env` file and ensure the PostgreSQL server is running.
- **Atlas migration errors**: Ensure the `migrations` directory exists and the `atlas.hcl` file is correctly configured.
- **GORM provider errors**: Verify the `./model` directory contains valid GORM models and the Atlas GORM provider is installed.

## Additional Notes
- Ensure your PostgreSQL database is running and accessible.
- Check the project’s `main.go` or documentation for custom routes or configurations.
- For production, consider using a process manager like `systemd` or a container like Docker.
- Regularly back up your database before applying migrations.
