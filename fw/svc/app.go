package svc

const Key_App = "tk:app"

type App interface {
	// AppID 表示当前这个app的唯一id, 可以用于分布式锁等
	AppID() string
	// // Version 定义当前版本
	// Version() string

	// BaseFolder 定义项目基础地址
	BaseFolder() string
	// MetaFolder 定义了配置文件的路径
	MetaFolder() string
	// LogFolder 定义了日志所在路径
	LogFolder() string
	// ServiceFolder 定义业务自己的服务提供者地址
	ServiceFolder() string
	// // MiddlewareFolder 定义业务自己定义的中间件
	// MiddlewareFolder() string
	// // CommandFolder 定义业务定义的命令
	// CommandFolder() string
	// // RuntimeFolder 定义业务的运行中间态信息
	// RuntimeFolder() string
	// // TestFolder 存放测试所需要的信息
	// TestFolder() string
	// // DeployFolder 存放部署的时候创建的文件夹
	// DeployFolder() string

	// // AppFolder 定义业务代码所在的目录，用于监控文件变更使用
	// AppFolder() string
	// // LoadAppConfig 加载新的AppConfig，key为对应的函数转为小写下划线，比如ConfigFolder => config_folder
	// LoadAppConfig(kv map[string]string)
}
