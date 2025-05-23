# Stage 1: Build the Go binary
FROM golang:1.24-alpine AS build

# Install necessary build tools
RUN apk add --no-cache gcc musl-dev

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files for dependency resolution
COPY go.mod go.sum ./

# Download and cache Go modules
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go binary with executable permissions
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o food-tinder ./cmd

# Set execute permission
RUN chmod +x food-tinder

# Stage 2: Create a minimal image to run the binary
FROM alpine:latest

# Install CA certificates to enable HTTPS
RUN apk add --no-cache ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the built binary from the previous stage
COPY --from=build /app/food-tinder /usr/local/bin/food-tinder

# Copy config
COPY --from=build /app/config /root/config

# Expose the port on which the service will run
EXPOSE 8080

# Specify the entry point command to run the binary
CMD ["/usr/local/bin/food-tinder"]