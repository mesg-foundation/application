package core

import (
	"context"
)

// ListServices lists services.
func (s *Server) ListServices(ctx context.Context, request *ListServicesRequest) (*ListServicesReply, error) {
	services, err := s.api.ListServices()
	return &ListServicesReply{Services: toGRPCServices(services)}, err
}
