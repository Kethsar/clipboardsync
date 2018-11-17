package main

import (
	"log"
	"net"

	pb "github.com/Kethsar/clipboardsync/clipboard_proto"

	"github.com/atotto/clipboard"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type csServer struct{}

func (css *csServer) Sync(ctx context.Context, in *pb.Clipboard) (*pb.Copied, error) {
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
	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterClipboardSyncServer(s, &csServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while serving: %v\n", err)
	}
}
