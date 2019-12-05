package main

import (
	"log"

	"github.com/ivost/myvault/server"
	"github.com/ivost/shared/pkg/config"
	"github.com/ivost/shared/pkg/version"
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
