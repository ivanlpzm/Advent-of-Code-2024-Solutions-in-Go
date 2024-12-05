# Use an official Golang image as the base image
FROM golang:1.20-alpine

# Set environment variables
ENV GO111MODULE=auto \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Copy the Go project files to the container
COPY . .

# Build the Go application
RUN go build -o adventCode5 adventCode5.go

# Expose any port your application needs (optional, if applicable)
# EXPOSE 8080 

# Command to run the application
CMD ["./adventCode5"]
