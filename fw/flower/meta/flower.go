package meta

import (
	"bytes"
	"os"
	"path"

	"github.com/spf13/viper"
	"github.com/sung1011/bloom/fw"
	"github.com/sung1011/bloom/fw/svc"
)

type Flower struct {
	sd fw.Seed
	svc.Meta
}

func Bud(seed fw.Seed, params ...interface{}) (interface{}, error) {
	sd := seed.(*Seed)

	confFile := path.Join(sd.svcApp.MetaFolder(), sd.svcEnv.AppEnv()+".yaml")
	data, err := os.ReadFile(confFile)
	if err != nil {
		return nil, err
	}
	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	return &Flower{sd: sd}, nil
}
func (flw *Flower) Get(key string) interface{} {
	return viper.Get(key)
}

func (flw *Flower) Load(key string, val interface{}) error {
	return viper.UnmarshalKey(key, val)
}
