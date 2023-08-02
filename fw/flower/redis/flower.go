package meta

import (
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/sung1011/bloom/fw"
	"github.com/sung1011/bloom/fw/svc"
)

type Flower struct {
	svc.Redis // implements
	sd        fw.Seed
	cMap      map[string]*redis.Client
	csMap     map[string][]*redis.Client
}

func Bud(seed fw.Seed, params ...interface{}) (interface{}, error) {
	sd := seed.(*Seed)
	// c
	cMap := map[string]*redis.Client{}
	for _, rds := range sd.svcMeta.Get().App.Storage.Redis {
		opt, err := redis.ParseURL(rds.URI)
		if err != nil {
			return nil, err
		}
		cMap[rds.Name] = redis.NewClient(&redis.Options{
			Addr:     opt.Addr,
			Password: opt.Password,
			DB:       rds.DB,
			// @@
		})
	}
	// cs
	return &Flower{sd: sd}, nil
}

func (flw *Flower) GetClient(key string) (*redis.Client, error) {
	c := flw.cMap[key]
	if c == nil {
		return nil, errors.New("redis client not found " + key)
	}
	return c, nil
}

func (flw *Flower) GetClientShared(key string, idx int) (*redis.Client, error) {
	cs := flw.csMap[key]
	if cs == nil {
		return nil, errors.New("redis client not found " + key)
	}
	c := cs[idx%len(cs)]
	return c, nil
}
