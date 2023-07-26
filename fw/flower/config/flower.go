package config

import (
	"github.com/sung1011/stickypack/fw/svc"
)

type Flower struct {
	svc.Config
}

func Bud(params ...interface{}) (interface{}, error) {
	// @@todo
	return &Flower{}, nil
}
