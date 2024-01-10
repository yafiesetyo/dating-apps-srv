package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yafiesetyo/dating-apps-srv/internal/adapter/usecase"
	"github.com/yafiesetyo/dating-apps-srv/internal/entity"
	mock_interfaces "github.com/yafiesetyo/dating-apps-srv/internal/mocks"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestLogin(t *testing.T) {
	type mockedDependencies struct {
		userGetterMock *mock_interfaces.MockRFindUserByUsername
		passwordMock   *mock_interfaces.MockBCrypt
	}

	type testCase struct {
		name     string
		args     entity.User
		mockFunc func(md *mockedDependencies)
		wantErr  assert.ErrorAssertionFunc
	}

	testCases := []testCase{
		{
			name: "success",
			args: entity.User{
				Username: "test",
				Password: "12345678",
			},
			mockFunc: func(md *mockedDependencies) {
				md.userGetterMock.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return(entity.User{
						ID:       1,
						Password: "12345",
					}, nil)

				md.passwordMock.EXPECT().
					CompareAndHash(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(tt, err)
			},
		},
		{
			name: "error hash password",
			args: entity.User{
				Username: "test",
				Password: "12345678",
			},
			mockFunc: func(md *mockedDependencies) {
				md.userGetterMock.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return(entity.User{
						ID:       1,
						Password: "12345",
					}, nil)

				md.passwordMock.EXPECT().
					CompareAndHash(gomock.Any(), gomock.Any()).
					Return(assert.AnError)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(tt, err)
			},
		},
		{
			name: "error get user not found",
			args: entity.User{
				Username: "test",
				Password: "12345678",
			},
			mockFunc: func(md *mockedDependencies) {
				md.userGetterMock.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return(entity.User{}, gorm.ErrRecordNotFound)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(tt, err)
			},
		},
		{
			name: "error get user",
			args: entity.User{
				Username: "test",
				Password: "12345678",
			},
			mockFunc: func(md *mockedDependencies) {
				md.userGetterMock.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return(entity.User{}, assert.AnError)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(tt, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			userGetterMock := mock_interfaces.NewMockRFindUserByUsername(ctrl)
			passwordMock := mock_interfaces.NewMockBCrypt(ctrl)

			if tc.mockFunc != nil {
				tc.mockFunc(&mockedDependencies{
					userGetterMock: userGetterMock,
					passwordMock:   passwordMock,
				})
			}

			uc := usecase.NewLogin(userGetterMock, passwordMock)
			_, err := uc.Do(context.Background(), tc.args)
			tc.wantErr(t, err)
		})
	}
}
