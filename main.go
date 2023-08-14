package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/sung1011/bloom/app/cmd"
	appHttp "github.com/sung1011/bloom/app/http"
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
	httpHandler, err := appHttp.NewHttpEngine(pot)
	if err != nil {
		panic(err)
	}
	if err := pot.Sow(&server.Seed{HttpHandler: httpHandler}); err != nil {
		panic(err)
	}
	if err := pot.Sow(&log.Seed{}); err != nil {
		panic(err)
	}
	defer pot.Make(svc.Key_Log).(svc.Log).Sync()

	if err := pot.Sow(&redis.Seed{}); err != nil {
		panic(err)
	}

	// @@ config, deploy, mw, mongo, server(kernel)
	fmt.Println("-------------------------------------")

	test_Meta(pot)
	test_Zap(pot)
	test_Redis(pot)

	// test_Server(pot)
	_ = cmd.RunCommand(pot)

}

func test_Server(pot fw.Pot) {
	server := pot.Make(svc.Key_Server).(svc.Server)
	http.ListenAndServe(":8080", server.HttpHandler())
}

func test_Redis(pot fw.Pot) {
	redis := pot.Make(svc.Key_Redis).(svc.Redis)
	c, err := redis.GetClient("default")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	fmt.Println(c.Ping(ctx).Result())
}

func test_Meta(pot fw.Pot) {
	meta := pot.Make(svc.Key_Meta).(svc.Meta).Get()
	fmt.Println("env", meta.Env)
}

func test_Zap(pot fw.Pot) {
	ctx := context.Background()
	m := map[string]interface{}{
		"a": "b",
	}
	logger := pot.Make(svc.Key_Log).(svc.Log)

	logger.Info(ctx, "lala", m)
	// logger.Fatal(ctx, "fff", m)
	// logger.Error(ctx, "eee", m)
	// logger.Panic(ctx, "ppp", m)
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
}
