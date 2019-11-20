package main

import (
	"context"

	"fmt"

	pb "github.com/Felyne/gomicrorpc/proto"
)

// 仅做测试用
type Options struct {
	RedisAddr string
	DB        int
	Password  string
}

type SayService struct{}

//一个简单的服务端实现
func NewSayService(opt Options) *SayService {
	fmt.Printf("%+v\n", opt)
	return &SayService{}
}

func (s *SayService) Hello(ctx context.Context, req *pb.SayParam, rsp *pb.SayResponse) error {
	fmt.Println("received: ", req.Msg)
	rsp.Header = make(map[string]*pb.Pair)
	rsp.Header["name"] = &pb.Pair{Key: 1, Values: "abc"}
	rsp.Msg = "hello world"
	rsp.Values = append(rsp.Values, "a", "b")
	rsp.Type = pb.RespType_DESCEND
	return nil
}
