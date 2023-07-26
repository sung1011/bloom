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

func (sd *Seed) Boot(pot fw.Pot) error {
	return nil
}

func (sd *Seed) Register(pot fw.Pot) fw.Bud {
	return Bud
}

func (sd *Seed) Params(pot fw.Pot) []interface{} {
	return []interface{}{pot, sd.BaseFolder}
}

func (sd *Seed) IsDefer() bool {
	return false
}
