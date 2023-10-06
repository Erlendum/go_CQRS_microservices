install-goimports:
	go install golang.org/x/tools/cmd/goimports@latest

goimports:
	goimports -w .

gen-server: _gen-server goimports

_gen-server:
	echo ${PWD}
	docker run --rm \
          -v ${PWD}:/go_CQRS_microservices openapitools/openapi-generator-cli generate \
          -i /go_CQRS_microservices/api_gateway_service/api/swagger.yaml \
          -g go-server \
          -o /go_CQRS_microservices/api_gateway_service/internal/server/ \
          --additional-properties=outputAsLibrary=true,sourceFolder=.



.PHONY:

run_api_gateway:
	go run api_gateway_service/cmd/main.go -config=./api_gateway_service/config/config.yaml


proto_reader:
	@echo Generating product reader microservice proto
	cd reader_service/proto/product_reader && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. product_reader.proto

proto_reader_message:
	@echo Generating product reader messages microservice proto
	cd reader_service/proto/product_reader && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. product_reader_messages.proto
