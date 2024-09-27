# Use Go 1.22 as the base image
FROM golang:1.22

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go mod files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code into the container
COPY main.go ./

# Build the Go application
RUN go build -o main .

# Command to run the executable
CMD ["./main"]