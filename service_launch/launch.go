package service_launch

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/micro/go-micro"

	"github.com/Felyne/config_center"

	"github.com/coreos/etcd/clientv3"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/registry/etcdv3"
)

//每个服务端都要提供
type SetupFunc func(s server.Server, cfgContent string) error

// example: ./server dev 0 localhost:2379
func Start(serviceName, version, buildTime string, setup SetupFunc) {
	if len(os.Args) < 4 {
		if len(os.Args) == 2 && os.Args[1] == "-v" {
			fmt.Printf("version: %s\nbuildTime: %s\nserviceName: %s\n",
				version, buildTime, serviceName)
		} else {
			help()
		}
		os.Exit(1)
	}
	envName := os.Args[1]
	port := os.Args[2] //监听端口，0表示自动分配
	etcdAddrs := os.Args[3:]

	err := run(serviceName, version, envName, port, etcdAddrs, setup)
	if err != nil {
		log.Fatal(err)
	}
}

func run(serviceName, version, envName, port string, etcdAddrs []string, setup SetupFunc) error {
	etcdAddrs = withHttpAddr(etcdAddrs...)
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdAddrs,
		DialTimeout: 15 * time.Second,
	})
	if err != nil {
		return err
	}

	cc := config_center.NewConfigCenter(cli, envName)
	cfgContent, err := cc.GetConfig(serviceName)
	if err != nil {
		return err
	}
	cli.Close()

	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = etcdAddrs
	})

	options := []micro.Option{
		micro.Name(serviceName),
		micro.Registry(reg),
	}
	listenAddr := checkPort(port)
	if listenAddr != "" {
		options = append(options, micro.Address(listenAddr))
	}
	if version != "" {
		options = append(options, micro.Version(version))
	}

	service := micro.NewService(options...)
	service.Init()
	err = setup(service.Server(), cfgContent)
	if err != nil {
		return err
	}

	return service.Run()
}

func help() {
	info := `
Usage:%s [envName] [port] etcdAddr...
envName  env name
port     port for listen.if value is 0,listen on a random port
`
	fmt.Printf(info, os.Args[0])
}

func withHttpAddr(addrs ...string) []string {
	var list []string
	for _, a := range addrs {
		if false == strings.HasPrefix(a, "http://") {
			a = "http://" + a
		}
		list = append(list, a)
	}
	return list
}

func checkPort(port string) (addr string) {
	if port == "0" || port == ":0" {
		return
	}
	if false == strings.HasPrefix(port, ":") {
		addr = ":" + port
	}
	return
}
