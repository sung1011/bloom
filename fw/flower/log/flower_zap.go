package log

import (
	"context"
	"io"
	"os"
	"path"

	"github.com/sung1011/bloom/fw"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type FlowerZap struct {
	sd         fw.Seed
	rootLogger *zap.SugaredLogger
}

func BudZap(seed fw.Seed, params ...interface{}) (interface{}, error) {
	sd := seed.(*Seed)
	meta := sd.svcMeta.Get()
	cfg := zap.NewDevelopmentConfig()

	// cfg.Development = meta.Env == "dev"
	// cfg.Encoding = meta.App.Log.Format
	// cfg.OutputPaths = []string{"stdout"}
	// cfg.OutputPaths = append(cfg.OutputPaths, path.Join(sd.svcApp.LogFolder(), "access.log"))
	// cfg.ErrorOutputPaths = []string{"stdout"}
	// cfg.ErrorOutputPaths = append(cfg.ErrorOutputPaths, path.Join(sd.svcApp.LogFolder(), "error.log"))
	// logger, err := cfg.Build()
	// if err != nil {
	// 	return nil, err
	// }

	var err error
	var encoder zapcore.Encoder
	if meta.App.Log.Format == "json" {
		encoder = zapcore.NewJSONEncoder(cfg.EncoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(cfg.EncoderConfig)
	}
	// access log
	accessPath := path.Join(sd.svcApp.LogFolder(), "access.log")
	c1 := zapcore.NewCore(encoder, zapcore.AddSync(getAccessLogWriter(accessPath)), zapcore.DebugLevel)
	if err != nil {
		return nil, err
	}
	// error log
	errorPath := path.Join(sd.svcApp.LogFolder(), "error.log")
	c2 := zapcore.NewCore(encoder, zapcore.AddSync(getErrorLogWriter(errorPath)), zapcore.ErrorLevel)
	if err != nil {
		return nil, err
	}
	// logger
	core := zapcore.NewTee(c1, c2)
	logger := zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.ErrorLevel),
	)

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

func (flw *FlowerZap) Sync() error {
	return flw.rootLogger.Sync()
}

func getAccessLogWriter(logPath string) zapcore.WriteSyncer {
	file, _ := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// file, _ := os.Create(logPath)
	ws := io.MultiWriter(file, os.Stdout)
	return zapcore.AddSync(ws)
}

func getErrorLogWriter(logPath string) zapcore.WriteSyncer {
	file, _ := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// file, _ := os.Create(logPath)
	ws := io.MultiWriter(file, os.Stderr)
	return zapcore.AddSync(ws)
}
