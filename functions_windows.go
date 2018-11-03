// +build windows

package main

import (
	"fmt"
	"syscall"
	"time"

	"github.com/atotto/clipboard"
)

const (
	address = "192.168.1.10:9002"
)

var (
	u32             = syscall.NewLazyDLL("user32.dll")
	getClipSequence = u32.NewProc("GetClipboardSequenceNumber")
	cbSeq           uintptr
)

func monitorClipboard() {
	delay := time.NewTicker(500 * time.Millisecond)

	for range delay.C {
		r1, _, _ := getClipSequence.Call()
		if r1 != cbSeq {
			cbSeq = r1

			cb, err := clipboard.ReadAll()
			if err != nil {
				fmt.Printf("Error reading clipboard: %v\n", err)
			} else {
				syncClipoard(cb)
			}
		}
	}
}
