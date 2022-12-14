.DEFAULT_GOAL := swagger_serve

check_install:
	which swagger || go get github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install
	swagger generate spec -o /swagger.yaml --scan-models
	swagger generate spec -o /swagger.json --scan-models

swagger_serve: check_install swagger 
	swagger serve -F=swagger swagger.yaml