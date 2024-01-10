package data

import (
	"errors"
	"lendra/todo/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}

// Insert implements user.UserDataInterface.
func (repo *userQuery) Insert(input user.Core) error {
	// proses mapping dari struct entities core ke model gorm
	userInputGorm := User{
		Name:        input.Name,
		Email:       input.Email,
		Password:    input.Password,
		Address:     input.Address,
		PhoneNumber: input.PhoneNumber,
		Role:        input.Role,
	}
	// simpan ke DB
	tx := repo.db.Create(&userInputGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}

	return nil
}

// Update implements user.UserDataInterface.
func (repo *userQuery) Update(id uint, input user.Core) (user.Core, error) {
	dataGorm := CoreToModel(input)
	user := User{}
	tx := repo.db.First(&user, id)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return input, errors.New("User id not found")
		}
		return input, tx.Error
	}

	updateTx := repo.db.Model(&User{}).Where("id = ?", id).Updates(dataGorm)
	if updateTx.Error != nil {
		// fmt.Println("err:", updateTx.Error)
		return input, updateTx.Error
	}
	return user.ModelToCore(), nil
}

// Login implements user.UserDataInterface.
func (repo *userQuery) Login(email string) (data *user.Core, err error) {
	var userGorm User
	tx := repo.db.Where("email = ?", email).First(&userGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}
	result := userGorm.ModelToCore()
	return &result, nil
}

// Delete implements user.UserDataInterface.
func (repo *userQuery) Delete(id uint) error {
	user := User{}
	tx := repo.db.First(&user, id)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return errors.New("User id not found")
		}
		return tx.Error
	}

	deleteTx := repo.db.Delete(&user)
	if deleteTx.Error != nil {
		return deleteTx.Error
	}
	return nil
}

// SelectUser implements user.UserDataInterface.
func (repo *userQuery) SelectUserById(id uint) (user.Core, error) {
	userModel := &User{}
    err := repo.db.Where("id = ?", id).First(userModel).Error
    if err != nil {
        return user.Core{}, err
    }

    // Proses mapping dari struct gorm model ke struct core
    userCore := user.Core{
        ID:          userModel.ID,
        Name:        userModel.Name,
        Email:       userModel.Email,
        Password:    userModel.Password,
        Address:     userModel.Address,
        PhoneNumber: userModel.PhoneNumber,
        Role:        userModel.Role,
        CreatedAt:   userModel.CreatedAt,
        UpdatedAt:   userModel.UpdatedAt,
    }

    return userCore, nil
}
