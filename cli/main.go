package main

import (
	"context"
	"fmt"

	"github.com/Felyne/gomicrorpc/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
)

var etcdList = []string{
	"http://127.0.0.1:2379",
}

func main() {
	// 我这里用的etcd 做为服务发现，如果使用consul可以去掉
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = etcdList
	})

	service := micro.NewService(
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
	sayClent := model.NewSayService(model.ServiceName_name[0], service.Client())
	rsp, err := sayClent.Hello(context.TODO(), &model.SayParam{Msg: "hello server, haha"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", rsp)
}
