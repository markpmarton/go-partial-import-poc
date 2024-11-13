package server

import (
	"context"
	"log"
	"net"

	"gitlab.com/mark-p-marton/go-partial-import-poc/pkg/proto/sampleproto"
	"google.golang.org/grpc"
)

type SampleServer struct {
	sampleproto.UnimplementedSampleServiceServer
}

func (*SampleServer) ReceiveSampleEvent(ctx context.Context, sampleevent *sampleproto.SampleEvent) (*sampleproto.SampleEventReponse, error) {
	if len(sampleevent.Content) == 0 {
		return &sampleproto.SampleEventReponse{
			Status: false,
		}, nil
	}
	return &sampleproto.SampleEventReponse{
		Status: true,
	}, nil
}

func (*SampleServer) StartServer() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("TCP listener error: %v", err)
	}

	grpcServer := grpc.NewServer()
	service := &SampleServer{}
	sampleproto.RegisterSampleServiceServer(grpcServer, service)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("gRPC server error: %v", err)
	}
}
