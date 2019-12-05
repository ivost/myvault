package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	mysclient "github.com/ivost/myservice/pkg/client"
	v1 "github.com/ivost/shared/grpc/myvault"
	"github.com/ivost/shared/pkg/config"
	"github.com/ivost/shared/pkg/system"
	"github.com/ivost/shared/pkg/version"
)

func (s *Server) Health(ctx context.Context, none *empty.Empty) (resp *v1.HealthResponse, err error) {
	log.Printf("*** Health request received")
	resp = &v1.HealthResponse{
		Name:    "myvault",
		Version: version.Version,
		Build:   version.Build,
		Status:  "OK",
		Time:    time.Now().Format(time.RFC3339),
		Address: system.MyIP(),
	}

	conf := config.New("test/myservice-client-config.yaml")
	log.Printf("Calling MyService.Health, config: %+v", conf)
	c := mysclient.New(conf)
	rsp, err := c.Health()
	if err != nil {
		log.Printf("Error %v", err.Error())
		resp.Status = err.Error()
		return resp, err
	}
	msg := fmt.Sprintf("%+v", rsp)
	log.Printf("=== myservice status: %s", msg)
	resp.Status = "*** OK " + msg
	log.Printf("=== myvault resp: %+v", resp)
	return resp, err
}
