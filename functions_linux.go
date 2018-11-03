// +build linux

package main

import (
	"time"

	"github.com/atotto/clipboard"
)

const (
	address = "192.168.122.193:9002"
)

func monitorClipboard() {
	delay := make(<-chan time.Time)

	for {
		delay = time.After(500 * time.Millisecond)

		select {
		case <-delay:
			cb, err := clipboard.ReadAll()
			if err == nil {
				syncClipoard(cb)
			}
		}
	}
}
