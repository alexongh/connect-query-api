package tests

import (
	"context"

	"connectrpc.com/connect"
	v1 "github.com/alexongh/connect-query-protobuf/gen/go/services/tests/v1"
)

type TestServer struct{}

func (s *TestServer) View(
    ctx context.Context,
    req *connect.Request[v1.ViewRequest],
) (*connect.Response[v1.ViewResponse], error) {

    response := connect.NewResponse(&v1.ViewResponse{
        Name: "Jon Doe",
    })
    
    return response, nil 
}

