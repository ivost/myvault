package main

import (
	"log"

	"github.com/ivost/shared/pkg/version"
	"github.com/ivost/shared/pkg/config"

	"github.com/ivost/myservice/server"
)

func main() {
	log.Printf("%v client %v %v", version.Name, version.Version, version.Build)
	cf := config.DefaultConfigFile
	conf := config.New(cf)
	server := server.New(conf)
	if server != nil {
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}
}
