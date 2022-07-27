package service

import (
	"fmt"
	"os"
)

type PingService interface {
	SendPing() string
	SendHello(name string) string
}

type pingServiceImpl struct{}

func (pingServiceImpl) SendPing() string {
	aproperty := os.Getenv("MY_PROPERTY")
	return fmt.Sprintf("pong with :%s", aproperty)
}

func (pingServiceImpl) SendHello(name string) string {
	return fmt.Sprintf("hello :%s", name)
}

func NewPingService() PingService {
	return pingServiceImpl{}
}
