// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/pranc1ngpegasus/go-playground/adapter/handler"
	"github.com/pranc1ngpegasus/go-playground/repository"
	"github.com/pranc1ngpegasus/go-playground/usecase"
	"net/http"
)

// Injectors from wire.go:

func initialize() (*app, error) {
	findByIDImpl := repository.NewFindByIDImpl()
	usecaseFindByIDImpl := usecase.NewFindByIDImpl(findByIDImpl)
	handlerHandler, err := handler.NewHandler(usecaseFindByIDImpl)
	if err != nil {
		return nil, err
	}
	mainApp := &app{
		handler: handlerHandler,
	}
	return mainApp, nil
}

// wire.go:

type app struct {
	handler http.Handler
}
