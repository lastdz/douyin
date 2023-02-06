package comment

import (
	"context"
	"errors"
	"github.com/RaymondCode/simple-demo/server/comment/db"
	"github.com/RaymondCode/simple-demo/transport"
	"github.com/RaymondCode/simple-demo/wire"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/server"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type Commentserver struct {
	tr *transport.Transport
}

func (ser *Commentserver) Start() {
	ser.tr = &transport.Transport{sync.Mutex{}, []string{"localhost:2379"}, make(map[string]*client.XClient, 0)}
	s := server.NewServer()
	//添加etcd注册插件
	ser.tr.AddRegistryPlugin(s, "localhost:12306")
	//启动监听
	go s.Serve("tcp", "localhost:12306")
	//启动的监听名
	s.RegisterName("comment", new(Commentserver), "")

}
func checkactionargs(args *wire.ActionArgs) string {
	if args.Token == "" || args.Action_Type == "" || args.Video_id == "" {
		return "error:wrong args"
	}
	if args.Action_Type != "1" && args.Action_Type != "2" {
		return "error:wrong action_type"
	}
	return ""
}
func checktoken(token string) error {
	return nil
}
func getcreatedate() string {
	return time.Now().Format("01-02")
}
func (ser *Commentserver) Action(ctx context.Context, args *wire.ActionArgs, reply *wire.ActionReply) error {

	if errmsg := checkactionargs(args); errmsg != "" {
		reply.Status_code = -1
		reply.Status_message = errmsg
		return errors.New("wrong args")
	}
	if err := checktoken(args.Token); err != nil {
		reply.Status_code = -1
		reply.Status_message = "wrong token"
		return err
	}
	if args.Action_Type == "1" {
		//user接口未实现
		//ser.tr.Call(context.Background(), "User", "Getuser", &getuserargs, &getuserreply)
		//uid := getuserreply.user.id
		var vid int
		var err error
		if vid, err = strconv.Atoi(args.Video_id); err != nil {
			reply.Status_code = -1
			reply.Status_message = "wrong commentid"
			return err
		}
		uid := rand.Int()
		comment := db.CommentDb{Uid: uid, Content: args.Comment_text, CreateDate: getcreatedate(), Vedioid: vid}
		if err := db.InsertComment(&comment); err != nil {
			reply.Status_code = -1
			reply.Status_message = "wrong comment"
			return err
		}

		//user:=getuserreply.user

		user := wire.User{int64(uid), "temp", 0, 0, false}
		com := wire.Comment{Id: comment.Id, User: user, Content: args.Comment_text, CreateDate: comment.CreateDate}
		reply.Status_code = 0
		reply.Comment = com
		reply.Status_message = "ok"
		return nil
	} else {
		var cid int
		var err error
		if cid, err = strconv.Atoi(args.Comment_id); err != nil {
			reply.Status_code = -1
			reply.Status_message = "wrong commentid"
			return err
		}
		if err := db.DeleteComment(int64(cid)); err != nil {
			reply.Status_code = -1
			reply.Status_message = "wrong commentid"
			return err
		}
		reply.Status_code = 0
		reply.Status_message = "ok"
		return nil
	}
	return nil
}
func checklistargs(args *wire.ListArgs) string {
	if args.Video_id == "" || args.Token == "" {
		return "error:wrong args"
	}
	return ""
}
func (ser *Commentserver) List(ctx context.Context, args *wire.ListArgs, reply *wire.ListReply) error {
	if errmsg := checklistargs(args); errmsg != "" {
		reply.Status_code = -1
		reply.Status_message = errmsg
		return errors.New("wrong args")
	}
	if err := checktoken(args.Token); err != nil {
		reply.Status_code = -1
		reply.Status_message = "wrong token"
		return err
	}
	reply.Status_code = 0
	reply.Status_message = "ok"
	var vid int
	var err error
	if vid, err = strconv.Atoi(args.Video_id); err != nil {
		reply.Status_code = -1
		reply.Status_message = "wrong token"
		return err
	}
	commentdb_list := db.GetAllComment(vid)
	comment_list := make([]wire.Comment, len(commentdb_list))
	for i := 0; i < len(comment_list); i++ {
		comment_list[i].Id = commentdb_list[i].Id
		comment_list[i].User = getuser(commentdb_list[i].Uid)
		comment_list[i].Content = commentdb_list[i].Content
		comment_list[i].CreateDate = commentdb_list[i].CreateDate
	}
	reply.Comment_List = comment_list
	return nil
}

//接口未实现
func getuser(uid int) wire.User {
	return wire.User{1, "1", 1, 1, true}
}
