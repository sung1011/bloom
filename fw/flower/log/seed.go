package log

import (
	"github.com/sung1011/bloom/fw"
	"github.com/sung1011/bloom/fw/svc"
)

type Seed struct {
	svcMeta svc.Meta
	Driver  string
}

func (sd *Seed) Name() fw.SvcKey {
	return svc.Key_Log
}

func (sd *Seed) Boot(pot fw.Pot) error {
	sd.svcMeta = pot.Make(svc.Key_Meta).(svc.Meta)
	return nil
}

func (sd *Seed) Register(pot fw.Pot) fw.Bud {
	switch sd.Driver {
	case "zap":
		return BudZap
	}
	return BudZap
}

func (sd *Seed) Params(pot fw.Pot) []interface{} {
	return []interface{}{}
}

func (sd *Seed) IsDefer() bool {
	return false
}
