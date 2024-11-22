API_PROTO_FILES=$(shell find api-proto/src/proto/auth -name *.proto)
.PHONY: generate
generate:
	protoc --proto_path=./api-proto/src \
 	       --go_out=paths=source_relative:./gen/go \
		   --go-grpc_out=paths=source_relative:./gen/go \
	       $(API_PROTO_FILES)