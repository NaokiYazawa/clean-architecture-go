package usecase_test

import (
	"reflect"
	"testing"

	"github.com/NaokiYazawa/clean-architecture-go/domain/model"
	mock "github.com/NaokiYazawa/clean-architecture-go/domain/repository/mock"
	"github.com/NaokiYazawa/clean-architecture-go/usecase"

	"github.com/golang/mock/gomock"
)

func Test_userUsecase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type fields struct {
		userRepository *mock.MockUserRepository
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		prepare func(f *fields)
		want    *model.User
		wantErr bool
	}{
		{
			name: "【正常系】ユーザの新規登録",
			prepare: func(f *fields) {
				f.userRepository.EXPECT().Create(&model.User{ID: 0, Name: "sample"}).Return(&model.User{ID: 0, Name: "sample"}, nil).Times(1)
			},
			args:    args{name: "sample"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				userRepository: mock.NewMockUserRepository(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			userUsecase := usecase.NewUserUsecase(f.userRepository)
			got, err := userUsecase.Create(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, &model.User{ID: 0, Name: "sample"}) {
				t.Errorf("userUsecase.Create() = %v, want %v", got, &model.User{ID: 0, Name: "sample"})
			}
		})
	}
}
