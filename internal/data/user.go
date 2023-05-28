package data

import (
	"context"
	"playground-kratos/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Save(ctx context.Context, user *biz.User) (*biz.User, error) {
	return user, nil
}

func (r *userRepo) Update(ctx context.Context, user *biz.User) (*biz.User, error) {
	return user, nil
}

func (r *userRepo) FindByID(ctx context.Context, id int64) (*biz.User, error) {
	row := r.data.db.QueryRowContext(ctx, `select id, name from users where id = $1`, id)
	if row.Err() != nil {
		return nil, row.Err()
	}
	user := &biz.User{}
	if err := row.Scan(&user.Id, &user.Name); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepo) ListAll(context.Context) ([]*biz.User, error) {
	return nil, nil
}
