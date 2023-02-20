package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/RaymondCode/simple-demo/wire"

	"github.com/RaymondCode/simple-demo/transport"
	"github.com/gin-gonic/gin"
)

type ChatResponse struct {
	Response
	MessageList []Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	content := c.Query("content")
	actionType := c.Query("action_type")
	var toUserId int
	var err error
	if toUserId, err = strconv.Atoi(c.Query("to_user_id")); err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "ToUser doesn't exist"})
		return
	}
	//将消息插入数据库
	args := wire.MessageActionArgs{
		Token:      token,
		ToUserId:   toUserId,
		ActionType: actionType,
		Content:    content,
	}
	var reply wire.MessageActionReply
	err = transport.RpcTp.Call(context.Background(), "message", "Action", &args, &reply)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "error"})
		return
	}
	c.JSON(http.StatusOK, Response{StatusCode: 0})
}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
	token := c.Query("token")
	var toUserId int
	var err error
	if toUserId, err = strconv.Atoi(c.Query("to_user_id")); err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "ToUser doesn't exist"})
		return
	}

	args := wire.MessageListArgs{
		Token:  token,
		UserId: int64(toUserId),
	}
	reply := wire.MessageListReply{}
	//查询数据库
	err = transport.RpcTp.Call(context.Background(), "message", "List", &args, &reply)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "ErrorRpcCall"})
		return
	}
	var MessageList []Message
	for _, message := range reply.MessageList {
		MessageList = append(MessageList, Message{
			Id:         message.Id,
			Content:    message.Content,
			CreateTime: message.CreateTime,
		})
	}
	c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0}, MessageList: MessageList})
}
