.DEFAULT_GOAL := swagger_serve

check_install:
	where swagger || go get -u github.com/go-swagger/cmd/swagger

swagger: check_install
	swagger generate spec -o ./swagger.yaml --scan-models

swagger_serve: check_install swagger 
	swagger serve -F=swagger swagger.yaml