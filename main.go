package main

import (
	"fmt"
	"os"

	"github.com/sung1011/bloom/fw"
	"github.com/sung1011/bloom/fw/flower/app"
	"github.com/sung1011/bloom/fw/flower/env"
	"github.com/sung1011/bloom/fw/flower/meta"
	"github.com/sung1011/bloom/fw/flower/server"
	"github.com/sung1011/bloom/fw/flower/uuid"
	"github.com/sung1011/bloom/fw/svc"
)

func main() {
	pot := fw.NewTicklesPot()
	if err := pot.Sow(&env.Seed{}); err != nil {
		panic(err)
	}
	if err := pot.Sow(&uuid.Seed{Mode: "google"}); err != nil {
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
	if err := pot.Sow(&server.Seed{Mode: "gin"}); err != nil {
		panic(err)
	}

	// @@ config, deploy, log, mongo, redis, server(kernel)

	fmt.Println("-------------------------------------")
	fmt.Println("app", pot.Make(svc.Key_App).(svc.App).MetaFolder())

	yaml := struct {
		Env string
	}{}
	pot.Make(svc.Key_Meta).(svc.Meta).Load(&yaml)
	fmt.Println("meta", yaml.Env)

	fmt.Println("", pot.Make(svc.Key_Server).(svc.Server).HttpHandler())
	// viper.GetViper().Debug()
}
