// インフラ層
// データベースは永続化するため persistent という命名
// Data Access
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
// userRepository に依存
func (userRepository *UserRepository) Create(User *model.User) (*model.User, error) {
	user := model.User{}
	// copier.Copy(&user, &User)
	if err := copier.Copy(&user, &User); err != nil {
		return nil, err
	}
	if err := userRepository.Conn.Create(&user).Error; err != nil {
		return nil, err
	}
	userModel := new(model.User)
	// copier.Copy(&userModel, &user)
	if err := copier.Copy(&userModel, &user); err != nil {
		return nil, err
	}

	return userModel, nil
}

// ReadByID Read an user by ID
func (userRepository *UserRepository) ReadByID(id int) (*model.User, error) {
	user := model.User{ID: id}
	if err := userRepository.Conn.First(&user).Error; err != nil {
		return nil, err
	}
	userModel := new(model.User)
	// copier.Copy(&userModel, &user)
	if err := copier.Copy(&userModel, &user); err != nil {
		return nil, err
	}

	return userModel, nil
}

// ReadAll Read users
func (userRepository *UserRepository) ReadAll() (*model.Users, error) {
	users := []model.User{}
	// gorm.Find from v2 doesn't return ErrRecordNotFound
	userRepository.Conn.Find(&users)
	userModels := new(model.Users)
	// copier.Copy(&userModels, &users)
	if err := copier.Copy(&userModels, &users); err != nil {
		return nil, err
	}

	return userModels, nil
}

// Update Update an user
func (userRepository *UserRepository) Update(User *model.User) (*model.User, error) {
	user := model.User{}
	// copier.Copy(&user, &User)
	if err := copier.Copy(&user, &User); err != nil {
		return nil, err
	}
	// Save vs Update
	// どちらの方が better なのか。
	if err := userRepository.Conn.Model(&user).Updates(&user).Error; err != nil {
		return nil, err
	}
	userModel := new(model.User)
	// copier.Copy(&userModel, &user)
	if err := copier.Copy(&userModel, &user); err != nil {
		return nil, err
	}

	return userModel, nil
}

// Delete Delete an user
func (userRepository *UserRepository) Delete(User *model.User) error {
	user := model.User{}
	// copier.Copy(&user, &User)
	if err := copier.Copy(&user, &User); err != nil {
		return err
	}
	if err := userRepository.Conn.Delete(&user).Error; err != nil {
		return err
	}
	userModel := new(model.User)
	// copier.Copy(&userModel, &user)
	if err := copier.Copy(&userModel, &user); err != nil {
		return err
	}

	return nil
}
