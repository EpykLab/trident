package logstream

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Epyklab/trident/logstream/rpcstream"
)

func ShipLogs(lineChan chan string) {
	conn, err := grpc.NewClient(
		"host.docker.internal:5555",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewLogServiceClient(conn)

	for line := range lineChan {
		logEntry := &pb.LogEntry{LogMessage: line}
		response, err := client.LogMessage(context.Background(), logEntry)
		if err != nil {
			log.Printf("could not log message: %v", err)
			continue
		}

		fmt.Printf("Log entry successful: %v\n", response.Success)
	}
}
