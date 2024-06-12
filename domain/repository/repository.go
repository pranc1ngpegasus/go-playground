package repository

import (
	"context"
)

type Repository[Input any, Output any] interface {
	Handle(context.Context, Input) (Output, error)
}
