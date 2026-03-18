package demo_microservice

import (
	"context"
)

type HealthServerImpl struct {
	UnimplementedHealthServer
}

func (h HealthServerImpl) Check(ctx context.Context, request *HealthCheckRequest) (*HealthCheckResponse, error) {
	return &HealthCheckResponse{
		Status: HealthCheckResponse_SERVING,
	}, nil
}

func (h HealthServerImpl) Watch(request *HealthCheckRequest, server Health_WatchServer) error {
	return server.Send(&HealthCheckResponse{
		Status: HealthCheckResponse_SERVING,
	})
}
