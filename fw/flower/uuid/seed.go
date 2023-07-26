package uuid

import (
	"github.com/sung1011/stickypack/fw"
	"github.com/sung1011/stickypack/fw/svc"
)

type Seed struct {
	fw.Pot
	Mode string
}

func (sd *Seed) Name() fw.SvcKey {
	return svc.Key_UUID
}

func (sd *Seed) Boot(c fw.Pot) error {
	sd.Mode = "gofrs" // @@todo get config
	// s.Mode = "google" // @@todo get config
	return nil
}

func (sd *Seed) Register(c fw.Pot) fw.Bud {
	switch sd.Mode {
	case "google":
		return BudGoole
	case "gofrs":
		return BudGofrs
	}
	return nil
}

func (sd *Seed) Params(p fw.Pot) []interface{} {
	return []interface{}{}
}

func (sd *Seed) IsDefer() bool {
	return false
}
