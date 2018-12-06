# Clipboardsync
Syncs the clipboard between computers, physical or virtual.

There are three modes: client, server, and dual. Client attempts to connect to the server to send/receive clipboard data. Server runs to allow clients to connect and propagates the clipboard data to all connected clients. Dual runs both at the same time, and can be connected to by its own client, so you don't have to run multiple processes.

An attempt is made to listen to clipboard change notifications. If this fails, it falls back to polling. If you want a new feature, fork the repo and add it or open an issue. Feel free to send pull requests, or keep your changes to your own fork if you wish.

I don't actually know how to program haha.

## Platforms:
* Linux
* Windows

## Requirements
### Linux
 * xclip or xsel (required by [clipboard](https://github.com/atotto/clipboard))

### Windows
 * none, everything should Just Werk™

## Usage
Create a clipboardsync.toml file using the example as a base, and then run the program. Do this for the two machines you want to sync clipboards between.

The config file is searched for using [awconf](https://github.com/awused/awconf)

## TODO
- [x] One server, many clients. Server, Client, and dual mode
  - [x] Make the Sync RPC function use streams
  - [x] Keep connection open and send data back and forth
  - [x] Client auto-attempts reconnection every X seconds when disconnected
- Possibly encrypt the clipboard to make this safe for use if anything is not on the local network