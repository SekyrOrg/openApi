gen/types:
	oapi-codegen -package gateway -generate types -o gen/schema.gen.go sekyr-openapi.yml
gen/server:
	oapi-codegen -package gateway  -generate gin -o gen/server.gen.go sekyr-openapi.yml
gen/spec:
	oapi-codegen -package gateway  -generate spec -o gen/spec.gen.go sekyr-openapi.yml
gen/supa:
	oapi-codegen -package gateway -include-tags creator  -generate spec,gin,types -o gen/supa.gen.go sekyr-openapi.yml
gen: gen/types gen/server gen/spec
