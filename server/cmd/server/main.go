package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Epyklab/trident/server/logstream"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedLogServiceServer
}

func (s *server) LogMessage(ctx context.Context, in *pb.LogEntry) (*pb.LogResponse, error) {
	log.Printf("Received: %v", in.GetLogMessage())
	return &pb.LogResponse{Success: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:5555")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.StatsHandler(otelgrpc.NewServerHandler()))
	pb.RegisterLogServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
