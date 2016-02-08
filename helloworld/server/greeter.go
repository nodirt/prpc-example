package server

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/nodirt/prpc-example/helloworld/proto"
)

type greeterService struct{}

func (s *greeterService) SayHello(c context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	if req.Name == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "Name unspecified")
	}

	return &helloworld.HelloReply{
		Message: "Hello " + req.Name,
	}, nil
}
