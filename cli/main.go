package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Felyne/launcher"

	"github.com/Felyne/gomicrorpc/client_factory"

	pb "github.com/Felyne/gomicrorpc/proto"
	"github.com/micro/go-micro/client"
)

func GetClient(envName string, etcdAddrs []string) pb.SayService {
	serviceName := launcher.GenServiceRegName(
		envName, pb.ServiceName_name[0])

	cli := client_factory.NewClient(etcdAddrs,
		func(c client.Client) interface{} {
			return pb.NewSayService(serviceName, c)
		})
	return cli.(pb.SayService)
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
