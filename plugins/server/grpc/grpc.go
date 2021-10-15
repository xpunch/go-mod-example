package grpc

import (
	"fmt"

	example "github.com/xpunch/go-mod-example/v3"
)

func NewGrpcServer() example.Plugin {
	return &server{}
}

type server struct{}

func (s *server) Load() error {
	fmt.Println("Loading grpc...")
	return nil
}
