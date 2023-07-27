package uuid

import (
	"github.com/gofrs/uuid"
	"github.com/sung1011/bloom/fw"
	"github.com/sung1011/bloom/fw/svc"
)

type FlowerGofrs struct {
	svc.UUID
}

func BudGofrs(seed fw.Seed, params ...interface{}) (interface{}, error) {
	return &FlowerGofrs{}, nil
}

func (s *FlowerGofrs) NewID() string {
	v, _ := uuid.NewV6()
	return v.String()
}
