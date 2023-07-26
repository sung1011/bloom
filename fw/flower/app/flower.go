package app

import (
	"errors"
	"path/filepath"

	"github.com/sung1011/stickypack/fw"
	"github.com/sung1011/stickypack/fw/svc"
)

type Flower struct {
	svc.App
	pot        fw.Pot // 服务容器
	baseFolder string // 基础路径
	appId      string // 表示当前这个app的唯一id, 可以用于分布式锁等
}

func Bud(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("app service need 2 params")
	}
	pot := params[0].(fw.Pot)
	baseFolder := params[1].(string)
	return &Flower{
		pot:        pot,
		baseFolder: baseFolder,
	}, nil
}

func (flw *Flower) AppID() string {
	return flw.appId
}

// BaseFolder 定义项目基础地址
func (flw *Flower) BaseFolder() string {
	return flw.baseFolder
}

// ConfigFolder 定义了配置文件的路径
func (flw *Flower) ConfigFolder() string {
	// envSvc := flw.pot.Make(svc.Key_Env).(svc.Env)
	// return filepath.Join(flw.BaseFolder(), "config", envSvc.AppEnv())
	return filepath.Join(flw.BaseFolder(), "config")
}

// LogFolder 定义了日志所在路径
func (flw *Flower) LogFolder() string {
	return filepath.Join(flw.BaseFolder(), "log")
}

// ProviderFolder 定义业务自己的服务提供者地址
func (flw *Flower) ServiceFolder() string {
	return filepath.Join(flw.BaseFolder(), "service")
}
