package biz

import (
	"fmt"
	"github.com/killlowkey/golang-testing/store"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
)

func MockUser(id int) *store.User {
	return &store.User{
		Id:   id,
		Name: fmt.Sprintf("test%d", id),
		Age:  18,
	}
}

func TestUserService_GetUserById(t *testing.T) {
	ctl := gomock.NewController(t)

	mockUserStore := store.NewMockUserStore(ctl)
	mockUserStore.EXPECT().GetUserById(gomock.Eq(1)).Return(MockUser(1), nil)

	mockUserStore.EXPECT().GetUserById(gomock.AnyOf(2, 3, 6)).DoAndReturn(func(id int) (*store.User, error) {
		return MockUser(id), nil
	})

	userService := NewUserServiceImpl(mockUserStore)
	user, err := userService.GetUserById(1)
	assert.NoError(t, err)
	reflect.DeepEqual(MockUser(1), user)

	user, err = userService.GetUserById(3)
	assert.NoError(t, err)
	reflect.DeepEqual(MockUser(3), user)
}
