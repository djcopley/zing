# Zing!

A tiny, user-friendly gRPC messaging command-line interface (CLI).

What you can do:
- run the local server,
- login as a demo user,
- send messages,
- read your inbox.

## Build
- Using Go: go build -o build/zing main.go
- Using Make: make zing (outputs build/zing)
- Run tests: make test

## Serve (start the server)
- ./build/zing serve
- Defaults: localhost:5132
- Change host/port: ./build/zing serve -a 0.0.0.0 -p 50051

## Login
- ./build/zing login localhost:5132
- You’ll be prompted for username and password.
- Demo users (built-in):
  - user1 / pass
  - user2 / pass

## Send messages
- To user2 with a flag: ./build/zing message send user2 -m "Hello!"
- Or pipe stdin: echo "Hi" | ./build/zing message send user2
- If no -m and no stdin, your editor opens (uses $VISUAL or $EDITOR; falls back to nano).

## Read messages
- ./build/zing message read
- Optional paging:
  - --page-size 10
  - --page-token <token>

## Config
Zing stores settings and your auth token in a config file.
- Location (XDG conventions):
  - macOS/Linux: ~/.config/zing/config.toml
  - Windows: %AppData%/zing/config.toml

Default keys:
- server_addr = "localhost"
- server_port = 5132
- token = ""  (filled after successful login)

Notes:
- Login stores your token in the config.
- Commands use server_addr:server_port from the config.
