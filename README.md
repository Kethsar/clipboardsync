# Clipboardsync
Syncs the clipboard between two machines, physical or virtual.

An attempt is made to listen to clipboard change notifications. If this fails, it falls back to polling.

I don't expect this to be useful to anyone else. If you want a new feature, fork the repo and add it or open an issue. Feel free to send pull requests, or keep your changes to your own fork if you wish.

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
- One server, many clients. Server, Client, and dual mode
  - Make the Sync RPC function use streams
  - Keep connection open and send data back and forth
  - Client auto-attempts reconnection ever X seconds when disconnected
- Possibly encrypt the clipboard to make this safe for use if anything is not on the local network