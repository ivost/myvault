package server

import (
	"context"
	"net"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	v1 "github.com/ivost/shared/grpc/myservice"
	"github.com/ivost/shared/pkg/version"
)

func (s *Server) Health(ctx context.Context, none *empty.Empty) (resp *v1.HealthResponse, err error) {
	resp = &v1.HealthResponse{
		Name:    "myservice",
		Version: version.Version,
		Build:   version.Build,
		Status:  "OK",
		Time:    time.Now().Format(time.RFC3339),
		Address: MyIP(),
	}
	return
}

func MyIP() string {
	ifaces, err := net.Interfaces()
	// handle err
	_ = err
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		// handle err
		_ = err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil {
				continue
			}
			s := ip.String()
			if strings.Contains(s, ":") {
				continue
			}
			if s == "127.0.0.1" {
				continue
			}
			//log.Printf("addr: %v", ip)
			return s
		}
	}
	return ""
}
