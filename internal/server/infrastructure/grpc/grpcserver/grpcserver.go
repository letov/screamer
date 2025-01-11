package grpcserver

import (
	"context"
	"net"
	"screamer/internal/common/application/dto"
	"screamer/internal/server/application/services"
	"screamer/internal/server/infrastructure/config"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	pb "screamer/proto"
)

type GRPCServer struct {
	pb.UnimplementedScreamerServiceServer
	Ms *services.MetricService
}

func (m *GRPCServer) UpdateValue(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	jm := dto.NewJsonMetricFromPb(in)
	r, err := m.Ms.UpdateMetricJSON(ctx, jm)
	if err != nil {
		return nil, err
	}
	dm, err := r.GetDomainMetric()
	if err != nil {
		return nil, err
	}
	return &pb.Response{
		Value: float32(dm.Value),
		Ident: &pb.Ident{
			Type: dm.Ident.Type.String(),
			Name: dm.Ident.Name,
		},
	}, nil
}

func (m *GRPCServer) GetValue(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	jm := dto.NewJsonMetricFromPb(in)
	mtr, err := m.Ms.ValueMetricJSON(ctx, jm)
	if err != nil {
		return nil, err
	}

	return &pb.Response{
		Value: float32(mtr.Value),
		Ident: &pb.Ident{
			Type: mtr.Ident.Type.String(),
			Name: mtr.Ident.Name,
		},
	}, nil
}

func NewGRPCServer(lc fx.Lifecycle, log *zap.SugaredLogger, c *config.Config, ms *services.MetricService) *grpc.Server {
	srv := grpc.NewServer()
	grpcServer := &GRPCServer{}
	grpcServer.Ms = ms
	pb.RegisterScreamerServiceServer(srv, grpcServer)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			addr := c.NetAddressGrpc.String()
			ln, err := net.Listen("tcp", addr)
			if err != nil {
				return err
			}
			log.Info("Starting GRPC server: ", addr)
			go func() {
				err := srv.Serve(ln)
				if err != nil {
					log.Error(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			srv.Stop()
			return nil
		},
	})

	return srv
}
