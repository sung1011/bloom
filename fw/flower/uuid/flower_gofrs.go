package uuid

import (
	"github.com/gofrs/uuid"
	"github.com/sung1011/bloom/fw"
)

type FlowerGofrs struct {
	sd fw.Seed
}

func BudGofrs(seed fw.Seed) (interface{}, error) {
	return &FlowerGofrs{
		sd: seed.(*Seed),
	}, nil
}

func (flw *FlowerGofrs) NewID() string {
	v, _ := uuid.NewV6()
	return v.String()
}
