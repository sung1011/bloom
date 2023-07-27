package uuid

import (
	"github.com/sung1011/bloom/fw"
	"github.com/sung1011/bloom/fw/svc"
)

type Seed struct {
	Mode string
}

func (sd *Seed) Name() fw.SvcKey {
	return svc.Key_UUID
}

func (sd *Seed) Boot(pot fw.Pot) error {
	return nil
}

func (sd *Seed) Register(pot fw.Pot) fw.Bud {
	switch sd.Mode {
	case "google":
		return BudGoole
	case "gofrs":
		return BudGofrs
	}
	return nil
}

func (sd *Seed) Params(pot fw.Pot) []interface{} {
	return []interface{}{}
}

func (sd *Seed) IsDefer() bool {
	return false
}
