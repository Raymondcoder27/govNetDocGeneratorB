# FROM golang:1.22.4 AS development

# # Set the Current Working Directory inside the container
# WORKDIR /app

# # Copy go.mod and go.sum files to the workspace
# COPY go.mod go.sum ./

# # Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
# RUN go mod download

# # Copy the source from the current directory to the working directory inside the container
# COPY . .

# # Install reflex for live reloading
# RUN go install github.com/cespare/reflex@latest

# # Expose port 8080 to the outside world
# EXPOSE 8080

# # Command to run the application
# CMD ["reflex", "-g", "*.go", "go", "run", "main.go"]
