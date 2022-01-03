# Swagger commands

## Install the go-swagger

> **go get -u github.com/go-swagger/go-swagger/cmd/swagger**

## Generate the swagger doc for the spec, the same can be run from the MakeFile

> **swagger generate spec -o ./swagger.yml --scan-models**

## Generate the swagger client for the automation tests

> **swagger generate client -f ./swagger.yml**
