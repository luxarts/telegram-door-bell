package service

import (
	"strconv"
	"telegram-door-bell/internal/domain"
	"telegram-door-bell/internal/repository"
)

type UserService interface {
	Create(ID int64, token string) error
	Read(ID int64) (*domain.UserDTO, error)
}

type userService struct {
	userRepo repository.UsersRepository
}

func NewUserService(userRepo repository.UsersRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Create(ID int64, token string) error {
	dto := domain.UserDTO{
		ID:    ID,
		Token: token,
	}
	return s.userRepo.Create(dto.ToUser())
}
func (s *userService) Read(ID int64) (*domain.UserDTO, error) {
	user, err := s.userRepo.Read(strconv.FormatInt(ID, 10))
	if err != nil {
		return nil, err
	}

	dto := user.ToDTO()

	return &dto, nil
}
