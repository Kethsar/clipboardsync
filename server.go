package main

import (
	"io"
	"log"
	"net"
	"sync"

	pb "github.com/Kethsar/clipboardsync/clipboard_proto"

	"google.golang.org/grpc"
)

type csServer struct{}

var clients map[pb.ClipboardSync_SyncServer]bool
var smux sync.Mutex

func (css *csServer) Sync(stream pb.ClipboardSync_SyncServer) error {
	printToConsole("Server: Client connected")
	clients[stream] = true

	if cboard != "" { // Send the currently stored clipboard, if it exists
		stream.Send(&pb.Clipboard{Data: cboard})
	}

	// Loop and attempt to read new data from the client
	for {
		in, err := stream.Recv()
		if err != nil {
			printToConsole("Server: Client disconnected")

			// Remove the client from our list of clients, since we lost them
			smux.Lock()
			delete(clients, stream)
			smux.Unlock()

			if err != io.EOF {
				return err
			}

			return nil
		}

		// Send clipboard out only if it is different
		if setClipboard(in.GetData()) {
			printToConsole("Server: New clipboard received")
			propagate(cboard, stream)
		}
	}
}

// Send the specified text to all connected clients, except the one that sent it to the server
func propagate(text string, source pb.ClipboardSync_SyncServer) {
	smux.Lock()
	defer smux.Unlock()

	for stream := range clients {
		if stream != source {
			stream.Send(&pb.Clipboard{Data: text})
		}
	}
}

func startServer() {
	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	clients = make(map[pb.ClipboardSync_SyncServer]bool)

	s := grpc.NewServer()
	pb.RegisterClipboardSyncServer(s, &csServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while serving: %v\n", err)
	}
}
