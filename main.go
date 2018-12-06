package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"github.com/atotto/clipboard"
	"github.com/awused/awconf"

	pb "github.com/Kethsar/clipboardsync/clipboard_proto"

	"google.golang.org/grpc"
)

type config struct {
	Port   string
	Server string
	Mode   int
}

const (
	clientMode = 1
	serverMode = 2
	dualMode   = 3
)

var (
	cboard string
	mux    sync.Mutex
	c      *config
	stream pb.ClipboardSync_SyncClient
)

func main() {
	err := awconf.LoadConfig("clipboardsync", &c)
	if err != nil {
		log.Fatalln(err)
	}

	if c.Mode == serverMode {
		startServer()
	}

	if c.Mode == dualMode {
		go startServer()
	}

	if c.Mode == clientMode || c.Mode == dualMode {
		go createAndMonitorStream()
		monitorClipboard()
	}
}

func createAndMonitorStream() {
	conn, err := grpc.Dial(c.Server, grpc.WithInsecure())
	if err != nil {
		printToConsole(fmt.Sprintf("Client: Failed to connect: %s", err))
		return
	}
	defer conn.Close()

	client := pb.NewClipboardSyncClient(conn)
	stream, err = client.Sync(context.Background())
	if err != nil {
		printToConsole(fmt.Sprintf("Client: Error creating stream: %s", err))
		return
	}

	printToConsole("Client: Stream opened")

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			printToConsole("Client: Reached end of stream")
			break
		}

		if err != nil {
			printToConsole(fmt.Sprintf("Client: Failed to receive clipboard: %s", err))
		} else if setClipboard(in.GetData()) {
			printToConsole("Client: New clipboard received")

			err = clipboard.WriteAll(cboard)
			if err != nil {
				printToConsole(fmt.Sprintf("Client: Failed to set clipboard: %s", err))
			}
		}
	}

	// Set stream to nil so we know we need to re-create it
	stream = nil
}

// Send the clipboard to the server specified in the config if it is different
func syncClipoard(text string) {
	if !setClipboard(text) {
		return
	}

	if stream == nil {
		printToConsole("Client: No connection to a server is open, unable to send clipboard")
		return
	}

	err := stream.Send(&pb.Clipboard{Data: text})
	if err != nil {
		printToConsole(fmt.Sprintf("Client: Error sending clipboard: %s", err))
		return
	}

	printToConsole("Client: New clipboard sent")
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
