package env

import (
	"github.com/sung1011/stickypack/fw"
	"github.com/sung1011/stickypack/fw/svc"
)

type Seed struct {
	fw.Pot
}

func (sd *Seed) Name() fw.SvcKey {
	return svc.Key_Env
}

func (sd *Seed) Boot(c fw.Pot) error {
	return nil
}

func (sd *Seed) Register(c fw.Pot) fw.Bud {
	return Bud
}

func (sd *Seed) Params(c fw.Pot) []interface{} {
	return []interface{}{}
}

func (sd *Seed) IsDefer() bool {
	return false
}
