package service

import (
	"errors"
	"lendra/todo/features/user"
	"lendra/todo/utils/encrypts"
	"lendra/todo/utils/middlewares"
	"strings"
)

type userService struct {
	userData user.UserDataInterface
	h        encrypts.HashInterface
}

// dependency injection
func New(repo user.UserDataInterface, hash encrypts.HashInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
		h:        hash,
	}
}

// Create implements user.UserServiceInterface.
func (service *userService) Create(input user.Core) error {
	// logic validation
	if input.Email == "" {
		return errors.New("[validation] email harus diisi")
	}

	encrypPass, err := service.h.HashPassword(input.Password)
	if err != nil {
		return errors.New("terjadi masalah saat memproses data")
	}

	input.Password = encrypPass

	err = service.userData.Insert(input)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return errors.New("data yang dimasukkan sudah terdaftar")
		}
		return errors.New("terjadi kesalahan pada sistem")
	}

	return nil

}

// Update implements user.UserServiceInterface.
func (service *userService) Update(id uint, input user.Core) (user.Core, error) {
	//validasi
	updatedUser, err := service.userData.Update(id, input)
	if err != nil {
		return user.Core{}, err
	}

	return updatedUser, nil
}

// Login implements user.UserServiceInterface.
func (service *userService) Login(email string, password string) (data *user.Core, token string, err error) {
	result, err := service.userData.Login(email)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, "", errors.New("data tidak ditemukan")
		}
		return nil, "", errors.New("terjadi kesalahan pada sistem")
	}

	isValid := service.h.CheckPasswordHash(result.Password, password)
	if !isValid {
		return nil, "", errors.New("password yang diinputkan salah")
	}

	token, errJwt := middlewares.CreateToken(uint(result.ID))
	if errJwt != nil {
		return nil, "", errJwt
	}

	return result, token, nil
}

// Delete implements user.UserServiceInterface.
func (service *userService) Delete(id uint) error {
	err := service.userData.Delete(id)
	if err != nil {
		return err
	}
	return nil
}


// GetUser implements user.UserServiceInterface.
func (service *userService) GetUserById(id uint) (user.Core, error) {
	results, err := service.userData.SelectUserById(id)
	return results, err
}