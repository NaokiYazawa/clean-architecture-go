package repository

import (
	"github.com/NaokiYazawa/clean-architecture-go/domain/model"
)

// メソッド名(とコメント)で機能を「定義」だけする。
// データにアクセスするためのインターフェース
// Data Access Interface
// UserRepository Interface of an user repository
// UserRepository is interface for infrastructure
type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	ReadByID(id int) (*model.User, error)
	ReadAll() (*model.Users, error)
	Update(user *model.User) (*model.User, error)
	Delete(user *model.User) error
}
