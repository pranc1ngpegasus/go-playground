package usecase

import (
	"context"
)

type Usecase[Input any, Output any] interface {
	Handle(context.Context, Input) (Output, error)
}
