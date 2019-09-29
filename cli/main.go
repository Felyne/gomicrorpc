package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Felyne/gomicrorpc/client_factory"

	pb "github.com/Felyne/gomicrorpc/proto"
	"github.com/micro/go-micro/client"
)

func GetClient(etcdAddrs []string) pb.SayService {
	cli := client_factory.NewClient(etcdAddrs,
		func(c client.Client) interface{} {
			return pb.NewSayService(pb.ServiceName_name[0], c)
		})
	return cli.(pb.SayService)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf(`Usage:%s [etcdAddr...]\n`, os.Args[0])
		os.Exit(1)
	}
	etcdAddrs := os.Args[1:]
	cli := GetClient(etcdAddrs)
	rsp, err := cli.Hello(context.TODO(), &pb.SayParam{Msg: "hello server"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", rsp)
}
