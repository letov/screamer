package grpcserver

import (
	"context"
	"net"
	"screamer/internal/common/metric"
	"screamer/internal/server/config"
	"screamer/internal/server/services"
	"strings"

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
	delta := in.GetDelta()
	value := float64(in.GetValue())

	jm := metric.JSONMetric{
		ID:    in.GetId(),
		MType: strings.ToLower(in.GetMtype().String()),
		Delta: &delta,
		Value: &value,
	}

	mtr, err := m.Ms.UpdateMetricJSON(ctx, jm)
	if err != nil {
		return nil, err
	}

	res := &pb.Response{
		Value: float32(mtr.Value),
		Ident: &pb.Ident{
			Type: mtr.Ident.Type.String(),
			Name: mtr.Ident.Name,
		},
	}

	return res, nil
}

func (m *GRPCServer) GetValue(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	delta := in.GetDelta()
	value := float64(in.GetValue())

	jm := metric.JSONMetric{
		ID:    in.GetId(),
		MType: strings.ToLower(in.GetMtype().String()),
		Delta: &delta,
		Value: &value,
	}

	mtr, err := m.Ms.ValueMetricJSON(ctx, jm)
	if err != nil {
		return nil, err
	}

	res := &pb.Response{
		Value: float32(mtr.Value),
		Ident: &pb.Ident{
			Type: mtr.Ident.Type.String(),
			Name: mtr.Ident.Name,
		},
	}

	return res, nil
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
