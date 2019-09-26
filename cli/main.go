package main

import (
	"context"
	"fmt"

	"github.com/Felyne/gomicrorpc/client_factory"

	pb "github.com/Felyne/gomicrorpc/proto"
	"github.com/micro/go-micro/client"
)

var etcdList = []string{
	"http://127.0.0.1:2379",
}

func GetClient(etcdAddrs []string) pb.SayService {
	cli := client_factory.NewClient(etcdAddrs,
		func(c client.Client) interface{} {
			return pb.NewSayService(pb.ServiceName_name[0], c)
		})
	return cli.(pb.SayService)
}

func main() {
	cli := GetClient(etcdList)
	rsp, err := cli.Hello(context.TODO(), &pb.SayParam{Msg: "hello server"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", rsp)
}
