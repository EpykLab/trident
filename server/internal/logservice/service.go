package logservice

import (
	"context"

	pb "github.com/Epyklab/trident/server/pkg/protobuf"
)

type LogServiceServer struct {
	pb.UnimplementedLogServiceServer
}

func (s *LogServiceServer) LogMessage(ctx context.Context, in *pb.LogEntry) (*pb.LogResponse, error) {
	// Implement your logic to process the log message here
	return &pb.LogResponse{Success: true}, nil
}
