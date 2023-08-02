package redis

import (
	"github.com/sung1011/bloom/fw"
	"github.com/sung1011/bloom/fw/svc"
)

type Seed struct {
	svcMeta svc.Meta
	svcLog  svc.Log
}

func (sd *Seed) Name() fw.SvcKey {
	return svc.Key_Redis
}

func (sd *Seed) Boot(pot fw.Pot) error {
	sd.svcMeta = pot.Make(svc.Key_Meta).(svc.Meta)
	sd.svcLog = pot.Make(svc.Key_Log).(svc.Log)
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
