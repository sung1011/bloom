package env

import (
	"github.com/sung1011/bloom/fw"
	"github.com/sung1011/bloom/fw/svc"
)

type Seed struct {
	fw.Pot
}

func (sd *Seed) Name() fw.SvcKey {
	return svc.Key_Env
}

func (sd *Seed) Boot(pot fw.Pot) error {
	return nil
}

func (sd *Seed) Register(pot fw.Pot) fw.Bud {
	return Bud
}

func (sd *Seed) Params(pot fw.Pot) []interface{} {
	return []interface{}{}
}

func (sd *Seed) IsDefer() bool {
	return false
}
