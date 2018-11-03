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
	delay := time.NewTicker(500 * time.Millisecond)

	for range delay.C {
		cb, err := clipboard.ReadAll()
		if err == nil {
			syncClipoard(cb)
		}
	}
}
