package container

import (
	"api-server/config"
	"api-server/db"
	"api-server/handler"
	"api-server/routes"
	"context"
	"fmt"
	"go.uber.org/fx"
)

var Module = fx.Options(
	handler.Module,
	routes.Module,
	config.Module,
	db.Module,
	fx.Invoke(registerHooks),
)

func registerHooks(lifecycle fx.Lifecycle, h *handler.HttpHandler, config config.Config) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go h.Gin.Run(fmt.Sprintf(":%s", config.Port()))
				return nil
			},

			OnStop: func(ctx context.Context) error {
				// TODO log
				return nil
			},
		},
	)
}
