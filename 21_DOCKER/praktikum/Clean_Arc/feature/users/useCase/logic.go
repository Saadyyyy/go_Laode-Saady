package usecase

import (
	users "Laode_Saady/21_DOCKER/praktikum/Clean_Arc/feature/users"
	"errors"
)

type userUseCase struct {
	userRepository users.DataInterface
}

// Login implements model.UserCaseInterface.
// Login implements model.UserCaseInterface.
func (us *userUseCase) Login(data users.Users) (users.Users, string, error) {
	user, token, err := us.userRepository.Login(data)
	if err != nil {
		return users.Users{}, "", err
	}
	return user, token, nil
}

// GetAll implements model.UserCaseInterface.
func (us *userUseCase) GetAll() ([]users.Users, error) {
	user, err := us.userRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetAllUsers implements model.UserCaseInterface.

func (us *userUseCase) Create(data users.Users) error {
	if data.Email == "" || data.Password == "" {
		return errors.New("error Email dan Password Harus di isi")
	}

	err := us.userRepository.Insert(data)
	return err
}

func New(userRepo users.DataInterface) users.UserCaseInterface {
	return &userUseCase{
		userRepository: userRepo,
	}
}
