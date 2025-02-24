# Use official Golang image
FROM golang:1-alpine AS builder

# Set environment variables
WORKDIR /app

# Copy files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN go build -o main .

# Use a minimal image for running the app
FROM gcr.io/distroless/base-debian11

WORKDIR /root/
COPY --from=builder /app/main .

# Expose the port
EXPOSE 8880

# Start the server
CMD ["/root/main"]