package main

import (
	"connect-query-api/src/handlers/tests"
	"net/http"

	connectcors "connectrpc.com/cors"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/alexongh/connect-query-protobuf/gen/go/services/tests/v1/v1connect"
)

func main() {

    mux := http.NewServeMux()

    mux.Handle(v1connect.NewTestServiceHandler(&tests.TestServer{}))

    c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:3000"},
        AllowCredentials: true,
        AllowedMethods: connectcors.AllowedMethods(),
        AllowedHeaders: connectcors.AllowedHeaders(),
        ExposedHeaders: connectcors.ExposedHeaders(),
    })

    // Insert the middleware
    handler2 := c.Handler(mux)

    http.ListenAndServe(
        "localhost:3001",
        // Use h2c so we can serve HTTP/2 without TLS.
        h2c.NewHandler(handler2, &http2.Server{}),
    )
}

