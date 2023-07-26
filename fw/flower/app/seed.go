package app

import (
	"github.com/sung1011/stickypack/fw"
	"github.com/sung1011/stickypack/fw/svc"
)

type Seed struct {
	fw.Pot
	BaseFolder string
}

func (sd *Seed) Name() fw.SvcKey {
	return svc.Key_App
}

func (sd *Seed) Boot(c fw.Pot) error {
	return nil
}

func (sd *Seed) Register(c fw.Pot) fw.Bud {
	return Bud
}

func (sd *Seed) Params(c fw.Pot) []interface{} {
	return []interface{}{c, sd.BaseFolder}
}

func (sd *Seed) IsDefer() bool {
	return false
}
