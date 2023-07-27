package uuid

import (
	"github.com/google/uuid"
	"github.com/sung1011/bloom/fw"
	"github.com/sung1011/bloom/fw/svc"
)

type FlowerGoogle struct {
	svc.UUID
}

func BudGoole(seed fw.Seed, params ...interface{}) (interface{}, error) {
	return &FlowerGoogle{}, nil
}

func (s *FlowerGoogle) NewID() string {
	return uuid.New().String()
}
