package middleware

import (
	"github.com/ezio1119/fishapp-user/conf"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func (m *Middleware) RecoveryInterceptor() grpc.UnaryServerInterceptor {
	customFunc := func(p interface{}) (err error) {
		return grpc.Errorf(codes.Internal, "panic triggered: %v", p)
	}
	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []grpc_recovery.Option{}
	if conf.C.Sv.Debug {
		opts = append(opts, grpc_recovery.WithRecoveryHandler(customFunc))
	}

	return grpc_recovery.UnaryServerInterceptor(opts...)
}
