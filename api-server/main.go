package main

import (
	"api-server/container"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		container.Module,
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
	).Run()
}
