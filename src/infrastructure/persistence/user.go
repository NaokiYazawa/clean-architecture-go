// インフラ層
// データベースは永続化するため persistent という命名
package persistence

import (
	"github.com/NaokiYazawa/clean-architecture-go/domain/model"
	"github.com/NaokiYazawa/clean-architecture-go/domain/repository"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

// UserRepository struct of an user repository
type UserRepository struct {
	Conn *gorm.DB
}

// NewUserRepository Constructor of an user repository
func NewUserRepository(conn *gorm.DB) repository.UserRepository {
	return &UserRepository{Conn: conn}
}

// Create Create an user
func (userRepository *UserRepository) Create(User *model.User) (*model.User, error) {
	user := model.User{}
	copier.Copy(&user, &User)
	if err := userRepository.Conn.Create(&user).Error; err != nil {
		return nil, err
	}
	userModel := new(model.User)
	copier.Copy(&userModel, &user)

	return userModel, nil
}

// ReadByID Read an user by ID
func (userRepository *UserRepository) ReadByID(id int) (*model.User, error) {
	user := model.User{ID: id}
	if err := userRepository.Conn.First(&user).Error; err != nil {
		return nil, err
	}
	userModel := new(model.User)
	copier.Copy(&userModel, &user)

	return userModel, nil
}

// ReadAll Read users
func (userRepository *UserRepository) ReadAll() (*model.Users, error) {
	users := []model.User{}
	// gorm.Find from v2 doesn't return ErrRecordNotFound
	userRepository.Conn.Find(&users)
	userModels := new(model.Users)
	copier.Copy(&userModels, &users)

	return userModels, nil
}

// Update Update an user
func (userRepository *UserRepository) Update(User *model.User) (*model.User, error) {
	user := model.User{}
	copier.Copy(&user, &User)
	// Save vs Update
	// どちらの方が better なのか。
	if err := userRepository.Conn.Model(&user).Updates(&user).Error; err != nil {
		return nil, err
	}
	userModel := new(model.User)
	copier.Copy(&userModel, &user)

	return userModel, nil
}

// Delete Delete an user
func (userRepository *UserRepository) Delete(User *model.User) error {
	user := model.User{}
	copier.Copy(&user, &User)
	if err := userRepository.Conn.Delete(&user).Error; err != nil {
		return err
	}
	userModel := new(model.User)
	copier.Copy(&userModel, &user)

	return nil
}
