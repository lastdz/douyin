package controller

import (
	"github.com/RaymondCode/simple-demo/server/relation"
	"github.com/RaymondCode/simple-demo/wire"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	to_user_id := c.Query("to_user_id")
	action_type := c.Query("action_type")
	args := wire.RelationActionArgs{
		Token:       token,
		To_user_id:  to_user_id,
		Action_type: action_type,
	}
	reply := wire.RelationActionReply{}
	err := relation.RelationAction(&args, &reply)
	if reply.StatusCode == 0 {
		c.JSON(http.StatusOK, Response{StatusCode: reply.StatusCode})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: reply.StatusCode, StatusMsg: err.Error()})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	user_id := c.Query("user_id")
	token := c.Query("token")
	args := wire.RelationListArgs{
		Token:   token,
		User_id: user_id,
	}
	reply := wire.RelationListReply{}
	relation.RelationFollowList(&args, &reply)
	c.JSON(http.StatusOK, reply)
	//c.JSON(http.StatusOK, UserListResponse{
	//	Response: Response{
	//		StatusCode: 0,
	//	},
	//	UserList: []User{DemoUser},
	//})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	user_id := c.Query("user_id")
	token := c.Query("token")
	args := wire.RelationListArgs{
		Token:   token,
		User_id: user_id,
	}
	reply := wire.RelationListReply{}
	relation.RelationFollowerList(&args, &reply)
	c.JSON(http.StatusOK, reply)
	//c.JSON(http.StatusOK, UserListResponse{
	//	Response: Response{
	//		StatusCode: 0,
	//	},
	//	UserList: []User{DemoUser},
	//})
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	user_id := c.Query("user_id")
	token := c.Query("token")
	args := wire.RelationListArgs{
		Token:   token,
		User_id: user_id,
	}
	reply := wire.RelationListReply{}
	relation.RelationFriendList(&args, &reply)
	c.JSON(http.StatusOK, reply)
	//c.JSON(http.StatusOK, UserListResponse{
	//	Response: Response{
	//		StatusCode: 0,
	//	},
	//	UserList: []User{DemoUser},
	//})
}
