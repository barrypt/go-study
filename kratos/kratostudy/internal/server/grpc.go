package server

import (
	v1 "kratostudy/api/helloworld/v1"
	"kratostudy/internal/conf"
	"kratostudy/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}

	// create grpc conn
// only for demo, use single instance in production env
     conn, err := grpc.DialInsecure(ctx,
	grpc.WithEndpoint("127.0.0.1:9000"),
	grpc.WithMiddleware(
		recovery.Recovery(),
		tracing.Client(),
	),
	grpc.WithTimeout(2*time.Second),
	// for tracing remote ip recording
	grpc.WithOptions(grpcx.WithStatsHandler(&tracing.ClientHandler{})),
)
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterGreeterServer(srv, greeter)
	return srv
}
