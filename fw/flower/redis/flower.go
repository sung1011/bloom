package redis

import (
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/sung1011/bloom/fw"
)

type Flower struct {
	sd    fw.Seed
	cMap  map[string]*redis.Client
	csMap map[string][]*redis.Client
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
		if nil != cMap[rds.Name] {
			return nil, errors.New("duplicate redis name " + rds.Name)
		}
		opt.DB = rds.DB
		cMap[rds.Name] = redis.NewClient(opt)
	}
	// cs
	csMap := map[string][]*redis.Client{}
	for _, rds := range sd.svcMeta.Get().App.Storage.SharedRedis {
		if nil != csMap[rds.Name] {
			return nil, errors.New("duplicate redis shared name " + rds.Name)
		}
		for k, url := range rds.URLs {
			opt, err := redis.ParseURL(url)
			if err != nil {
				return nil, err
			}
			opt.DB = rds.DBs[k]
			// csMap[rds.Name] = []*redis.Client{}
			csMap[rds.Name] = append(csMap[rds.Name], redis.NewClient(opt))
		}
	}

	return &Flower{
		sd:    sd,
		cMap:  cMap,
		csMap: csMap,
	}, nil
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
