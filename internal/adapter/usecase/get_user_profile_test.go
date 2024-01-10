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

func TestGetUserProfile(t *testing.T) {
	type mockedDependencies struct {
		getterMock *mock_interfaces.MockRFindUserByID
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
				ID: 1,
			},
			mockFunc: func(md *mockedDependencies) {
				md.getterMock.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return(
						entity.User{
							ID: 1,
						},
						nil,
					)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(tt, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			getterMock := mock_interfaces.NewMockRFindUserByID(ctrl)

			if tc.mockFunc != nil {
				tc.mockFunc(&mockedDependencies{
					getterMock: getterMock,
				})
			}

			uc := usecase.NewUserProfileGetter(getterMock)
			_, err := uc.Do(context.Background(), tc.args)

			tc.wantErr(t, err)
		})
	}
}
