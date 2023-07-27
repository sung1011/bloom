package server

import (
	"net/http"

	"github.com/gohade/hade/framework/gin"
	"github.com/sung1011/bloom/fw"
)

// 引擎服务
type FlowerGin struct {
	engine *gin.Engine
}

// 初始化web引擎服务实例
func BudGin(seed fw.Seed, params ...interface{}) (interface{}, error) {
	sd := seed.(*Seed)
	return &FlowerGin{engine: sd.ginEngine}, nil
}

// 返回web引擎
func (s *FlowerGin) HttpHandler() http.Handler {
	return s.engine
}
