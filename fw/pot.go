package fw

import (
	"bytes"
	"errors"
	"fmt"
	"sync"
)

type SvcKey string

type Pot interface {
	// Sow 绑定一个服务提供者，如果关键字凭证已经存在，会进行替换操作，返回error; 无锁, 启动服务时候调用
	Sow(Seed) error
	// Bloom 根据svcKey从seeds中获取一个服务, 单例模式的; 没有会报错
	Bloom(SvcKey) interface{}
}

type TicklesPot struct {
	Pot
	// seeds 存储注册的服务提供者
	seeds map[SvcKey]Seed
	// flowers 存储具体的实例
	flowers map[SvcKey]interface{}
	// lock 用于锁住对容器的变更操作
	lock sync.RWMutex
}

func NewTicklesPot() *TicklesPot {
	return &TicklesPot{
		seeds:   make(map[SvcKey]Seed),
		flowers: make(map[SvcKey]interface{}),
	}
}

func (pot *TicklesPot) Sow(sd Seed) error {
	if err := pot.setSeed(sd); err != nil {
		return err
	}
	if !sd.IsDefer() {
		if _, err := pot.getFlower(sd); err != nil {
			return err
		}
	}
	return nil
}

func (pot *TicklesPot) Bloom(key SvcKey) interface{} {
	pot.lock.Lock()
	defer pot.lock.Unlock()
	// svc, err := pot.make(key)
	sd := pot.getSeed(key)
	if sd == nil {
		panic(fmt.Sprintf("svc not found, svcKey: %s", key))
	}
	svc, err := pot.getFlower(sd)

	if err != nil {
		panic(err)
	}
	return svc
}

func (pot *TicklesPot) Pretty() {
	pretty := func(v interface{}) string {
		if v == nil {
			return "<nil>"
		}
		return fmt.Sprintf("%#v", v)
	}
	var buf bytes.Buffer
	buf.WriteString("seeds:\n")
	for k, v := range pot.seeds {
		buf.WriteString(fmt.Sprintf("  %s: %s\n", k, pretty(v)))
	}
	buf.WriteString("flowers:\n")
	for k, v := range pot.flowers {
		buf.WriteString(fmt.Sprintf("  %s: %s\n", k, pretty(v)))
	}
	fmt.Printf("%v", buf.String())
}

func (pot *TicklesPot) getSeed(key SvcKey) Seed {
	if sd, ok := pot.seeds[key]; ok {
		return sd
	}
	return nil
}

func (pot *TicklesPot) setSeed(sd Seed) error {
	key := sd.Name()
	if pot.seeds[key] != nil {
		return fmt.Errorf("seed already sow, svcKey: %s", key)
	}
	pot.seeds[key] = sd
	return nil
}

func (pot *TicklesPot) getFlower(sd Seed) (interface{}, error) {
	key := sd.Name()
	if flower, ok := pot.flowers[key]; ok {
		return flower, nil
	}
	flower, err := pot.newFlower(sd)
	if err != nil {
		return nil, err
	}
	pot.flowers[key] = flower

	return flower, nil
}

func (pot *TicklesPot) newFlower(sd Seed) (interface{}, error) {
	if err := sd.Base(pot); err != nil {
		return nil, err
	}
	bud := sd.Build(pot)
	if bud == nil {
		return nil, fmt.Errorf("seed not sow, svcKey: %s", sd.Name())
	}
	flower, err := bud(sd)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return flower, err
}
