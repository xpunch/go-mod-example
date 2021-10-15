package grpc

import (
	"fmt"

	example "github.com/xpunch/go-mod-example/v4"
)

func NewETCDRegistry() example.Plugin {
	return &registry{}
}

type registry struct{}

func (s *registry) Load() error {
	fmt.Println("Loading etcd...")
	return nil
}
