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

func TestPurchase(t *testing.T) {
	type mockedDependencies struct {
		userGetterMock           *mock_interfaces.MockRFindUserByID
		premiumFeatureSetterMock *mock_interfaces.MockRSetPremiumFeature
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
				ID:             1,
				UnlimitedSwipe: true,
			},
			mockFunc: func(md *mockedDependencies) {
				md.userGetterMock.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return(
						entity.User{
							ID: 1,
						},
						nil,
					)

				md.premiumFeatureSetterMock.EXPECT().
					Do(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(tt, err)
			},
		},
		{
			name: "success",
			args: entity.User{
				ID:           1,
				VerifiedUser: true,
			},
			mockFunc: func(md *mockedDependencies) {
				md.userGetterMock.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return(
						entity.User{
							ID: 1,
						},
						nil,
					)

				md.premiumFeatureSetterMock.EXPECT().
					Do(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(tt, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			userGetterMock := mock_interfaces.NewMockRFindUserByID(ctrl)
			premiumFeatureSetterMock := mock_interfaces.NewMockRSetPremiumFeature(ctrl)

			if tc.mockFunc != nil {
				tc.mockFunc(&mockedDependencies{
					userGetterMock:           userGetterMock,
					premiumFeatureSetterMock: premiumFeatureSetterMock,
				})
			}

			uc := usecase.NewPurchase(userGetterMock, premiumFeatureSetterMock)
			err := uc.Do(context.Background(), tc.args)

			tc.wantErr(t, err)
		})
	}
}
