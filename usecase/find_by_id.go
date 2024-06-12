package usecase

import (
	"context"

	"github.com/google/wire"
	"github.com/pranc1ngpegasus/go-playground/domain/model"
	"github.com/pranc1ngpegasus/go-playground/domain/repository"
	domain "github.com/pranc1ngpegasus/go-playground/domain/usecase"
)

var _ domain.Usecase[model.User, *model.User] = (*FindByIDImpl)(nil)

var NewFindByIDImplSet = wire.NewSet(
	wire.Bind(new(domain.Usecase[model.User, *model.User]), new(*FindByIDImpl)),
	NewFindByIDImpl,
)

type FindByIDImpl struct {
	repository repository.Repository[*model.User, *model.User]
}

func NewFindByIDImpl(
	repository repository.Repository[*model.User, *model.User],
) *FindByIDImpl {
	return &FindByIDImpl{
		repository: repository,
	}
}

func (u *FindByIDImpl) Handle(ctx context.Context, user model.User) (*model.User, error) {
	return &model.User{}, nil
}
