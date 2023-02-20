package comment

import (
	"context"
	"database/sql"
	"errors"

	"github.com/RaymondCode/simple-demo/server/message/db"

	"github.com/RaymondCode/simple-demo/transport"
	"github.com/RaymondCode/simple-demo/wire"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/server"
)

type MessageServer struct {
	tr *transport.Transport
}

func (ser *MessageServer) Start() {
	ser.tr = &transport.Transport{Etcdaddrs: []string{"localhost:2379"}, XClientMap: make(map[string]*client.XClient, 0)}
	s := server.NewServer()
	//添加etcd注册插件
	ser.tr.AddRegistryPlugin(s, "localhost:12307")
	//启动监听
	go s.Serve("tcp", "localhost:12307")
	//启动的监听名
	s.RegisterName("message", new(MessageServer), "")

}
func checkactionargs(args *wire.MessageActionArgs) string {
	if args.Token == "" || args.ActionType == "" || args.ToUserId == 0 {
		return "error:wrong args"
	}
	if args.ActionType != "1" {
		return "error:wrong ActionType"
	}
	return ""
}
func checkToken(token string) error {
	return nil
}
func (ser *MessageServer) Action(ctx context.Context, args *wire.MessageActionArgs, reply *wire.MessageActionReply) error {
	if errmsg := checkactionargs(args); errmsg != "" {
		reply.StatusCode = -1
		reply.StatusMessage = errmsg
		return errors.New("wrong args")
	}
	if err := checkToken(args.Token); err != nil {
		reply.StatusCode = -1
		reply.StatusMessage = "wrong token"
		return err
	}
	//token get id接口未实现
	//ser.tr.Call(context.Background(), "User", "Getuser", &getuserargs, &getuserreply)
	//uid := getuserreply.user.id
	uid := 1
	var toUserId = args.ToUserId
	message := db.MessageDb{
		Uid:      int64(uid),
		ToUserId: int64(toUserId),
		Content: sql.NullString{
			String: args.Content,
			Valid:  true,
		},
	}
	if err := db.InsertMessage(&message); err != nil {
		reply.StatusCode = -1
		reply.StatusMessage = "wrong message"
		return err
	}
	reply.StatusCode = 0
	reply.StatusMessage = "ok"
	reply.Message = wire.Message{
		Id:         message.Id,
		Content:    message.Content.String,
		CreateTime: message.CreateTime.String(),
	}
	return nil
}
func checkListArgs(args *wire.MessageListArgs) string {
	if args.UserId == 0 || args.Token == "" {
		return "error:wrong args"
	}
	return ""
}
func (ser *MessageServer) List(ctx context.Context, args *wire.MessageListArgs, reply *wire.MessageListReply) error {
	if errmsg := checkListArgs(args); errmsg != "" {
		reply.StatusCode = -1
		reply.StatusMessage = errmsg
		return errors.New("wrong args")
	}
	if err := checkToken(args.Token); err != nil {
		reply.StatusCode = -1
		reply.StatusMessage = "wrong token"
		return err
	}
	//token get id接口未实现
	//ser.tr.Call(context.Background(), "User", "Getuser", &getuserargs, &getuserreply)
	//uid := getuserreply.user.id
	uid := 1
	messageList, err := db.GetAllMessage(int64(uid), args.UserId)
	if err != nil {
		reply.StatusCode = -1
		reply.StatusMessage = "error"
		return err
	}
	var MessageList []wire.Message
	for _, message := range messageList {
		MessageList = append(MessageList, wire.Message{
			Id:         message.Id,
			Content:    message.Content.String,
			CreateTime: message.CreateTime.String(),
		})
	}
	reply.StatusCode = 0
	reply.StatusMessage = "ok"
	reply.MessageList = MessageList
	return nil
}
