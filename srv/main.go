package main

import (
	"log"

	"github.com/Felyne/gomicrorpc/impl"

	"github.com/astaxie/beego/config"

	pb "github.com/Felyne/gomicrorpc/proto"
	"github.com/Felyne/service_launch"
	"github.com/micro/go-micro/server"
)

var (
	Version   = ""
	BuildTime = ""
)

func main() {
	serviceName := pb.ServiceName_name[0]
	service_launch.Start(serviceName, Version, BuildTime, setup)
}

func setup(s server.Server, cfgContent string) error {
	cfg, err := config.NewConfigData("ini", []byte(cfgContent))
	if err != nil {
		log.Printf("NewConfigData() failed: %v", err)
		return err
	}
	h := impl.NewSayService(cfg)
	return pb.RegisterSayHandler(s, h)
}
