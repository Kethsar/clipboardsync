package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/awused/awconf"

	pb "github.com/Kethsar/clipboardsync/clipboard_proto"

	"google.golang.org/grpc"
)

type config struct {
	Port    string
	Server  string
	Timeout int
}

var (
	cboard string
	mux    sync.Mutex
	c      *config
)

func main() {
	err := awconf.LoadConfig("clipboardsync", &c)
	if err != nil {
		log.Fatalln(err)
	}

	go startServer()
	monitorClipboard()
}

// Send the clipboard to the server specified in the config if it is different
func syncClipoard(text string) {
	if !setClipboard(text) {
		return
	}

	conn, err := grpc.Dial(c.Server, grpc.WithInsecure())
	if err != nil {
		printToConsole(fmt.Sprintf("Failed to connect: %v", err))
		return
	}
	defer conn.Close()

	client := pb.NewClipboardSyncClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.Timeout)*time.Millisecond)
	defer cancel()

	_, err = client.Sync(ctx, &pb.Clipboard{Data: text})
	if err != nil {
		printToConsole(fmt.Sprintf("Error sending clipboard: %v", err))
		return
	}

	printToConsole("New clipboard sent")
}

// We have multiple threads accessing cboard, so use a mute when accessing it
func setClipboard(cb string) bool {
	mux.Lock()
	defer mux.Unlock()

	if cb == cboard {
		return false
	}

	cboard = cb
	return true
}

// Eh, I like formatted timestamps
func printToConsole(text string) {
	t := time.Now()
	fmt.Printf("[%d/%02d/%02d %02d:%02d:%02d] %s\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), text)
}
