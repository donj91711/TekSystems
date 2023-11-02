# Use an official Golang runtime as the base image
FROM golang:1.21

# Set the working directory in the container
WORKDIR /app

# Copy the local code to the container
COPY . .

# Download Go module dependencies
RUN go mod download

# Build the Go application
RUN go build -o myapp

# Expose the port your application listens on (if needed)
EXPOSE 8080

# Command to run the application
CMD ["./myapp"]
