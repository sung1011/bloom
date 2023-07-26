package fw

import (
	"errors"
	"fmt"
	"sync"
)

type SvcKey string

type Pot interface {
	// Sow 绑定一个服务提供者，如果关键字凭证已经存在，会进行替换操作，返回error
	Sow(seed) error
	// IsSow 关键字凭证是否已经绑定服务提供者
	IsSow(SvcKey) bool
	// Make 根据svcKey从seeds中获取一个服务, 单例模式的; 没有会报错
	Make(SvcKey) interface{}
	// New 根据svcKey从seeds中获取一个服务, 非单例模式的; 没有会报错
	New(SvcKey, []interface{}) (interface{}, error)
}

type TicklesPot struct {
	Pot
	// seeds 存储注册的服务提供者
	seeds map[SvcKey]seed
	// flowers 存储具体的实例
	flowers map[SvcKey]interface{}
	// lock 用于锁住对容器的变更操作
	lock sync.RWMutex
}

func NewTicklesPot() *TicklesPot {
	return &TicklesPot{
		seeds:   make(map[SvcKey]seed),
		flowers: make(map[SvcKey]interface{}),
	}
}

func (pot *TicklesPot) Sow(sd seed) error {
	pot.lock.Lock()
	defer pot.lock.Unlock()

	key := sd.Name()
	if pot.seeds[key] != nil {
		return fmt.Errorf("seed already sow, svcKey: %s", key)
	}
	pot.seeds[key] = sd

	if !sd.IsDefer() {
		if ins, err := pot.newInstance(sd, nil); err == nil {
			pot.flowers[key] = ins
		}
	}
	return nil
}

func (pot *TicklesPot) IsSow(key SvcKey) bool {
	return pot.getSeed(key) != nil
}

func (pot *TicklesPot) Make(key SvcKey) interface{} {
	serv, err := pot.make(key, false, nil)
	if err != nil {
		panic(err)
	}
	return serv
}

func (pot *TicklesPot) New(key SvcKey, params []interface{}) (interface{}, error) {
	return pot.make(key, true, params)
}

func (pot *TicklesPot) Keys() []SvcKey {
	pot.lock.RLock()
	defer pot.lock.RUnlock()

	keys := make([]SvcKey, 0, len(pot.seeds))
	for k := range pot.seeds {
		keys = append(keys, k)
	}
	return keys
}

func (pot *TicklesPot) getSeed(key SvcKey) seed {
	pot.lock.RLock()
	defer pot.lock.RUnlock()
	if sd, ok := pot.seeds[key]; ok {
		return sd
	}
	return nil
}

func (pot *TicklesPot) getInstance(key SvcKey) interface{} {
	pot.lock.RLock()
	defer pot.lock.RUnlock()
	if ins, ok := pot.flowers[key]; ok {
		return ins
	}
	return nil
}

func (pot *TicklesPot) newInstance(sd seed, params []interface{}) (interface{}, error) {
	if err := sd.Boot(pot); err != nil {
		return nil, err
	}
	if params == nil {
		params = sd.Params(pot)
	}
	method := sd.Register(pot)
	if method == nil {
		return nil, fmt.Errorf("seed not sow, svcKey: %s", sd.Name())
	}
	ins, err := method(params...)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return ins, err
}

func (pot *TicklesPot) make(key SvcKey, forceNew bool, params []interface{}) (interface{}, error) {
	sd := pot.getSeed(key)
	if sd == nil {
		return nil, fmt.Errorf("seed not sow, svcKey: %s", key)
	}
	if forceNew {
		return pot.newInstance(sd, params)
	}
	if ins := pot.getInstance(key); ins != nil {
		return ins, nil
	}
	inst, err := pot.newInstance(sd, nil)
	if err != nil {
		return nil, err
	}
	pot.lock.Lock()
	defer pot.lock.Unlock()
	pot.flowers[key] = inst
	return inst, nil
}
