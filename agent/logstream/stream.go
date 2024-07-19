package logstream

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Epyklab/trident/agent/logstream/rpcstream"
)

func ShipLogs(lineChan chan string) {
	remote := os.Getenv("TRIDENT_SERVER_ADDRESS")

	conn, err := grpc.NewClient(
		remote,
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
