package transport

import (
	"context"
	"github.com/rcrowley/go-metrics"
	etcd_client "github.com/rpcxio/rpcx-etcd/client"
	"github.com/rpcxio/rpcx-etcd/serverplugin"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/server"
	"log"
	"sync"
	"time"
)

const (
	Basepath = "/rpcx"
)

type Transport struct {
	Mu         sync.Mutex
	Etcdaddrs  []string
	XClientMap map[string]*client.XClient
}

var RpcTp Transport

func init() {
	RpcTp = Transport{sync.Mutex{}, []string{"localhost:2379"}, make(map[string]*client.XClient, 0)}
}

//给RPC增加etcd注册中心插件
func (t *Transport) AddRegistryPlugin(s *server.Server, ipandport string) {

	r := &serverplugin.EtcdRegisterPlugin{
		ServiceAddress: "tcp@" + ipandport,         //服务器端口号
		EtcdServers:    []string{"localhost:2379"}, //etcd集群的所有ip
		BasePath:       "/rpcx",                    //固定为/rpcx
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}

//通过服务名获取XClient
func (t *Transport) getxclient(servicepath string) (*client.XClient, error) {
	t.Mu.Lock()
	defer t.Mu.Unlock()
	xclient, ok := t.XClientMap[servicepath]
	if !ok { //找不到创建一个 重复10次 间隔一秒
		cnt := 10
		for {
			if cnt < 0 {
				panic("not found service")
			}
			d, err := etcd_client.NewEtcdDiscovery(Basepath, servicepath, t.Etcdaddrs, false, nil)
			if err != nil {
				time.Sleep(time.Second)
				cnt--
				continue
			}
			tmp := client.NewXClient(servicepath, client.Failover, client.RoundRobin, d, client.DefaultOption)
			t.XClientMap[servicepath] = &tmp
			xclient = &tmp
			break
		}

	}
	return xclient, nil
}

//A.B() 可以通过t.Call(ctx,"A","B",args,reply)调用
func (t *Transport) Call(ctx context.Context, servicepath string, funcname string, args interface{}, reply interface{}) error {
	xclient, err := t.getxclient(servicepath)
	if err != nil {
		return err
	}
	return (*xclient).Call(ctx, funcname, args, reply)
}
