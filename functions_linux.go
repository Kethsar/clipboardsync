// +build linux

package main

import (
	"fmt"
	"time"

	"github.com/BurntSushi/xgb/xfixes"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/atotto/clipboard"
)

const (
	cb = "CLIPBOARD"
)

var (
	cbatom xproto.Atom // This probably isn't actually needed
)

// Attempt to listen for selection events, else just poll the clipboard
func monitorClipboard() {
	xu, err := registerClipboardEvents()

	if err != nil {
		printToConsole(fmt.Sprintf("Error setting up event listening: %v", err))
		printToConsole("Using clipboard polling instead")
	} else {
		err = eventListening(xu)
		printToConsole(fmt.Sprintf("Error during event listening: %v", err))
		printToConsole("Switching to clipboard polling")
	}

	pollClipboard()
}

func eventListening(xu *xgbutil.XUtil) error {
	for {
		/*
			We don't actually care about the event returned
			since we only registered for selection "CLIPBOARD" owner changes
		*/
		_, xgberr := xu.Conn().WaitForEvent()

		// If we ever get an error, just stop, since I don't know of a good way to recover
		if xgberr != nil {
			return xgberr
		}

		cb, err := clipboard.ReadAll()
		if err == nil {
			syncClipoard(cb)
		}
	}
}

func pollClipboard() {
	delay := time.NewTicker(500 * time.Millisecond)

	for range delay.C {
		cb, err := clipboard.ReadAll()
		if err == nil {
			syncClipoard(cb)
		}
	}
}

// Set up "CLIPBOARD" selection owner change event notifications
func registerClipboardEvents() (*xgbutil.XUtil, error) {
	xu, err := xgbutil.NewConn()
	if err != nil {
		return nil, err
	}

	err = xfixes.Init(xu.Conn())
	if err != nil {
		return nil, err
	}

	cbinternatom, err := xproto.InternAtom(xu.Conn(), false, uint16(len(cb)), cb).Reply()
	if err != nil {
		return nil, err
	}
	cbatom = cbinternatom.Atom

	/*
		We don't actually care about what version is available, since any works
		But this is required before we can call SelectSelectionInput
	*/
	_, err = xfixes.QueryVersion(xu.Conn(), 5, 0).Reply()
	if err != nil {
		return nil, err
	}

	xfixes.SelectSelectionInput(xu.Conn(), xu.RootWin(), cbatom, xfixes.SelectionEventMaskSetSelectionOwner)

	return xu, nil
}
