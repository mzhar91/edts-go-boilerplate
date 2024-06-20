package config

import (
	"time"
	
	_api "sg-edts.com/edts-go-boilerplate/pkg/api"
)

func GetTimeoutContext() time.Duration {
	return time.Duration(Cfg.Context.Timeout) * time.Second
}

func ApiSetup() {
	_api.Setup()
}
