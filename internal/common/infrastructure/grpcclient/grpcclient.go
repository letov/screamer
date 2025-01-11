package grpcclient

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "screamer/proto"
)

type GRPCClient struct {
	conn   *grpc.ClientConn
	Client pb.ScreamerServiceClient
}

func NewGRPCClient(
	lc fx.Lifecycle,
	log *zap.SugaredLogger,
) *GRPCClient {
	conn, err := grpc.NewClient(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			conn.Close()
			return nil
		},
	})

	client := pb.NewScreamerServiceClient(conn)
	return &GRPCClient{conn: conn, Client: client}
}
