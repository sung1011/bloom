package app

import (
	"path/filepath"

	"github.com/sung1011/bloom/fw"
)

type Flower struct {
	sd fw.Seed

	pot        fw.Pot // 服务容器
	baseFolder string // 基础路径
	appId      string // 表示当前这个app的唯一id, 可以用于分布式锁等
}

func Bud(seed fw.Seed, params ...interface{}) (interface{}, error) {
	sd := seed.(*Seed)
	pot := params[0].(fw.Pot)

	return &Flower{
		pot:        pot,
		baseFolder: sd.BaseFolder,
		appId:      sd.svcUUID.NewID(),
		sd:         sd,
	}, nil
}

func (flw *Flower) AppID() string {
	return flw.appId
}

// BaseFolder 定义项目基础地址
func (flw *Flower) BaseFolder() string {
	return flw.baseFolder
}

// MetaFolder 定义了配置文件的路径
func (flw *Flower) MetaFolder() string {
	return filepath.Join(flw.BaseFolder(), "meta")
}

// LogFolder 定义了日志所在路径
func (flw *Flower) LogFolder() string {
	return filepath.Join(flw.BaseFolder(), "log")
}

// ProviderFolder 定义业务自己的服务提供者地址
func (flw *Flower) ServiceFolder() string {
	return filepath.Join(flw.BaseFolder(), "service")
}
