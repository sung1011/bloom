package env

import (
	"github.com/sung1011/bloom/fw"
	"github.com/sung1011/bloom/fw/svc"
)

type Seed struct {
}

func (sd *Seed) Name() fw.SvcKey {
	return svc.Key_Env
}

func (sd *Seed) Base(pot fw.Pot) error {
	return nil
}

func (sd *Seed) Build(pot fw.Pot) fw.Bud {
	return Bud
}

func (sd *Seed) IsDefer() bool {
	return false
}
