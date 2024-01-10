package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yafiesetyo/dating-apps-srv/internal/adapter/usecase"
	"github.com/yafiesetyo/dating-apps-srv/internal/constants"
	"github.com/yafiesetyo/dating-apps-srv/internal/entity"
	mock_interfaces "github.com/yafiesetyo/dating-apps-srv/internal/mocks"
	"go.uber.org/mock/gomock"
)

func TestSwipeProfile(t *testing.T) {
	type mockedDependencies struct {
		cacheMock         *mock_interfaces.MockCache
		userGetterMock    *mock_interfaces.MockRFindUserByID
		swipeCreatorMock  *mock_interfaces.MockRCreateUserSwipe
		unlikedUserGetter *mock_interfaces.MockRFindUserProfile
	}

	type testCase struct {
		name string
		args struct {
			id     uint
			swipe  uint
			action constants.SwipeAction
		}
		mockFunc func(md *mockedDependencies)
		wantErr  assert.ErrorAssertionFunc
	}

	testCases := []testCase{
		{
			name: "success",
			args: struct {
				id     uint
				swipe  uint
				action constants.SwipeAction
			}{
				id:     1,
				swipe:  2,
				action: constants.Like,
			},
			mockFunc: func(md *mockedDependencies) {
				md.userGetterMock.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return(entity.User{
						ID: 1,
					}, nil)

				md.userGetterMock.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return(entity.User{
						ID: 2,
					}, nil)

				md.unlikedUserGetter.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return([]entity.User{
						{
							ID: 2,
						},
					}, nil)

				md.cacheMock.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return("0", nil)

				md.swipeCreatorMock.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return(nil)

				md.cacheMock.EXPECT().
					Set(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "success unlimited swipe",
			args: struct {
				id     uint
				swipe  uint
				action constants.SwipeAction
			}{
				id:     1,
				swipe:  2,
				action: constants.Like,
			},
			mockFunc: func(md *mockedDependencies) {
				md.userGetterMock.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return(entity.User{
						ID:             1,
						UnlimitedSwipe: true,
					}, nil)

				md.userGetterMock.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return(entity.User{
						ID: 2,
					}, nil)

				md.unlikedUserGetter.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return([]entity.User{
						{
							ID: 2,
						},
					}, nil)

				md.swipeCreatorMock.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return(nil)

			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "error swiped user equal user",
			args: struct {
				id     uint
				swipe  uint
				action constants.SwipeAction
			}{
				id:     1,
				swipe:  1,
				action: constants.Like,
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
			},
		},
		{
			name: "error get user",
			args: struct {
				id     uint
				swipe  uint
				action constants.SwipeAction
			}{
				id:     1,
				swipe:  2,
				action: constants.Like,
			},
			mockFunc: func(md *mockedDependencies) {
				md.userGetterMock.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return(entity.User{}, assert.AnError)

			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
			},
		},
		{
			name: "error get liked user",
			args: struct {
				id     uint
				swipe  uint
				action constants.SwipeAction
			}{
				id:     1,
				swipe:  2,
				action: constants.Like,
			},
			mockFunc: func(md *mockedDependencies) {
				md.userGetterMock.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return(entity.User{
						ID: 1,
					}, nil)

				md.userGetterMock.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return(entity.User{}, assert.AnError)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
			},
		},
		{
			name: "error liked user not found",
			args: struct {
				id     uint
				swipe  uint
				action constants.SwipeAction
			}{
				id:     1,
				swipe:  2,
				action: constants.Like,
			},
			mockFunc: func(md *mockedDependencies) {
				md.userGetterMock.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return(entity.User{
						ID: 1,
					}, nil)

				md.userGetterMock.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return(entity.User{}, nil)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
			},
		},
		{
			name: "error unliked user not found",
			args: struct {
				id     uint
				swipe  uint
				action constants.SwipeAction
			}{
				id:     1,
				swipe:  2,
				action: constants.Like,
			},
			mockFunc: func(md *mockedDependencies) {
				md.userGetterMock.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return(entity.User{
						ID: 1,
					}, nil)

				md.userGetterMock.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return(entity.User{
						ID: 2,
					}, nil)

				md.unlikedUserGetter.
					EXPECT().Do(gomock.Any(), gomock.Any()).
					Return([]entity.User{}, assert.AnError)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
			},
		},
		{
			name: "error already swiped",
			args: struct {
				id     uint
				swipe  uint
				action constants.SwipeAction
			}{
				id:     1,
				swipe:  2,
				action: constants.Like,
			},
			mockFunc: func(md *mockedDependencies) {
				md.userGetterMock.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return(entity.User{
						ID: 1,
					}, nil)

				md.userGetterMock.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return(entity.User{
						ID: 2,
					}, nil)

				md.unlikedUserGetter.
					EXPECT().Do(gomock.Any(), gomock.Any()).
					Return([]entity.User{}, nil)
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			cacheMock := mock_interfaces.NewMockCache(ctrl)
			userGetterMock := mock_interfaces.NewMockRFindUserByID(ctrl)
			swipeCreatorMock := mock_interfaces.NewMockRCreateUserSwipe(ctrl)
			unlikedUserGetter := mock_interfaces.NewMockRFindUserProfile(ctrl)

			if tc.mockFunc != nil {
				tc.mockFunc(&mockedDependencies{
					cacheMock:         cacheMock,
					userGetterMock:    userGetterMock,
					swipeCreatorMock:  swipeCreatorMock,
					unlikedUserGetter: unlikedUserGetter,
				})
			}

			uc := usecase.NewSwipeProfile(
				cacheMock,
				userGetterMock,
				swipeCreatorMock,
				unlikedUserGetter,
			)

			err := uc.Do(context.Background(), tc.args.id, tc.args.swipe, tc.args.action)
			tc.wantErr(t, err)
		})
	}
}
