# Step 1: Build the Go app
FROM golang:1.23 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the application files
COPY . .

# Build the Go binary
RUN go build -o my-microservice .

# Step 2: Create a minimal image with the Go binary
FROM alpine:latest

# Install necessary certificates
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the binary from the build stage
COPY --from=build /app/my-microservice .

# Expose the port that the app runs on
EXPOSE 8080

# Run the Go app
CMD ["./my-microservice"]
