# Stage 1: Build the Go application
FROM golang:1.23.0-alpine3.20 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o incident-status-exporter

# Stage 2: Create a minimal runtime image
FROM alpine:3.20

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/incident-status-exporter .

# Expose the necessary port (replace with your app's port)
EXPOSE 8001

# Run the exporter
CMD ["./incident-status-exporter"]
