package main

import (
	"api-server/container"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		container.Module,
	).Run()
}
