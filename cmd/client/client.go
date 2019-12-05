package main

import (
	"log"

	"github.com/ivost/myvault/pkg/client"
	"github.com/ivost/shared/pkg/config"
	"github.com/ivost/shared/pkg/version"
)

func main() {
	log.Printf("%v client %v %v", version.Name, version.Version, version.Build)
	cf := config.DefaultConfigFile
	conf := config.New(cf)
	c := client.New(conf)
	log.Printf("myvault grpc client endpoint:  %v, secure: %v", conf.GrpcAddr, conf.Secure)
	rsp, err := c.Health()
	if err != nil {
		log.Printf("error %v", err)
		return
	}

	log.Printf("myvault health response: %+v", rsp)
}
