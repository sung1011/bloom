package main

import (
	"fmt"
	"os"

	"github.com/sung1011/stickypack/fw"
	"github.com/sung1011/stickypack/fw/flower/app"
	"github.com/sung1011/stickypack/fw/flower/env"
	"github.com/sung1011/stickypack/fw/flower/uuid"
	"github.com/sung1011/stickypack/fw/svc"
)

func main() {
	pot := fw.NewTicklesPot()
	// env
	pot.Sow(&env.Seed{})
	// config
	// app
	dir, err := os.Getwd() // @@ to boot or config
	if err != nil {
		panic(err)
	}
	pot.Sow(&app.Seed{BaseFolder: dir})
	// ...
	pot.Sow(&uuid.Seed{})

	fmt.Println("", pot.Keys())

	s := pot.Make(svc.Key_UUID).(svc.UUID)

	o, err := pot.New(svc.Key_UUID, []interface{}{})
	if err != nil {
		panic(err)
	}

	fmt.Println("", s, s.NewID())
	fmt.Println("", o, o.(svc.UUID).NewID())
	fmt.Println("", pot.Make(svc.Key_App).(svc.App).BaseFolder())
	fmt.Println("", pot.Make(svc.Key_App).(svc.App).ConfigFolder())
}
