package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/killlowkey/golang-testing/biz"
	"github.com/killlowkey/golang-testing/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// UserControllerTestSuite https://gin-gonic.com/docs/testing/
type UserControllerTestSuite struct {
	suite.Suite
	controller *UserController
	router     *gin.Engine
}

// SetupSuite 只运行一次
func (suite *UserControllerTestSuite) SetupSuite() {
	ctl := gomock.NewController(suite.T())
	userService := biz.NewMockUserService(ctl)
	userService.EXPECT().GetUserById(1).Return(&store.User{
		Id:   1,
		Name: "test1",
		Age:  18,
	}, nil)

	userService.EXPECT().GetUserById(gomock.Cond(func(id any) bool {
		return id.(int) >= 2
	})).Return(nil, fmt.Errorf("user not found"))

	suite.controller = NewUserController(userService)
	suite.router = gin.Default()
	suite.router.GET("/api/v1/user/:id", suite.controller.GetUserById)
}

func (suite *UserControllerTestSuite) TearDownSuite() {
	suite.router = nil
	suite.controller = nil
}

func (suite *UserControllerTestSuite) TestUserController_GetUserById() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/user/1", nil)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
	assert.True(suite.T(), reflect.DeepEqual(`{"code":200,"data":{"id":1,"name":"test1","age":18},"msg":"success"}`, w.Body.String()))
	fmt.Println(w.Body.String())
}

func (suite *UserControllerTestSuite) TestUserController_GetUserById_InvalidId() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/user/abc", nil)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
	assert.True(suite.T(), reflect.DeepEqual(`{"code":400,"msg":"invalid id abc"}`, w.Body.String()))
	fmt.Println(w.Body.String())
}

func (suite *UserControllerTestSuite) TestUserController_GetUserById_NotFound() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/user/2", nil)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
	assert.True(suite.T(), reflect.DeepEqual(`{"code":400,"msg":"get user by id 2 failed"}`, w.Body.String()))
	fmt.Println(w.Body.String())
}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
