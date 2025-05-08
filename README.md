# Setup and Running Guide

This guide explains how to set up and run a Go project using the Gin framework.

## Prerequisites
- **Go**: Ensure Go (version 1.16 or higher) is installed. Download from [golang.org](https://golang.org/dl/).
- **Git**: Required to clone the repository.
- **IDE**: Optional (e.g., VSCode, GoLand) for a better development experience.

## Setup Instructions

1. **Clone the Repository**
   ```bash
   git clone https://github.com/PhongndhTeamwork/mdm.git
   cd mdm
   ```

2. **Install Dependencies**
   Run the following command to install the required Go modules, including Gin:
   ```bash
   go mod tidy
   ```

3. **Environment Configuration**
   - If the project uses environment variables (e.g., `.env` file), create a `.env` file in the root directory.
   - Example `.env` file:
     ```
     PORT=8080
     DATABASE_URL=
     JWT_SECRET=
     JWT_EXPIRE=
     ```
   - Ensure you have the `github.com/joho/godotenv` package if the project loads `.env` files.

4. **Verify Go Installation**
   Ensure Go is installed and configured:
   ```bash
   go version
   ```

## Running the Project

1. **Start the Server**
   Run the main application file (e.g., `main.go`):
   ```bash
   go run main.go
   ```
   The server will start on the specified port (default: `8080`) or as configured in the `.env` file.

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
- **Database connection error**: Verify database credentials in the `.env` file.

## Additional Notes
- Ensure your database (if used) is running and accessible.
- Check the project’s `main.go` or documentation for custom routes or configurations.
- For production, consider using a process manager like `systemd` or a container like Docker.
