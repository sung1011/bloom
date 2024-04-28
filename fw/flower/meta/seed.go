package meta

import (
	"github.com/sung1011/bloom/fw"
	"github.com/sung1011/bloom/fw/svc"
)

type Seed struct {
	svcEnv svc.Env
	svcApp svc.App
}

func (sd *Seed) Name() fw.SvcKey {
	return svc.Key_Meta
}

func (sd *Seed) Base(pot fw.Pot) error {
	sd.svcApp = pot.Bloom(svc.Key_App).(svc.App)
	sd.svcEnv = pot.Bloom(svc.Key_Env).(svc.Env)
	return nil
}

func (sd *Seed) Build(pot fw.Pot) fw.Bud {
	return Bud
}

func (sd *Seed) IsDefer() bool {
	return false
}
