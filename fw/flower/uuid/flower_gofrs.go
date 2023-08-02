package uuid

import (
	"github.com/gofrs/uuid"
	"github.com/sung1011/bloom/fw"
	"github.com/sung1011/bloom/fw/svc"
)

type FlowerGofrs struct {
	svc.UUID // implements
	sd       fw.Seed
}

func BudGofrs(seed fw.Seed, params ...interface{}) (interface{}, error) {
	return &FlowerGofrs{
		sd: seed.(*Seed),
	}, nil
}

func (flw *FlowerGofrs) NewID() string {
	v, _ := uuid.NewV6()
	return v.String()
}
