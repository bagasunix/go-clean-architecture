# Stage 1: Build the Go application
FROM golang:alpine3.20 as builder

# Set environment variables
ENV CGO_ENABLED=0

# Install dependencies
RUN apk add --no-cache make git build-base

# Set the working directory
WORKDIR /usr/app/go

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -ldflags="-s -w" -o go-clean

# Stage 2: Create the final container image
FROM ubuntu:latest

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /usr/app/go/go-clean /app/

# Copy the configuration file from the builder stage
COPY --from=builder /usr/app/go/pkg/config/config.yaml /app/config/

# Install dependencies
RUN apt update && \
    apt upgrade -y && \
    apt install -y tzdata ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Set the timezone
ENV TZ="Asia/Jakarta"

# Expose the port
EXPOSE 8080

# Command to run the application
CMD ["./go-clean"]
