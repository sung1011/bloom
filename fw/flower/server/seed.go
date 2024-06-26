package server

import (
	"net/http"

	"github.com/sung1011/bloom/fw"
	"github.com/sung1011/bloom/fw/svc"
)

type Seed struct {
	HttpHandler http.Handler
}

func (sd *Seed) Name() fw.SvcKey {
	return svc.Key_Server
}

func (sd *Seed) Base(pot fw.Pot) error {
	return nil
}

func (sd *Seed) Build(pot fw.Pot) fw.Bud {
	return BudGin
}

func (sd *Seed) IsDefer() bool {
	return false
}
