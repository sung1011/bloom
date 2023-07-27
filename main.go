package main

import (
	"fmt"
	"os"

	"github.com/sung1011/bloom/fw"
	"github.com/sung1011/bloom/fw/flower/app"
	"github.com/sung1011/bloom/fw/flower/env"
	"github.com/sung1011/bloom/fw/flower/meta"
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

	o, err := pot.New(svc.Key_UUID, []interface{}{})
	if err != nil {
		panic(err)
	}

	fmt.Println("-------------------------------------")
	s := pot.Make(svc.Key_UUID).(svc.UUID)
	fmt.Println("", s, s.NewID())
	fmt.Println("", o, o.(svc.UUID).NewID())
	fmt.Println("", pot.Make(svc.Key_App).(svc.App).BaseFolder())
	fmt.Println("", pot.Make(svc.Key_App).(svc.App).MetaFolder())
}
