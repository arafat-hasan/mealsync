FROM golang:1.24-alpine

WORKDIR /app

# Install build dependencies and Swagger
RUN apk add --no-cache gcc musl-dev && \
    go install github.com/swaggo/swag/cmd/swag@latest

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Generate Swagger documentation
RUN swag init -g cmd/main.go

# Build the application
RUN go build -o main cmd/main.go

# Expose port
EXPOSE 8080

# Command to run the application
CMD ["./main"] 