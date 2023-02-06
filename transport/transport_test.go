package transport

import (
	"context"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/server"
	"sync"
	"testing"
	"time"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

type Arith int

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A + args.B
	return nil
}
func TestCall(t *testing.T) {
	tr := &Transport{sync.Mutex{}, []string{"localhost:2379"}, make(map[string]*client.XClient, 0)}
	s := server.NewServer()
	//添加etcd注册插件
	tr.AddRegistryPlugin(s, "localhost:8972")
	//启动监听
	go s.Serve("tcp", "localhost:8972")
	//启动的监听名
	s.RegisterName("Arith", new(Arith), "")
	//测试
	tests := []struct {
		servicepath string
		funcname    string
		args        Args
		reply       Reply
		ans         Reply
	}{
		{
			"Arith",
			"Mul",
			Args{2, 3},
			Reply{0},
			Reply{5},
		},
		{
			"Arith",
			"Mul",
			Args{1, 3},
			Reply{0},
			Reply{4},
		},
		{
			"Arith",
			"Mul",
			Args{6, 3},
			Reply{0},
			Reply{9},
		},
	}

	for _, test := range tests {
		tr.Call(context.Background(), test.servicepath, test.funcname, &test.args, &test.reply)
		if test.reply.C != test.ans.C {
			t.Errorf("test error %d %d", test.reply.C, test.ans.C)
		}
	}
}

type calltest struct {
	servicepath string
	funcname    string
	args        Args
	reply       Reply
	ans         Reply
}

//并发测试
func TestCallconcurrently(t *testing.T) {
	tr := &Transport{sync.Mutex{}, []string{"localhost:2379"}, make(map[string]*client.XClient, 0)}
	s := server.NewServer()
	tr.AddRegistryPlugin(s, "localhost:8972")
	go s.Serve("tcp", "localhost:8972")
	//可以通过该函数注册服务
	s.RegisterName("Arith", new(Arith), "")
	tests := make([]calltest, 1e5)

	for i := 1; i <= 1e5; i++ {
		tests[i-1].servicepath = "Arith"
		tests[i-1].funcname = "Mul"
		tests[i-1].args = Args{i, i * 2}
		tests[i-1].reply = Reply{0}
		tests[i-1].ans = Reply{3 * i}
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 1e5; i++ {
		wg.Add(1)
		go func(i int) {
			tr.Call(context.Background(), tests[i].servicepath, tests[i].funcname, &tests[i].args, &tests[i].reply)
			wg.Done()
		}(i)
	}
	wg.Wait()
	for _, test := range tests {
		if test.reply.C != test.ans.C {
			t.Errorf("test error %d %d", test.reply.C, test.ans.C)
		}
	}
}

//测试新上线的服务能否被感知  连续测试会报错因为etcd中的key还没过期
func TestDe_registername(t *testing.T) {
	tr := &Transport{sync.Mutex{}, []string{"localhost:2379"}, make(map[string]*client.XClient, 0)}
	s := server.NewServer()
	tr.AddRegistryPlugin(s, "localhost:8972")
	go s.Serve("tcp", "localhost:8972")
	go func() {
		//等7s再注册
		time.Sleep(7 * time.Second)
		s.RegisterName("Arith", new(Arith), "")

	}()
	tests := make([]calltest, 1e3)

	for i := 1; i <= 1e5; i++ {
		tests[i-1].servicepath = "Arith"
		tests[i-1].funcname = "Mul"
		tests[i-1].args = Args{i, i * 2}
		tests[i-1].reply = Reply{0}
		tests[i-1].ans = Reply{3 * i}
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 1e5; i++ {
		tr.Call(context.Background(), tests[i].servicepath, tests[i].funcname, &tests[i].args, &tests[i].reply)
		//t.Errorf("done")
	}
	wg.Wait()
	for _, test := range tests {
		if test.reply.C != test.ans.C {
			t.Errorf("test error %d %d", test.reply.C, test.ans.C)
		}
	}
}
