package main

import (
	"fmt"
	pb "kethsar/clipboardsync/clipboard_proto"
	"log"
	"net"
	"time"

	"github.com/atotto/clipboard"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":9002"
)

type csServer struct{}

func (css *csServer) SendClipboard(ctx context.Context, in *pb.Clipboard) (*pb.Copied, error) {
	// Copy data to clipboard
	cboard = in.Data

	t := time.Now()
	fmt.Printf("[%d/%02d/%02d %02d:%02d:%02d] New clipboard received\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	err := clipboard.WriteAll(cboard)
	if err != nil {
		return &pb.Copied{Success: false}, err
	}

	return &pb.Copied{Success: true}, nil
}

func startServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterClipboardSyncServer(s, &csServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while serving: %v\n", err)
	}
}
