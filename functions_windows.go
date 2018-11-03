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
	u32                            = syscall.NewLazyDLL("user32.dll")
	procGetClipboardSequenceNumber = u32.NewProc("GetClipboardSequenceNumber")
	cbSeq                          uintptr
)

func monitorClipboard() {
	delay := make(<-chan time.Time)

	for {
		delay = time.After(500 * time.Millisecond)

		select {
		case <-delay:
			r1, _, _ := procGetClipboardSequenceNumber.Call()
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
}
