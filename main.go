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

	t := time.Now()
	fmt.Printf("[%d/%02d/%02d %02d:%02d:%02d] New clipboard sent\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		return
	}
	defer conn.Close()

	client := pb.NewClipboardSyncClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = client.SendClipboard(ctx, &pb.Clipboard{Data: text})
	if err != nil {
		log.Printf("Error sending clipboard: %v", err)
	}
}
