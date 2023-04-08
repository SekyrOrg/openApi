folder:
	mkdir -p gen
generate/server: folder
	oapi-codegen \
	-package server \
	-generate gin,types,spec \
	-o gen/server.gen.go \
	 specs/server.yaml

generate/creator: folder
	oapi-codegen \
	-package creator \
	-generate gin,types,spec \
	-o gen/creator.gen.go \
	 specs/creator.yaml
generate/gateway: folder
	go run main.go
	oapi-codegen \
	-package gen \
 	-generate gin,types,spec \
 	-o gen/gateway.gen.go \
 	specs/gateway.yaml

generate: generate/server generate/creator generate/gateway
