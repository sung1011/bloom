package uuid

import (
	"github.com/google/uuid"
	"github.com/sung1011/stickypack/fw/svc"
)

type FlowerGoogle struct {
	svc.UUID
}

func BudGoole(...interface{}) (interface{}, error) {
	return &FlowerGoogle{}, nil
}

func (s *FlowerGoogle) NewID() string {
	return uuid.New().String()
}
