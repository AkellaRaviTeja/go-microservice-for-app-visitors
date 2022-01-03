# App visitor Go Microservice

A microservice in Go lang to perform Create and Read operations on app visitors data on a MongoDB.

Microservice runs on :9191 port.

Read the docker setup doc [here](docs/docker-commands.md)

Read the swagger setup doc [here](docs/swagger-commands.md)

Start the service by running

> **go run service.go**

Once the service starts

Mongo DB can be monitored from the Mongo express by hitting [http://localhost:8081]

Swagger documentation for the APIs can be viewed by hitting [http://localhost:9191/docs]
