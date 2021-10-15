package example

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func NewService(plugins ...Plugin) Service {
	return &service{plugins: plugins}
}

type service struct {
	plugins []Plugin
}

func (s *service) Run() error {
	exits := make(chan os.Signal, 1)
	signal.Notify(exits, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	for _, p := range s.plugins {
		if err := p.Load(); err != nil {
			return err
		}
	}
	fmt.Println("Running...")
	<-exits
	return nil
}
