package env

import (
	"os"
	"strings"

	"github.com/sung1011/bloom/fw"
)

type Flower struct {
	sd fw.Seed

	maps map[string]string
}

func Bud(seed fw.Seed) (interface{}, error) {
	maps := map[string]string{}
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if len(pair) < 2 {
			continue
		}
		maps[pair[0]] = pair[1]
	}

	return &Flower{
		sd:   seed.(*Seed),
		maps: maps,
	}, nil
}

func (flw *Flower) AppEnv() string {
	return flw.maps["APP_ENV"]
}
