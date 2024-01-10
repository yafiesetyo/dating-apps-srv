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

func TestGetDatingProfileTest(t *testing.T) {
	type mockedDependencies struct {
		cacheMock      *mock_interfaces.MockCache
		userGetterMock *mock_interfaces.MockRFindUserByID
		profileGetter  *mock_interfaces.MockRFindUserProfile
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
				md.userGetterMock.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return(entity.User{
						ID: 1,
					}, nil)

				md.cacheMock.EXPECT().
					Keys(gomock.Any(), gomock.Any()).
					Return([]string{"view.1.12"}, nil)

				md.profileGetter.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return([]entity.User{
						{
							ID: 1,
						},
					}, nil)

				md.cacheMock.EXPECT().
					Set(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
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
			cacheMock := mock_interfaces.NewMockCache(ctrl)
			userGetterMock := mock_interfaces.NewMockRFindUserByID(ctrl)
			profileGetter := mock_interfaces.NewMockRFindUserProfile(ctrl)

			if tc.mockFunc != nil {
				tc.mockFunc(&mockedDependencies{
					cacheMock:      cacheMock,
					userGetterMock: userGetterMock,
					profileGetter:  profileGetter,
				})
			}

			uc := usecase.NewDatingProfileGetter(
				cacheMock, userGetterMock, profileGetter,
			)
			_, err := uc.Do(context.Background(), tc.args)
			tc.wantErr(t, err)
		})
	}
}
