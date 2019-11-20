package main

import (
	"context"
	"fmt"
	"os"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"

	"github.com/micro/go-micro"

	"github.com/Felyne/launcher"

	pb "github.com/Felyne/gomicrorpc/proto"
)

func GetClient(envName string, etcdAddrs []string) pb.SayService {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = etcdAddrs
	})
	service := micro.NewService(
		micro.Registry(reg),
	)
	service.Init()
	serviceName := launcher.GenServiceRegName(
		envName, pb.ServiceName_name[0])
	return pb.NewSayService(serviceName, service.Client())
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf(`Usage:%s [envName] [etcdAddr...]\n`, os.Args[0])
		os.Exit(1)
	}
	envName := os.Args[1]
	etcdAddrs := os.Args[2:]
	cli := GetClient(envName, etcdAddrs)
	rsp, err := cli.Hello(context.TODO(), &pb.SayParam{Msg: "hello server"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", rsp)
}
