package main

import (
	"github.com/micro/go-micro"

	"github.com/BurntSushi/toml"
	pb "github.com/Felyne/gomicrorpc/proto"
	"github.com/Felyne/launcher"
)

var (
	Version   = ""
	BuildTime = ""
)

func main() {
	serviceName := pb.ServiceName_name[0]
	launcher.Run(serviceName, Version, BuildTime, setup)
}

func setup(service micro.Service, _, cfgContent string) error {
	opt := Options{}
	if _, err := toml.Decode(cfgContent, &opt); err != nil {
		return err
	}
	h := NewSayService(opt)
	return pb.RegisterSayHandler(service.Server(), h)
}
