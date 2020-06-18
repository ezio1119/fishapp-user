package middleware

import (
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
)

func (*Middleware) ValidatorInterceptor() grpc.UnaryServerInterceptor {
	return grpc_validator.UnaryServerInterceptor()
}
