# Use the official Golang image as a base
FROM golang:1.24.1

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# # Download Go modules
RUN go mod download

# # Copy the source code into the container
COPY . .


WORKDIR /app/cmd/server
# # # Build the Go application
RUN go build -o main .

# # Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]