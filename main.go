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
	Port          string
	Server        string
	Mode          int
	RetryInterval time.Duration
	MaxRetries    int
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
		go startClient()
		monitorClipboard()
	}
}

func startClient() {
	attempts := 0
	delaySecs := c.RetryInterval
	if delaySecs < 5 {
		delaySecs = 5
	}
	delay := time.NewTicker(delaySecs * time.Second)

	for {
		attempts++

		conn, err := grpc.Dial(c.Server, grpc.WithInsecure())
		if err != nil {
			printToConsole(fmt.Sprintf("Client: Failed to connect: %s", err))
			if attempts >= c.MaxRetries {
				break
			}

			printToConsole("Will retry in 30 seconds")
			<-delay.C
			continue
		}

		client := pb.NewClipboardSyncClient(conn)
		stream, err = client.Sync(context.Background())
		if err != nil {
			conn.Close()
			printToConsole(fmt.Sprintf("Client: Error creating stream: %s", err))
			if attempts >= c.MaxRetries {
				break
			}

			printToConsole("Will retry in 30 seconds")
			<-delay.C
			continue
		}

		attempts = 0

		printToConsole("Client: Stream opened")
		monitorClientStream()
		printToConsole("Client: Stream closed")

		// Set stream to nil so we don't try to use it somewhere else
		stream = nil
		conn.Close()
	}
	delay.Stop()
}

// Continously monitor the stream, break when receiving errors
func monitorClientStream() {
	for {
		in, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				printToConsole("Client: Reached end of stream")
			} else {
				printToConsole(fmt.Sprintf("Client: Failed to receive clipboard: %s", err))
			}

			break
		}

		if setClipboard(in.GetData()) {
			printToConsole("Client: New clipboard received")

			err = clipboard.WriteAll(cboard)
			if err != nil {
				printToConsole(fmt.Sprintf("Client: Failed to set clipboard: %s", err))
			}
		}
	}
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
