package controller

import (
	"context"
	"fmt"
	"github.com/RaymondCode/simple-demo/transport"
	"github.com/RaymondCode/simple-demo/util"
	"github.com/RaymondCode/simple-demo/wire"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	hash := util.GetMd5String(username + password)

	args := &wire.RegisterArgs{
		Username: username,
		Password: hash,
	}
	reply := &wire.RegisterReply{}

	err := transport.RpcTp.Call(context.Background(), "user", "Register", args, reply)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.JSON(http.StatusOK, reply)
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	hash := util.GetMd5String(username + password)
	fmt.Println("Logging in")

	args := &wire.LoginArgs{
		Username: username,
		Password: hash,
	}
	reply := &wire.LoginReply{}

	err := transport.RpcTp.Call(context.Background(), "user", "Login", args, reply)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.JSON(http.StatusOK, reply)
}

func UserInfo(c *gin.Context) {
	uid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid user id")
	}
	token := c.Query("token")

	args := &wire.GetUserArgs{
		UserId: int64(uid),
		Token:  token,
	}
	reply := &wire.GetUserReply{}
	err = transport.RpcTp.Call(context.Background(), "user", "GetUser", args, reply)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.JSON(http.StatusOK, reply)
}
