package main

import (
	pb "kethsar/clipboardsync/clipboard_proto"
	"log"
	"net"

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
	if !setClipboard(in.Data) {
		return &pb.Copied{Success: false}, nil
	}

	printToConsole("New clipboard received")

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
