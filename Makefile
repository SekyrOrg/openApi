generate/responses:
	oapi-codegen \
    	-package responses \
    	-generate spec,types \
    	-o gen/responses/responses.gen.go \
    	 openapi/responses.yaml
generate/server:
	oapi-codegen \
	-package server \
	-generate gin,types,spec \
	-o gen/server/server.gen.go \
	 openapi/server.yaml

generate/creator:
	oapi-codegen \
	-package creator \
	-generate gin,types,spec \
	-o gen/creator/creator.gen.go \
	 openapi/creator.yaml
generate/gateway:
	go run main.go
	oapi-codegen \
	-package gen \
 	-generate gin,types,spec \
 	-o gen/gateway/gateway.gen.go \
 	openapi/gateway.yaml

generate: gen/types gen/server gen/spec
