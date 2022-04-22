package container

import (
	"api-server/config"
	"api-server/db"
	"api-server/handler"
	"api-server/routes"
	"context"
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Options(
	handler.Module,
	routes.Module,
	config.Module,
	db.Module,
	//logger.Module,
	fx.Invoke(registerHooks),
)

func registerHooks(lifecycle fx.Lifecycle, h *handler.HttpHandler, config config.Config, logger *zap.SugaredLogger) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go h.Gin.Run(fmt.Sprintf(":%s", config.Port()))
				return nil
			},

			OnStop: func(ctx context.Context) error {
				logger.Info("Container: OnStop called")
				return nil
			},
		},
	)
}
