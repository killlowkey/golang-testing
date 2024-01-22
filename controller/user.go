package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/killlowkey/golang-testing/biz"
	"strconv"
)

type UserController struct {
	service biz.UserService
}

func NewUserController(service biz.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) GetUserById(ctx *gin.Context) {
	idStr := ctx.Param("id")

	// 转为int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  fmt.Sprintf("invalid id %s", idStr),
		})
		return
	}

	user, err := c.service.GetUserById(id)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  fmt.Sprintf("get user by id %d failed", id),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": user,
	})
}
