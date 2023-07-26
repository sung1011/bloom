package env

import (
	"os"
	"strings"

	"github.com/sung1011/stickypack/fw/svc"
)

type Flower struct {
	svc.Env
	maps map[string]string
}

func Bud(...interface{}) (interface{}, error) {
	svc := &Flower{
		maps: map[string]string{},
	}
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if len(pair) < 2 {
			continue
		}
		svc.maps[pair[0]] = pair[1]
	}
	return svc, nil
}

func (flw *Flower) AppEnv() string {
	return flw.maps["APP_ENV"]
}
