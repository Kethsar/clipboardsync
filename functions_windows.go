// +build windows

package main

import (
	"fmt"
	"syscall"
	"time"

	"github.com/Kethsar/w32"

	"github.com/Kethsar/gform"

	"github.com/atotto/clipboard"
)

const (
	address = "192.168.1.10:9002"
)

var (
	cbWin *gform.Form // This is to make sure we keep it in memory, I guess
)

/*
	Attempt to listen for clipboard update events,
	else just poll the clipboard sequence number (Like MS says not to whoops)
*/
func monitorClipboard() {
	if createWindow() {
		gform.RunMainLoop()
	}

	printToConsole("Error when creating window for monitoring clipboard")
	printToConsole("Falling back to polling")

	pollClipboard()
}

// Create a basic window and register it to receive clipboard update events
func createWindow() bool {
	gform.Init()

	cbWin = gform.NewForm(nil)
	cbWin.Bind(w32.WM_CLIPBOARDUPDATE, cbUpdateHandler)

	return w32.AddClipboardFormatListener(cbWin.Handle())
}

func cbUpdateHandler(arg *gform.EventArg) {
	if data, ok := arg.Data().(*gform.RawMsg); ok {
		if data.Msg == w32.WM_CLIPBOARDUPDATE &&
			w32.IsClipboardFormatAvailable(w32.CF_UNICODETEXT) {
			cb, err := clipboard.ReadAll()

			if err != nil {
				printToConsole(fmt.Sprintf("Error reading clipboard: %v", err))
			} else {
				syncClipoard(cb)
			}
		}
	}
}

/*
	Poll and cache the clipboard sequence number
	This prevents constantly grabbing the clipboard itself
*/
func pollClipboard() {
	u32 := syscall.NewLazyDLL("user32.dll")
	getClipSequence := u32.NewProc("GetClipboardSequenceNumber")
	var cbSeq uintptr
	delay := time.NewTicker(500 * time.Millisecond)

	for range delay.C {
		r1, _, _ := getClipSequence.Call()
		if r1 != cbSeq {
			cbSeq = r1

			cb, err := clipboard.ReadAll()
			if err != nil {
				printToConsole(fmt.Sprintf("Error reading clipboard: %v", err))
			} else {
				syncClipoard(cb)
			}
		}
	}
}
