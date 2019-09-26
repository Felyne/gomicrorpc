package impl

import (
	"context"

	"fmt"

	pb "github.com/Felyne/gomicrorpc/proto"
	"github.com/astaxie/beego/config"
)

type SayService struct{}

//一个简单的服务端实现
func NewSayService(cfg config.Configer) *SayService {
	conf := cfg.String("redis.addr") //读取配置，初始化服务
	fmt.Println("conf:", conf)
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
