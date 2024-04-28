package fw

type Bud func(Seed) (interface{}, error)

// Seed 代表了一个服务提供者的接口, 其字段为基础属性, 可传入 或 由boot()将依赖的flower填充进来
type Seed interface {
	// Base 在调用实例化服务的时候会调用，可以把一些准备工作: 处理依赖关系, 填充seed默认的字段等
	Base(Pot) error
	// Build 在服务容器中注册了一个实例化服务的方法; 是否在注册的时候就实例化这个服务, 需要参考IsDefer接口
	Build(Pot) Bud
	// Name 代表了这个服务提供者的凭证
	Name() SvcKey
	// Params params定义传递给NewInstance的参数, 可以自定义多个, 建议将container作为第一个参数
	// Params(Pot) []interface{}
	// IsDefer 决定是否在注册的时候实例化这个服务; 如果不是注册的时候实例化, 那就延迟到第一次make的时候进行实例化操作
	IsDefer() bool
}
