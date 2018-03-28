OPENAPI_GEN=./openapi/gen
OPENAPI_SPEC=./openapi/openapi.yaml

.PHONY: all clean build openapi-server-gen openapi-server-clean

all: build

clean:
	rm -Rf ./bin

build:
	go build -o ./bin/octetstream ./cmd/octetstream.go

openapi-server-gen:
	mkdir -p ${OPENAPI_GEN}
	swagger generate server -t ${OPENAPI_GEN} -f ${OPENAPI_SPEC} --exclude-main
	
openapi-server-clean:
	rm -Rf ${OPENAPI_GEN}