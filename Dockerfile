# syntax=docker/dockerfile:1

# Latest version of go
FROM golang:latest

# Create a directory to copy the files from host system to the docker image
WORKDIR /app

# Download the necessary modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the go files
COPY *.go ./

# Build the go binary
RUN go build -o /visitor-service
CMD [ "/visitor-service" ]