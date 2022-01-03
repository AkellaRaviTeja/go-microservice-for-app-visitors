# GO Repo base repo
FROM golang:1.16.12-alpine3.15 as builder

RUN apk add git

# Add Maintainer Info
LABEL maintainer="<>"

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod go.sum ./

# Download all the dependencies
RUN go mod download

COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# GO Repo base repo
FROM alpine:latest

RUN apk --no-cache add ca-certificates curl

RUN mkdir /app

WORKDIR /app/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose ports
EXPOSE 9191
EXPOSE 27017
EXPOSE 8081

# Run Executable
CMD ["./main"]