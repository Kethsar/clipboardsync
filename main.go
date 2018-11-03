package main

import (
	"fmt"
	pb "kethsar/clipboardsync/clipboard_proto"
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	cboard string
)

func main() {
	go startServer()
	monitorClipboard()
}

func syncClipoard(text string) {
	if text == cboard {
		return
	}
	cboard = text
	fmt.Printf("New clipboard text: %s\n\n", text)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		return
	}
	defer conn.Close()

	client := pb.NewClipboardSyncClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ret, err := client.SendClipboard(ctx, &pb.Clipboard{Data: text})
	fmt.Printf("Clipboard Copied: %t\n", ret.Success)
	if err != nil {
		log.Printf("Error sending clipboard: %v", err)
	}
}
