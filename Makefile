generate/responses:
	oapi-codegen \
    	-package responses \
    	-generate spec,types \
    	-o gen/responses/responses.gen.go \
    	 specs/responses.yaml
generate/server:
	oapi-codegen \
	-package server \
	-generate gin,types,spec \
	-o gen/server/server.gen.go \
	 specs/server.yaml

generate/creator:
	oapi-codegen \
	-package creator \
	-generate gin,types,spec \
	-o gen/creator/creator.gen.go \
	 specs/creator.yaml
generate/gateway:
	go run main.go
	oapi-codegen \
	-package gen \
 	-generate gin,types,spec \
 	-o gen/gateway/gateway.gen.go \
 	specs/gateway.yaml

generate: gen/types gen/server gen/spec
