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

func (sd *Seed) Base(pot fw.Pot) error {
	sd.svcMeta = pot.Bloom(svc.Key_Meta).(svc.Meta)
	sd.svcLog = pot.Bloom(svc.Key_Log).(svc.Log)
	return nil
}

func (sd *Seed) Build(pot fw.Pot) fw.Bud {
	return Bud
}

func (sd *Seed) IsDefer() bool {
	return false
}
