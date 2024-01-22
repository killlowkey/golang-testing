package store

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
)

func mockUser(id int) *User {
	return &User{
		Id:   id,
		Name: fmt.Sprintf("test%d", id),
		Age:  18,
	}
}

// TestMockUserStore_GetUserById https://github.com/uber-go/mock/blob/main/sample/user_test.go
func TestMockUserStore_GetUserById(t *testing.T) {
	ctl := gomock.NewController(t)

	mockUserStore := NewMockUserStore(ctl)
	// 期待 id 为 1
	mockUserStore.EXPECT().GetUserById(gomock.Eq(1)).Return(mockUser(1), nil)

	// 期待 id 为 2、3、6
	mockUserStore.EXPECT().GetUserById(gomock.AnyOf(2, 3, 6)).DoAndReturn(func(id int) (*User, error) {
		return mockUser(id), nil
	})

	// 期待 id 为 10 或以上
	mockUserStore.EXPECT().GetUserById(gomock.Cond(func(id any) bool {
		return id.(int) >= 10
	})).DoAndReturn(func(id int) (*User, error) {
		return nil, fmt.Errorf("id must less than 10")
	})

	user, err := mockUserStore.GetUserById(1)
	assert.NoError(t, err)
	assert.True(t, reflect.DeepEqual(mockUser(1), user))

	user, err = mockUserStore.GetUserById(3)
	assert.NoError(t, err)
	assert.True(t, reflect.DeepEqual(mockUser(3), user))

	user, err = mockUserStore.GetUserById(10)
	assert.Error(t, err)
	assert.Nil(t, user)
}

// TestMockUserStore_List
func TestMockUserStore_List(t *testing.T) {
	ctl := gomock.NewController(t)
	userStore := NewMockUserStore(ctl)
	res := []*User{mockUser(1), mockUser(2), mockUser(3)}

	userStore.EXPECT().List().Return(res, nil)

	users, err := userStore.List()
	assert.NoError(t, err)
	assert.True(t, reflect.DeepEqual(res, users))
}

// TestMockUserService_TDD_GetUserById
func TestMockUserService_TDD_GetUserById(t *testing.T) {
	ctl := gomock.NewController(t)
	userStore := NewMockUserStore(ctl)

	userStore.EXPECT().GetUserById(gomock.Eq(1)).Return(mockUser(1), nil)
	userStore.EXPECT().GetUserById(gomock.Cond(func(x any) bool {
		return x.(int) >= 10
	})).DoAndReturn(func(id int) (*User, error) {
		return nil, fmt.Errorf("id must less than 10")
	})

	testCases := []struct {
		name string
		arg  int
		want *User
		err  error
	}{
		{
			name: "id = 1",
			arg:  1,
			want: mockUser(1),
			err:  nil,
		},
		{
			name: "id = 10",
			arg:  10,
			want: nil,
			err:  fmt.Errorf("id must less than 10"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			user, err := userStore.GetUserById(tc.arg)
			assert.Equal(t, tc.err, err)
			assert.True(t, reflect.DeepEqual(tc.want, user))
		})
	}
}
