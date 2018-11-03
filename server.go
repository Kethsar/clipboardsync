package main

import (
	"fmt"
	pb "kethsar/clipboardsync/clipboardsync"
	"log"
	"net"

	"github.com/atotto/clipboard"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type csServer struct{}

func (css *csServer) SendClipboard(ctx context.Context, in *pb.Clipboard) (*pb.Copied, error) {
	// Copy data to clipboard
	fmt.Printf("New Clipboard data: %s\n\n", in.Data)
	err := clipboard.WriteAll(in.Data)

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
