package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yafiesetyo/dating-apps-srv/internal/adapter/usecase"
	"github.com/yafiesetyo/dating-apps-srv/internal/entity"
	mock_interfaces "github.com/yafiesetyo/dating-apps-srv/internal/mocks"
	"go.uber.org/mock/gomock"
)

func TestRegister(t *testing.T) {
	type mockedDependencies struct {
		userCheckerMock *mock_interfaces.MockRFindUserByUsername
		userCreatorMock *mock_interfaces.MockRCreateUser
		passwordMock    *mock_interfaces.MockBCrypt
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
				Name: "a",
			},
			mockFunc: func(md *mockedDependencies) {
				md.userCheckerMock.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return(entity.User{}, nil)

				md.passwordMock.EXPECT().
					GenerateFromPassword(gomock.Any(), gomock.Any()).
					Return([]byte{}, nil)

				md.userCreatorMock.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			userCheckerMock := mock_interfaces.NewMockRFindUserByUsername(ctrl)
			userCreatorMock := mock_interfaces.NewMockRCreateUser(ctrl)
			passwordMock := mock_interfaces.NewMockBCrypt(ctrl)

			if tc.mockFunc != nil {
				tc.mockFunc(&mockedDependencies{
					userCheckerMock: userCheckerMock,
					userCreatorMock: userCreatorMock,
					passwordMock:    passwordMock,
				})
			}

			uc := usecase.NewRegister(
				userCheckerMock,
				userCreatorMock,
				passwordMock,
			)

			err := uc.Do(context.Background(), tc.args)
			tc.wantErr(t, err)
		})
	}
}
