.DEFAULT_GOAL := swagger

install_swagger:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: install_swagger
	swagger generate spec -o ./swagger.yml --scan-models