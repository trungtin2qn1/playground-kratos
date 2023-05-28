package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id   int64
	Name string
}

// UserRepo is a Greater repo.
type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	FindByID(context.Context, int64) (*User, error)
	// ListByHello(context.Context, string) ([]*Greeter, error)
	ListAll(context.Context) ([]*User, error)
}

// GreeterUsecase is a Greeter usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *UserUsecase) CreateUser(ctx context.Context, user *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", user.Name)
	return uc.repo.Save(ctx, user)
}

// GetUser creates a GetUserById, and returns the new Greeter.
func (uc *UserUsecase) GetUserById(ctx context.Context, id int64) (*User, error) {
	uc.log.WithContext(ctx).Infof("GetUserById: %v", id)
	return uc.repo.FindByID(ctx, id)
}
