# Use the official Golang image as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the rest of the application source code to the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
