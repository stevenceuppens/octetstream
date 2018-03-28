package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/stevenceuppens/octetstream/openapi/gen/restapi"
	"github.com/stevenceuppens/octetstream/openapi/gen/restapi/operations"

	"net/http"
	_ "net/http/pprof"
)

// Default ENV
var (
	Port = 3000
)

// Capture ENV
func init() {
	if p := os.Getenv("PORT"); p != "" {
		pInt, err := strconv.Atoi(p)
		if err != nil {
			panic("Cannot parse PORT env variable")
		}
		Port = pInt
	}
}

func main() {
	go startProfiling()

	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	openapi := operations.NewDownloadAPI(swaggerSpec)
	server := restapi.NewServer(openapi)
	defer server.Shutdown()
	server.Port = Port

	// Setup Handlers
	openapi.GetFileHandler = operations.GetFileHandlerFunc(download)

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

func download(params operations.GetFileParams) middleware.Responder {
	// take size from query header (10485760 by default = 10MB)
	size := *params.Size

	fmt.Println("download fake file of size: ", size/1024/1024, "MB")

	buffer := make([]byte, size, size)

	// inject random values into buffer
	rand.Read(buffer)

	data := ioutil.NopCloser(bytes.NewBuffer(buffer))

	return operations.NewGetFileOK().WithPayload(data)
}

func startProfiling() {
	http.ListenAndServe(":8080", http.DefaultServeMux)
}
