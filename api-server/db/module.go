package db

import "go.uber.org/fx"

var Module = fx.Options(fx.Provide(
	NewUserDao,
	GetDb,
))
