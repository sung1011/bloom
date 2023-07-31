package log

import (
	"context"
	"fmt"

	"github.com/sung1011/bloom/fw"
	"go.uber.org/zap"
)

type FlowerZap struct {
	sd         fw.Seed
	rootLogger *zap.SugaredLogger
}

type Config struct {
	Level     string
	Format    string
	Output    []string
	ErrOutput []string
}

func BudZap(seed fw.Seed, params ...interface{}) (interface{}, error) {
	sd := seed.(*Seed)
	var err error
	cfg := zap.NewDevelopmentConfig()

	conf := &Config{}
	if err := sd.svcMeta.Load("app.log", conf); err != nil {
		return nil, err
	}
	if cfg.Level, err = zap.ParseAtomicLevel(conf.Level); err != nil {
		return nil, err
	}
	cfg.Development = sd.svcMeta.Get("env") == "dev"
	cfg.Encoding = conf.Format
	cfg.OutputPaths = conf.Output
	cfg.ErrorOutputPaths = conf.ErrOutput

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	fmt.Println("zap logger", cfg.OutputPaths, cfg.ErrorOutputPaths)
	logger.Info("asd")
	// logger.Warn("abc")
	logger.Error("xyz")
	return &FlowerZap{
		sd:         seed.(*Seed),
		rootLogger: logger.Sugar(),
	}, nil
}

// Panic 表示会导致整个程序出现崩溃的日志信息
func (flw *FlowerZap) Panic(ctx context.Context, msg string, fields map[string]interface{}) {
	panic("not implemented") // TODO: Implement
}

// Fatal 表示会导致当前这个请求出现提前终止的错误信息
func (flw *FlowerZap) Fatal(ctx context.Context, msg string, fields map[string]interface{}) {
	panic("not implemented") // TODO: Implement
}

// Error 表示出现错误，但是不一定影响后续请求逻辑的错误信息
func (flw *FlowerZap) Error(ctx context.Context, msg string, fields map[string]interface{}) {
	panic("not implemented") // TODO: Implement
}

// Warn 表示出现错误，但是一定不影响后续请求逻辑的报警信息
func (flw *FlowerZap) Warn(ctx context.Context, msg string, fields map[string]interface{}) {
	panic("not implemented") // TODO: Implement
}

// Info 表示正常的日志信息输出
func (flw *FlowerZap) Info(ctx context.Context, msg string, fields map[string]interface{}) {
	flw.rootLogger.Infow(msg, fields)
}

// Debug 表示在调试状态下打印出来的日志信息
func (flw *FlowerZap) Debug(ctx context.Context, msg string, fields map[string]interface{}) {
	panic("not implemented") // TODO: Implement
}

// Trace 表示最详细的信息，一般信息量比较大，可能包含调用堆栈等信息
func (flw *FlowerZap) Trace(ctx context.Context, msg string, fields map[string]interface{}) {
	panic("not implemented") // TODO: Implement
}

// SetLevel 设置日志级别
// // SetCtxFielder 从context中获取上下文字段field
// SetCtxFielder(handler CtxFielder)
// // SetFormatter 设置输出格式
// SetFormatter(formatter Formatter)
// // SetOutput 设置输出管道
// SetOutput(out io.Writer)
