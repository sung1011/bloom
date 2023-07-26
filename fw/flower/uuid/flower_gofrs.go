package uuid

import (
	"github.com/gofrs/uuid"
	"github.com/sung1011/stickypack/fw/svc"
)

type FlowerGofrs struct {
	svc.UUID
}

func BudGofrs(...interface{}) (interface{}, error) {
	return &FlowerGofrs{}, nil
}

func (s *FlowerGofrs) NewID() string {
	v, _ := uuid.NewV6()
	return v.String()
}
