package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sung1011/bloom/app/http"
	"github.com/sung1011/bloom/fw"
	"github.com/sung1011/bloom/fw/flower/app"
	"github.com/sung1011/bloom/fw/flower/env"
	"github.com/sung1011/bloom/fw/flower/log"
	"github.com/sung1011/bloom/fw/flower/meta"
	"github.com/sung1011/bloom/fw/flower/redis"
	"github.com/sung1011/bloom/fw/flower/server"
	"github.com/sung1011/bloom/fw/flower/uuid"
	"github.com/sung1011/bloom/fw/svc"
)

func main() {
	pot := fw.NewTicklesPot()
	if err := pot.Sow(&env.Seed{}); err != nil {
		panic(err)
	}
	if err := pot.Sow(&uuid.Seed{Driver: "google"}); err != nil {
		panic(err)
	}
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if err := pot.Sow(&app.Seed{BaseFolder: dir}); err != nil { // app on: uuid
		panic(err)
	}
	if err := pot.Sow(&meta.Seed{}); err != nil { // meta on: app, env
		panic(err)
	}
	// server
	httpHandler, err := http.NewHttpEngine(pot)
	if err != nil {
		panic(err)
	}
	if err := pot.Sow(&server.Seed{HttpHandler: httpHandler}); err != nil {
		panic(err)
	}
	if err := pot.Sow(&log.Seed{}); err != nil {
		panic(err)
	}
	if err := pot.Sow(&redis.Seed{}); err != nil {
		panic(err)
	}

	// @@ config, deploy, mw, mongo, redis, server(kernel)
	fmt.Println("-------------------------------------")
	// tmpMeta(pot)
	// tmpZap(pot)
	tmpRedis(pot)
}

func tmpRedis(pot fw.Pot) {
	redis := pot.Make(svc.Key_Redis).(svc.Redis)
	c, err := redis.GetClient("default")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	fmt.Println(c.Ping(ctx).Result())
}

func tmpMeta(pot fw.Pot) {
	// yaml := struct {
	// 	Env string
	// }{}
	// pot.Make(svc.Key_Meta).(svc.Meta).Load("app", &yaml)
	// fmt.Println("meta", yaml.Env)
	// fmt.Println("app", pot.Make(svc.Key_App).(svc.App).MetaFolder())
}

func tmpZap(pot fw.Pot) {
	ctx := context.Background()
	m := map[string]interface{}{
		"a": "b",
	}
	pot.Make(svc.Key_Log).(svc.Log).Info(ctx, "lala", m)
	// zap yaml
	// 	rawJSON := []byte(`{
	//     "level":"error",
	//     "encoding":"json",
	//     "outputPaths": ["stdout", "server.log"],
	//     "errorOutputPaths": ["stderr"],
	//     "initialFields":{"init":"dj"},
	//     "encoderConfig": {
	//       "messageKey": "message",
	//       "levelKey": "level",
	//       "levelEncoder": "lowercase"
	//     }
	//   }`)
	// 	var cfg zap.Config
	// 	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
	// 		panic(err)
	// 	}
	// 	logger, err := cfg.Build()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer logger.Sync()

	// 	logger.Info("server start work successfully!")

	// 	logger.Info("lala")

	// zap config
	// cfg := zap.Config{
	// 	Level:             zap.AtomicLevel{},
	// 	Development:       false,
	// 	DisableCaller:     false,
	// 	DisableStacktrace: false,
	// 	Sampling: &zap.SamplingConfig{
	// 		Initial:    0,
	// 		Thereafter: 0,
	// 		Hook:       func(zapcore.Entry, zapcore.SamplingDecision) { panic("not implemented") },
	// 	},
	// 	Encoding: "",
	// 	EncoderConfig: zapcore.EncoderConfig{
	// 		MessageKey:          "",
	// 		LevelKey:            "",
	// 		TimeKey:             "",
	// 		NameKey:             "",
	// 		CallerKey:           "",
	// 		FunctionKey:         "",
	// 		StacktraceKey:       "",
	// 		SkipLineEnding:      false,
	// 		LineEnding:          "",
	// 		EncodeLevel:         func(zapcore.Level, zapcore.PrimitiveArrayEncoder) { panic("not implemented") },
	// 		EncodeTime:          func(time.Time, zapcore.PrimitiveArrayEncoder) { panic("not implemented") },
	// 		EncodeDuration:      func(time.Duration, zapcore.PrimitiveArrayEncoder) { panic("not implemented") },
	// 		EncodeCaller:        func(zapcore.EntryCaller, zapcore.PrimitiveArrayEncoder) { panic("not implemented") },
	// 		EncodeName:          func(string, zapcore.PrimitiveArrayEncoder) { panic("not implemented") },
	// 		NewReflectedEncoder: func(io.Writer) zapcore.ReflectedEncoder { panic("not implemented") },
	// 		ConsoleSeparator:    "",
	// 	},
	// 	OutputPaths:      []string{},
	// 	ErrorOutputPaths: []string{},
	// 	InitialFields: map[string]interface{}{
	// 		"": nil,
	// 	},
	// }
	// logger, err := cfg.Build()
	// if err != nil {
	// 	panic(err)
	// }
	// logger.Info("halo")

}
