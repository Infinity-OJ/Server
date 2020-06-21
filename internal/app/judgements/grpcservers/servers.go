package grpcservers

import (
	proto "github.com/infinity-oj/api/protobuf-spec"
	"github.com/infinity-oj/server/internal/pkg/transports/grpc"
	"github.com/google/wire"
	stdgrpc "google.golang.org/grpc"
)

func CreateInitServersFn(
	js *JudgementsService,
) grpc.InitServers {
	return func(s *stdgrpc.Server) {
		proto.RegisterJudgementsServer(s, js)
	}
}

var ProviderSet = wire.NewSet(NewJudgementsServer, CreateInitServersFn)
