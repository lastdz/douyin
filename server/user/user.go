package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/RaymondCode/simple-demo/server/user/db"
	"github.com/RaymondCode/simple-demo/transport"
	"github.com/RaymondCode/simple-demo/util"
	"github.com/go-sql-driver/mysql"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/server"
	"sync"
)

type UserServer struct {
	tr      *transport.Transport
	jwtServ *JwtService
}

type TestArgs struct {
}

type TestReply struct {
}

type RegisterArgs struct {
	Username string
	Password string
}

type RegisterReply struct {
	StatusCode int
	StatusMsg  string
	UserId     int64
	Token      string
}

type TokenValidateArgs struct {
	TokenStr string
	Username string
}

type TokenValidateReply struct {
	Success bool
	Msg     string
}

func (sv *UserServer) Start(dbAddr string, dbUser string, dbPasswd string, jwtSecret string) {
	sv.jwtServ = NewJwtService([]byte(jwtSecret))

	sv.tr = &transport.Transport{sync.Mutex{}, []string{"localhost:2379"}, make(map[string]*client.XClient, 0)}
	s := server.NewServer()
	sv.tr.AddRegistryPlugin(s, "localhost:12306")
	go s.Serve("tcp", "localhost:12306")
	s.RegisterName("user", new(UserServer), "")
	db.GetDb(dbAddr, dbUser, dbPasswd)
}

func (sv *UserServer) Test(ctx context.Context, args *TestArgs, reply *TestReply) error {
	fmt.Println("Test success")
	return nil
}

func (sv *UserServer) Register(ctx context.Context, args *RegisterArgs, reply *RegisterReply) error {
	hash := util.GetMd5String(args.Username + args.Password)
	id, err := db.InsertUser(args.Username, hash)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			reply.StatusCode = 1
			reply.StatusMsg = "Username already exists"
			return err
		} else {
			panic(err)
		}
	}
	reply.StatusCode = 0
	reply.StatusMsg = "Register successfully"
	reply.UserId = id
	reply.Token, _ = sv.jwtServ.GenerateToken(args.Username)
	return nil
}

func (sv *UserServer) ValidateToken(ctx context.Context, args *TokenValidateArgs, reply *TokenValidateReply) error {
	err := sv.jwtServ.ValidateToken(args.TokenStr, args.Username)
	if err != nil {
		reply.Success = false
		reply.Msg = err.Error()
		return nil
	}
	reply.Success = true
	return nil
}
