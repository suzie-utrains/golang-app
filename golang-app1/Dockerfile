# Stage 1: Build the Go application
FROM golang:1.20-alpine AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o server .

# Stage 2: Create a small image with the built application
FROM alpine:latest

# Install necessary CA certificates
RUN apk --no-cache add ca-certificates

# Create a non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the built application from the previous stage
COPY --from=build /app/server .

# Change ownership to the non-root user
RUN chown -R appuser:appgroup /app

# Set the user to run the application
USER appuser

# Expose port 8080 to the outside world
EXPOSE 8080

# Set environment variable for port
ENV PORT 8080

# Run the executable
ENTRYPOINT ["sh", "-c", "./server $PORT"]

