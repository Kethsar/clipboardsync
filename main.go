package main

import (
	"fmt"
	pb "kethsar/clipboardsync/clipboard_proto"
	"sync"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	cboard string
	mux    sync.Mutex
)

func main() {
	go startServer()
	monitorClipboard()
}

func syncClipoard(text string) {
	if !setClipboard(text) {
		return
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		printToConsole(fmt.Sprintf("Failed to connect: %v", err))
		return
	}
	defer conn.Close()

	client := pb.NewClipboardSyncClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = client.Sync(ctx, &pb.Clipboard{Data: text})
	if err != nil {
		printToConsole(fmt.Sprintf("Error sending clipboard: %v", err))
		return
	}

	printToConsole("New clipboard sent")
}

func setClipboard(cb string) bool {
	mux.Lock()
	defer mux.Unlock()

	if cb == cboard {
		return false
	}

	cboard = cb
	return true
}

func printToConsole(text string) {
	t := time.Now()
	fmt.Printf("[%d/%02d/%02d %02d:%02d:%02d] %s\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), text)
}
