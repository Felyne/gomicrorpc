package client_factory

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
)

type NewClientFunc func(c client.Client) interface{}

func NewClient(etcdAddrs []string, newFunc NewClientFunc) interface{} {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = etcdAddrs
	})
	service := micro.NewService(
		micro.Registry(reg),
	)
	service.Init()
	return newFunc(service.Client())
}
