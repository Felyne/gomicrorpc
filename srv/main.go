package main

import (
	"context"
	"fmt"

	model "github.com/Felyne/gomicrorpc/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
)

var etcdList = []string{"http://127.0.0.1:2379"}

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *model.SayParam, rsp *model.SayResponse) error {
	fmt.Println("received: ", req.Msg)
	rsp.Header = make(map[string]*model.Pair)
	rsp.Header["name"] = &model.Pair{Key: 1, Values: "abc"}
	rsp.Msg = "hello world"
	rsp.Values = append(rsp.Values, "a", "b")
	rsp.Type = model.RespType_DESCEND
	return nil
}

func main() {
	// 我这里用的etcd 做为服务发现，如果使用consul可以去掉
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = etcdList
	})

	service := micro.NewService(
		micro.Name(model.ServiceName_name[0]),
		micro.Registry(reg),
	)
	// 如果你用的是consul把上面的注释掉用下面的
	/*
		// 初始化服务
		service := micro.NewService(
			micro.Name("lp.srv.eg1"),
		)
	*/
	service.Init()
	model.RegisterSayHandler(service.Server(), new(Say))
	if err := service.Run(); err != nil {
		panic(err)
	}

}
