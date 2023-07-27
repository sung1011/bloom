package server

import (
	"github.com/gohade/hade/framework/gin"
	"github.com/sung1011/bloom/fw"
	"github.com/sung1011/bloom/fw/svc"
)

type Seed struct {
	Mode string

	ginEngine *gin.Engine
}

func (sd *Seed) Name() fw.SvcKey {
	return svc.Key_Server
}

func (sd *Seed) Boot(pot fw.Pot) error {
	return nil
}

func (sd *Seed) Register(pot fw.Pot) fw.Bud {
	switch sd.Mode {
	case "gin":
		return BudGin
	}
	return nil
}

func (sd *Seed) Params(pot fw.Pot) []interface{} {
	return []interface{}{}
}

func (sd *Seed) IsDefer() bool {
	return false
}
