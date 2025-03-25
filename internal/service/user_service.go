package service

import (
	"github.com/Ryan-18-system/autenticaPB/internal/dto"
	"github.com/Ryan-18-system/autenticaPB/internal/entity"
	"github.com/Ryan-18-system/autenticaPB/internal/infra/database"
)

type UserServiceInterface interface {
	Login(username, password string) (string, error)
	CreateUser(userDto dto.UserImputDTO) error
	UpdateUser(userDto dto.UserImputDTO) error
	DeleteUser(userDto dto.UserImputDTO) error
	GetAllUsers() ([]dto.UserImputDTO, error)
}

type UserService struct {
	UserRepository database.UserRepositoryInterface
}

func NewUserService(repository database.UserRepositoryInterface) *UserService {
	return &UserService{UserRepository: repository}
}
func (u *UserService) Login(username, password string) (string, error) {
	return "", nil
}
func (u *UserService) CreateUser(userDto dto.UserImputDTO) error {
	newUser, err := entity.NewUser(userDto.Username, userDto.Password, userDto.Email, userDto.Nome, userDto.CPF)
	if err != nil {
		return err
	}
	err = u.UserRepository.Create(newUser)
	return err
}
func (u *UserService) UpdateUser(userDto dto.UserImputDTO) error {
	return nil
}
func (u *UserService) DeleteUser(userDto dto.UserImputDTO) error {
	return nil
}
func (u *UserService) GetAllUsers() ([]dto.UserImputDTO, error) {
	return nil, nil
}
