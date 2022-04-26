package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"telegram-door-bell/internal/domain"
)

type UsersRepository interface {
	Create(user domain.User) error
	Read(ID string) (*domain.User, error)
}

type usersRepository struct {
	rc *redis.Client
}

func NewUsersRepository(rc *redis.Client) UsersRepository {
	return &usersRepository{rc: rc}
}

func (r *usersRepository) Create(user domain.User) error {
	return r.rc.Set(context.Background(), user.ID, user.Token, 0).Err()
}
func (r *usersRepository) Read(ID string) (*domain.User, error) {
	token, err := r.rc.Get(context.Background(), ID).Result()
	if err != nil {
		return nil, err
	}

	user := domain.User{
		ID:    ID,
		Token: token,
	}

	return &user, nil
}
