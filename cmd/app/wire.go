//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/google/wire"

	"github.com/pranc1ngpegasus/go-playground/adapter/handler"
	"github.com/pranc1ngpegasus/go-playground/repository"
	"github.com/pranc1ngpegasus/go-playground/usecase"
)

type app struct {
	handler http.Handler
}

func initialize() (*app, error) {
	wire.Build(
		// repository
		repository.NewFindByIDImplSet,

		// usecase
		usecase.NewFindByIDImplSet,

		// handler
		handler.NewHandlerSet,

		wire.Struct(new(app), "*"),
	)

	return nil, nil
}
