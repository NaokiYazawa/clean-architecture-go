package usecase

import (
	"github.com/NaokiYazawa/clean-architecture-go/domain/model"
	"github.com/NaokiYazawa/clean-architecture-go/domain/repository"
)

// UserUsecase Interface of an user usecase
// Input Boundary
// どういった引数（入力データ）が必要なのかを定義する
type UserUsecase interface {
	Create(name string) (*model.User, error)
	ReadByID(id int) (*model.User, error)
	ReadAll() (*model.Users, error)
	Update(id int, name string) (*model.User, error)
	Delete(id int) error
}

// 外側から渡されたユーザデータを、外側にあるDBに保存する。
// ここで外側に直接アクセスしてしまうと依存関係が内→外に向いてしまう。これは避けたい。
// そこで、「依存関係逆転の原則」を使う。
// Go ではこれを interface を定義することで実現する。

// 同じ層に repository インターフェースを定義し、このインターフェースに依存するようにする。


type userUsecase struct {
	userRepository repository.UserRepository
}

// NewUserUsecase Constructor of an user usecase
// インターフェースを受け入れて、構造体を返す
func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{userRepository: userRepository}
}

// UseCase は、あくまでアプリケーションとして何が出来るのかのみを表現するだけである。
// そのため、実装は持たない。
// ビジネスロジック
// Create Usecase to save an user
func (userUsecase *userUsecase) Create(name string) (*model.User, error) {
	// validation
	user, err := model.NewUser(name)
	if err != nil {
		return nil, err
	}

	createdUser, err := userUsecase.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	// presenter に対して出力
	return createdUser, nil
}

// ReadByID Usecase to read an user by id
func (userUsecase *userUsecase) ReadByID(id int) (*model.User, error) {
	foundUser, err := userUsecase.userRepository.ReadByID(id)
	if err != nil {
		return nil, err
	}

	return foundUser, nil
}

// ReadAll Usecase to read users
func (userUsecase *userUsecase) ReadAll() (*model.Users, error) {
	foundUsers, err := userUsecase.userRepository.ReadAll()
	if err != nil {
		return nil, err
	}

	return foundUsers, nil
}

// Update Usecase to update an user
func (userUsecase *userUsecase) Update(id int, name string) (*model.User, error) {
	targetUser, err := userUsecase.userRepository.ReadByID(id)
	if err != nil {
		return nil, err
	}

	if err := targetUser.SetUser(name); err != nil {
		return nil, err
	}

	updatedUser, err := userUsecase.userRepository.Update(targetUser)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

// Delete Usecase to delete an user
func (userUsecase *userUsecase) Delete(id int) error {
	targetUser, err := userUsecase.userRepository.ReadByID(id)
	if err != nil {
		return err
	}

	err = userUsecase.userRepository.Delete(targetUser)
	if err != nil {
		return err
	}

	return nil
}
