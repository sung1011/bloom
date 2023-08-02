package uuid

import (
	"github.com/google/uuid"
	"github.com/sung1011/bloom/fw"
	"github.com/sung1011/bloom/fw/svc"
)

type FlowerGoogle struct {
	svc.UUID // implements
	sd       fw.Seed
}

func BudGoole(seed fw.Seed, params ...interface{}) (interface{}, error) {
	return &FlowerGoogle{
		sd: seed.(*Seed),
	}, nil
}

func (flw *FlowerGoogle) NewID() string {
	return uuid.New().String()
}
