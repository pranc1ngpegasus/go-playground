package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/pranc1ngpegasus/go-playground/domain/model"
	domain "github.com/pranc1ngpegasus/go-playground/domain/repository"
)

var _ domain.Repository[*model.User, *model.User] = (*FindByIDImpl)(nil)

var NewFindByIDImplSet = wire.NewSet(
	wire.Bind(new(domain.Repository[*model.User, *model.User]), new(*FindByIDImpl)),
	NewFindByIDImpl,
)

type FindByIDImpl struct{}

func NewFindByIDImpl() *FindByIDImpl {
	return &FindByIDImpl{}
}

func (r *FindByIDImpl) Handle(ctx context.Context, user *model.User) (*model.User, error) {
	return &model.User{}, nil
}
