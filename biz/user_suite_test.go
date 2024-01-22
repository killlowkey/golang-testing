package biz

import (
	"fmt"
	"github.com/killlowkey/golang-testing/store"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"testing"
)

// UserServiceTestSuite https://github.com/stretchr/testify
type UserServiceTestSuite struct {
	suite.Suite
	service UserService
}

// SetupSuite 只运行一次
func (suite *UserServiceTestSuite) SetupSuite() {
	fmt.Println("SetupSuite")
	ctl := gomock.NewController(suite.T())

	mockUserStore := store.NewMockUserStore(ctl)
	mockUserStore.EXPECT().GetUserById(gomock.Eq(1)).Return(MockUser(1), nil)
	mockUserStore.EXPECT().List().DoAndReturn(func() ([]*store.User, error) {
		return []*store.User{MockUser(1), MockUser(2), MockUser(3)}, nil
	})

	suite.service = NewUserServiceImpl(mockUserStore)
}

func (suite *UserServiceTestSuite) TearDownSuite() {
	fmt.Println("TearDownSuite")
	suite.service = nil
}

func (suite *UserServiceTestSuite) SetupTest() {
	fmt.Println("SetupTest")
}

func (suite *UserServiceTestSuite) TearDownTest() {
	fmt.Println("TearDownTest")
}

func (suite *UserServiceTestSuite) BeforeTest(suiteName, testName string) {
	fmt.Printf("BeforeTest suiteName: %s, testName: %s\n", suiteName, testName)
}

func (suite *UserServiceTestSuite) AfterTest(suiteName, testName string) {
	fmt.Printf("AfterTest suiteName: %s, testName: %s\n", suiteName, testName)
}

func (suite *UserServiceTestSuite) TestGetUserById() {
	user, err := suite.service.GetUserById(1)
	suite.NoError(err)
	suite.Equal(MockUser(1), user)
}

func (suite *UserServiceTestSuite) TestList() {
	users, err := suite.service.List()
	suite.NoError(err)
	suite.Equal([]*store.User{MockUser(1), MockUser(2), MockUser(3)}, users)
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
