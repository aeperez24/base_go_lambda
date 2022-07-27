package service

import (
	"fmt"
	"os"
)

type PingService interface {
	SendPing() string
}

type pingServiceImpl struct{}

func (pingServiceImpl) SendPing() string {
	aproperty := os.Getenv("MY_PROPERTY")
	return fmt.Sprintf("pong with :%s", aproperty)
}

func NewPingService() PingService {
	return pingServiceImpl{}
}
