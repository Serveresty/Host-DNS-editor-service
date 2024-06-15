package internal

import (
	"YADROhostsDNS/pkg/api"
	"context"
	"fmt"
)

type Server struct {
	api.UnimplementedHostnameServiceServer
}

func (s *Server) SetHostname(_ context.Context, req *api.SetHostnameRequest) (*api.SetHostnameResponce, error) {
	hostname := req.GetHostname()
	err := SetHostname(hostname)
	if err != nil {
		return &api.SetHostnameResponce{Message: fmt.Sprintf("Failed to write to /etc/hostname: %v", err)}, err
	}
	return &api.SetHostnameResponce{Message: "Set hostname success"}, nil
}

func (s *Server) ListDNSServers(_ context.Context, req *api.ListDNSServersRequest) (*api.ListDNSServersResponce, error) {
	list, err := GetDNSList()
	if err != nil {
		return &api.ListDNSServersResponce{}, err
	}
	return &api.ListDNSServersResponce{Servers: list}, nil
}

func (s *Server) AddDNSServer(_ context.Context, req *api.AddDNSServerRequest) (*api.AddDNSServerResponce, error) {
	newServer := req.GetServer()
	if err := AddDNSServer(newServer); err != nil {
		return &api.AddDNSServerResponce{Message: fmt.Sprintf("Failed to add DNS: %v", err)}, err
	}
	return &api.AddDNSServerResponce{Message: "DNS successfully added"}, nil
}

func (s *Server) RemoveDNSServer(_ context.Context, req *api.RemoveDNSServerRequest) (*api.RemoveDNSServerResponce, error) {
	server := req.GetServer()
	if err := RemoveDNCServer(server); err != nil {
		return &api.RemoveDNSServerResponce{Message: fmt.Sprintf("Failed to remove DNS: %v", err)}, err
	}
	return &api.RemoveDNSServerResponce{Message: "DNS successfully removed"}, nil
}
