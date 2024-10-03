package main

import (
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		// Provide registers any number of constructor functions, teaching the
		fx.Provide(
			NewServerConfig,
			NewHandlerMapping,
			SetupRouter,
			NewHandler, NewService, NewRepo,
		),
		fx.Invoke(StartGinServer),
	)

	app.Run()
}
