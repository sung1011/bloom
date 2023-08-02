package log

import (
	"context"

	"github.com/sung1011/bloom/fw"
	"github.com/sung1011/bloom/fw/svc"
	"go.uber.org/zap"
)

type FlowerZap struct {
	svc.Log    // implements
	sd         fw.Seed
	rootLogger *zap.SugaredLogger
}

func BudZap(seed fw.Seed, params ...interface{}) (interface{}, error) {
	sd := seed.(*Seed)
	var err error
	cfg := zap.NewDevelopmentConfig()

	meta := sd.svcMeta.Get()
	cfg.Development = meta.Env == "dev"
	cfg.Encoding = meta.App.Log.Format
	cfg.OutputPaths = meta.App.Log.Output
	cfg.ErrorOutputPaths = meta.App.Log.ErrOutput

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	return &FlowerZap{
		sd:         seed.(*Seed),
		rootLogger: logger.Sugar(),
	}, nil
}

// Panic 表示会导致整个程序出现崩溃的日志信息
func (flw *FlowerZap) Panic(ctx context.Context, msg string, fields map[string]interface{}) {
	var kvs []interface{}
	for k, v := range fields {
		kvs = append(kvs, k, v)
	}
	flw.rootLogger.Panicw(msg, kvs...)
}

// Fatal 表示会导致当前这个请求出现提前终止的错误信息
func (flw *FlowerZap) Fatal(ctx context.Context, msg string, fields map[string]interface{}) {
	var kvs []interface{}
	for k, v := range fields {
		kvs = append(kvs, k, v)
	}
	flw.rootLogger.Fatalw(msg, kvs...)
}

// Error 表示出现错误，但是不一定影响后续请求逻辑的错误信息
func (flw *FlowerZap) Error(ctx context.Context, msg string, fields map[string]interface{}) {
	var kvs []interface{}
	for k, v := range fields {
		kvs = append(kvs, k, v)
	}
	flw.rootLogger.Errorw(msg, kvs...)
}

// Warn 表示出现错误，但是一定不影响后续请求逻辑的报警信息
func (flw *FlowerZap) Warn(ctx context.Context, msg string, fields map[string]interface{}) {
	var kvs []interface{}
	for k, v := range fields {
		kvs = append(kvs, k, v)
	}
	flw.rootLogger.Warnw(msg, kvs...)
}

// Info 表示正常的日志信息输出
func (flw *FlowerZap) Info(ctx context.Context, msg string, fields map[string]interface{}) {
	var kvs []interface{}
	for k, v := range fields {
		kvs = append(kvs, k, v)
	}
	flw.rootLogger.Infow(msg, kvs...)
}

// Debug 表示在调试状态下打印出来的日志信息
func (flw *FlowerZap) Debug(ctx context.Context, msg string, fields map[string]interface{}) {
	var kvs []interface{}
	for k, v := range fields {
		kvs = append(kvs, k, v)
	}
	flw.rootLogger.Debugw(msg, kvs...)
}

// Trace 表示最详细的信息，一般信息量比较大，可能包含调用堆栈等信息
// func (flw *FlowerZap) Trace(ctx context.Context, msg string, fields map[string]interface{}) {
// }

// SetLevel 设置日志级别
// // SetCtxFielder 从context中获取上下文字段field
// SetCtxFielder(handler CtxFielder)
// // SetFormatter 设置输出格式
// SetFormatter(formatter Formatter)
// // SetOutput 设置输出管道
// SetOutput(out io.Writer)
