package grpcclients

import (
	proto "github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/Infinity-OJ/Server/internal/pkg/transports/grpc"
	"github.com/pkg/errors"
)

func NewSubmissionsClient(client *grpc.Client) (proto.SubmissionsClient, error) {
	conn, err := client.Dial("Submissions")
	if err != nil {
		return nil, errors.Wrap(err, "submission client dial error")
	}
	c := proto.NewSubmissionsClient(conn)

	return c, nil
}