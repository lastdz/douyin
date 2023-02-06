package comment

import (
	"context"
	"github.com/RaymondCode/simple-demo/transport"
	"github.com/RaymondCode/simple-demo/wire"
	"github.com/smallnest/rpcx/client"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestActionHttp(t *testing.T) {
	ser := Commentserver{}
	ser.Start()
	time.Sleep(time.Hour)
}
func TestAction(t *testing.T) {
	ser := Commentserver{}
	ser.Start()
	args := make([]wire.ActionArgs, 1e3)
	for i := 0; i < 1e3; i++ {
		args[i] = wire.ActionArgs{"1", "1", "1", "你好哈哈", "0"}
	}
	replys := make([]wire.ActionReply, 1e3)
	tr := &transport.Transport{sync.Mutex{}, []string{"localhost:2379"}, make(map[string]*client.XClient, 0)}
	//插入
	for i := 0; i < 1e3; i++ {
		tr.Call(context.Background(), "comment", "Action", &args[i], &replys[i])
		if replys[i].Status_code != 0 {
			t.Error("wrong")
		}
	}
	//删除
	for i := 0; i < 1e3; i++ {
		args[i].Action_Type = "2"
		args[i].Comment_id = strconv.Itoa(int(replys[i].Comment.Id))
		tr.Call(context.Background(), "comment", "Action", &args[i], &replys[i])
		if replys[i].Status_code != 0 {
			t.Error("wrong")
		}
	}
}
func TestList(t *testing.T) {
	ser := Commentserver{}
	ser.Start()
	args := make([]wire.ActionArgs, 1e3)
	rd := rand.Int()
	rdstr := strconv.Itoa(rd)
	for i := 0; i < 1e3; i++ {
		args[i] = wire.ActionArgs{"1", rdstr, "1", "你好哈哈", "0"}
	}
	replys := make([]wire.ActionReply, 1e3)
	tr := &transport.Transport{sync.Mutex{}, []string{"localhost:2379"}, make(map[string]*client.XClient, 0)}
	//插入
	for i := 0; i < 1e3; i++ {
		tr.Call(context.Background(), "comment", "Action", &args[i], &replys[i])
		if replys[i].Status_code != 0 {
			t.Error("wrong")
		}
	}
	rep := wire.ListReply{}
	tr.Call(context.Background(), "comment", "List", &wire.ListArgs{"1", rdstr}, &rep)
	if len(rep.Comment_List) != 1e3 {
		t.Error("test error")
	}
}
