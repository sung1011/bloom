package config

import (
	"github.com/sung1011/stickypack/fw"
	"github.com/sung1011/stickypack/fw/svc"
)

type Seed struct {
}

func (sd *Seed) Name() fw.SvcKey {
	return svc.Key_Config
}

func (sd *Seed) Boot(pot fw.Pot) error {
	return nil
}

func (sd *Seed) Register(pot fw.Pot) fw.Bud {
	return Bud
}

func (sd *Seed) Params(pot fw.Pot) []interface{} {
	appEnv := pot.Make(svc.Key_Env).(svc.Env).AppEnv()
	return []interface{}{pot, appEnv}
}

func (sd *Seed) IsDefer() bool {
	return false
}
