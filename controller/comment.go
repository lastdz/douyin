package controller

import (
	"context"
	"fmt"
	"github.com/RaymondCode/simple-demo/transport"
	"github.com/RaymondCode/simple-demo/wire"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {

	token := c.Query("token")
	actionType := c.Query("action_type")
	videoid := c.Query("video_id")
	fmt.Println(actionType, token, videoid, "111")
	if actionType == "1" {
		commenttext := c.Query("comment_text")
		args := wire.ActionArgs{token, videoid, actionType, commenttext, ""}
		reply := wire.ActionReply{}
		err := transport.RpcTp.Call(context.Background(), "comment", "Action", &args, &reply)
		if err != nil || reply.Status_code != 0 {
			c.JSON(http.StatusInternalServerError, reply.Status_message)
		} else {
			c.JSON(http.StatusOK, wire.CommentActionResponse{wire.Response{int32(reply.Status_code), reply.Status_message}, reply.Comment})
		}
	} else if actionType == "2" {
		commentid := c.Query("comment_id")
		args := wire.ActionArgs{token, videoid, actionType, "", commentid}
		reply := wire.ActionReply{}
		err := transport.RpcTp.Call(context.Background(), "comment", "Action", &args, &reply)
		if err != nil || reply.Status_code != 0 {
			c.JSON(http.StatusInternalServerError, reply.Status_message)
		} else {
			c.JSON(http.StatusOK, wire.CommentActionResponse{wire.Response{int32(reply.Status_code), reply.Status_message}, reply.Comment})
		}
	} else {
		c.JSON(http.StatusBadRequest, nil)
	}

}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	token := c.Query("token")
	videoid := c.Query("video_id")
	args := wire.ListArgs{token, videoid}
	reply := wire.ListReply{}
	err := transport.RpcTp.Call(context.Background(), "comment", "List", &args, &reply)
	if err != nil || reply.Status_code != 0 {
		c.JSON(http.StatusInternalServerError, reply.Status_message)
	}

	c.JSON(http.StatusOK, wire.CommentListResponse{wire.Response{int32(reply.Status_code), reply.Status_message}, reply.Comment_List})
}
