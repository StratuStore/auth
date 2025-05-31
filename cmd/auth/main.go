package main

import (
	"github.com/StratuStore/auth/internal/app/app"
	"github.com/StratuStore/auth/internal/libs/config"
	"go.uber.org/fx"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	fx.New(app.CreateApp(cfg)).Run()
}
