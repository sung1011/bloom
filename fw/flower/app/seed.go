package app

import (
	"github.com/sung1011/bloom/fw"
	"github.com/sung1011/bloom/fw/svc"
)

type Seed struct {
	BaseFolder string

	svcUUID svc.UUID
}

func (sd *Seed) Name() fw.SvcKey {
	return svc.Key_App
}

func (sd *Seed) Boot(pot fw.Pot) error {
	sd.svcUUID = pot.Make(svc.Key_UUID).(svc.UUID)
	return nil
}

func (sd *Seed) Register(pot fw.Pot) fw.Bud {
	return Bud
}

func (sd *Seed) Params(pot fw.Pot) []interface{} {
	return []interface{}{pot}
}

func (sd *Seed) IsDefer() bool {
	return false
}
