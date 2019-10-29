package main

import (
	"log"

	"github.com/astaxie/beego/config"

	pb "github.com/Felyne/gomicrorpc/proto"
	"github.com/Felyne/launcher"
	"github.com/micro/go-micro/server"
)

var (
	Version   = ""
	BuildTime = ""
)

func main() {
	serviceName := pb.ServiceName_name[0]
	launcher.Run(serviceName, Version, BuildTime, setup)
}

func setup(s server.Server, cfgContent string) error {
	cfg, err := config.NewConfigData("ini", []byte(cfgContent))
	if err != nil {
		log.Printf("NewConfigData() failed: %v", err)
		return err
	}
	h := NewSayService(cfg)
	return pb.RegisterSayHandler(s, h)
}
