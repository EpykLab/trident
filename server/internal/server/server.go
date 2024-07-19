package server

import (
	"github.com/Epyklab/trident/server/internal/logservice"
	"google.golang.org/grpc"

	pb "github.com/Epyklab/trident/server/pkg/protobuf"
)

func RegisterServices(s *grpc.Server) {
	logService := logservice.LogServiceServer{}
	pb.RegisterLogServiceServer(s, &logService)
}
