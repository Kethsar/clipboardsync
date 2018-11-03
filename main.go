package main

import (
	"fmt"
	pb "kethsar/clipboardsync/clipboardsync"
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port    = ":9002"
	address = "192.168.1.10:9002"
)

var (
	client pb.ClipboardSyncClient
	server csServer
)

func main() {
	go startServer()
	monitorClipboard()
}

func syncClipoard(text string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		return
	}
	defer conn.Close()

	client = pb.NewClipboardSyncClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ret, err := client.SendClipboard(ctx, &pb.Clipboard{Data: "yes"})
	fmt.Printf("Clipboard Copied: %t\n", ret.Success)
	if err != nil {
		log.Printf("Error sending clipboard: %v", err)
	}
}
