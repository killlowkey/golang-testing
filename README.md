## 前置准备：安装 mock
```shell
go install go.uber.org/mock/mockgen@latest
```

## 创建项目
```shell
go mod init github.com/killlowkey/golang-testing
go get go.uber.org/mock/gomock
go get github.com/stretchr/testify
go get github.com/gin-gonic/gin
```

## 创建 mock
```shell
mockgen -source ./store/user.go -destination ./store/user_mock.go -package store UserStore
mockgen -source ./biz/user.go -destination ./biz/user_mock.go -package biz UserService
```

## 运行测试
1. 运行所有测试：go test ./...
2. 运行单个测试：go test ./biz -run TestUserService_GetUserById
3. 运行指定包下测试：go test -v ./biz
4. 运行指定测试文件(需要指定依赖)：go test -v ./biz/user_test.go ./biz/user_mock.go ./biz/user.go
   > user_test.go 依赖了 user_mock.go 和 user.go，如果不指定依赖，会报错：undefined: UserServiceMock

## 参考资料
1. https://github.com/uber-go/mock
2. https://github.com/stretchr/testify
3. https://gin-gonic.com/docs/testing/