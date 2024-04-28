package uuid

import (
	"github.com/google/uuid"
	"github.com/sung1011/bloom/fw"
)

type FlowerGoogle struct {
	sd fw.Seed
}

func BudGoole(seed fw.Seed) (interface{}, error) {
	return &FlowerGoogle{
		sd: seed.(*Seed),
	}, nil
}

func (flw *FlowerGoogle) NewID() string {
	return uuid.New().String()
}
