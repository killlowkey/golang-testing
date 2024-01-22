package biz

import (
	"errors"
	"github.com/killlowkey/golang-testing/store"
)

type UserService interface {
	GetUserById(id int) (*store.User, error)
	List() ([]*store.User, error)
	Create(user *store.User) error
}

type UserServiceImpl struct {
	store store.UserStore
}

func NewUserServiceImpl(store store.UserStore) UserService {
	return &UserServiceImpl{store: store}
}

func (s *UserServiceImpl) GetUserById(id int) (*store.User, error) {
	return s.store.GetUserById(id)
}

func (s *UserServiceImpl) List() ([]*store.User, error) {
	return s.store.List()
}

func (s *UserServiceImpl) Create(user *store.User) error {
	return errors.New("not implemented")
}
