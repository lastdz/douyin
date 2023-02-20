package comment

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/RaymondCode/simple-demo/transport"
	"github.com/RaymondCode/simple-demo/wire"
	"github.com/smallnest/rpcx/client"
)

func TestActionHttp(t *testing.T) {
	ser := MessageServer{}
	ser.Start()
	time.Sleep(time.Hour)
}
func TestAction(t *testing.T) {
	ser := MessageServer{}
	ser.Start()
	args := make([]wire.MessageActionArgs, 10)
	for i := 0; i < 10; i++ {
		args[i] = wire.MessageActionArgs{
			Token:      "1",
			ToUserId:   1,
			ActionType: "1",
			Content:    "1",
		}
	}
	replys := make([]wire.MessageActionReply, 10)
	tr := &transport.Transport{sync.Mutex{}, []string{"localhost:2379"}, make(map[string]*client.XClient, 0)}
	//插入
	for i := 0; i < 10; i++ {
		tr.Call(context.Background(), "message", "Action", &args[i], &replys[i])
		if replys[i].StatusCode != 0 {
			t.Error("wrong")
		}
	}
}
func TestList(t *testing.T) {
	ser := MessageServer{}
	ser.Start()
	tr := &transport.Transport{sync.Mutex{}, []string{"localhost:2379"}, make(map[string]*client.XClient, 0)}
	rep := wire.MessageListReply{}
	tr.Call(context.Background(), "message", "List", &wire.MessageListArgs{"1", 2}, &rep)
	fmt.Println(rep.MessageList)
	if rep.StatusCode != 0 {
		t.Error("test error")
	}
}
