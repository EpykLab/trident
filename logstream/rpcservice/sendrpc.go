package rpcservice

import "fmt"

// Args holds the arguments to pass to the logging function.
type Args struct {
	LogEntry string
}

// LogService provides RPC access to logging functionality.
type LogService struct{}

// LogMessage is an RPC method that clients can call to log a message.
func (t *LogService) LogMessage(args *Args, reply *bool) error {
	fmt.Println("Log Entry:", args.LogEntry)
	*reply = true // Indicate success
	return nil
}
