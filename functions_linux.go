// +build linux

package main

import (
	"fmt"
	"time"
)

func monitorClipboard() {
	delay := make(<-chan time.Time)

	for {
		delay = time.After(500 * time.Millisecond)

		select {
		case <-delay:
			fmt.Println("Not yet implemented")
		}
	}
}
