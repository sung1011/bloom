package server

import (
	"net/http"

	"github.com/sung1011/bloom/fw"
)

// 引擎服务
type FlowerGin struct {
	httpHandler http.Handler
}

// 初始化web引擎服务实例
func BudGin(seed fw.Seed, params ...interface{}) (interface{}, error) {
	sd := seed.(*Seed)
	return &FlowerGin{httpHandler: sd.HttpHandler}, nil
}

// 返回web引擎
func (flw *FlowerGin) HttpHandler() http.Handler {
	return flw.httpHandler
}
