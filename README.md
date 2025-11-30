# Zing!

A tiny, user-friendly gRPC messaging command-line interface (CLI).

Talk to your friends from the terminal with zero fuss. Start the server, register or log in, send a message, and read your inbox — all in seconds.

## ✨ Features
- 📝 Register: create a new account and auto-save your login token
- 🔐 Login: authenticate and securely store your token in your config
- 📤 Send: write and send messages to other users (from flag, stdin, or your editor)
- 📥 Read: fetch your inbox; messages are delivered and cleared atomically
- ⏳ Session expiry: sessions automatically expire after 7 days; you’ll be prompted to log in again if needed
- ⚙️ Simple config: server settings and token are managed for you

## Build
- Using Go: `go build -o build/zing main.go`
- Using Make: `make zing` (outputs `build/zing`)
- Run tests: `make test`

## Serve (start the server)
- `./build/zing serve`
- Defaults: `localhost:50051`
- Change host/port: `./build/zing serve -a 0.0.0.0 -p 50051`

## Register
- `./build/zing register localhost:50051`
- You’ll be prompted for a username and password (with confirmation).
- On success, your token and connection settings are saved to your config.

## Login
- `./build/zing login localhost:50051`
- You’ll be prompted for username and password.
- On success, your token and connection settings are saved to your config.

## Send messages
- To a user with a flag: `./build/zing message send user2 -m "Hello!"`
- Or pipe stdin: `echo "Hi" | ./build/zing message send user2`
- If no `-m` and no stdin, your editor opens (uses `$VISUAL` or `$EDITOR`; falls back to `nano`).

## Read messages
- `./build/zing message read`
- Optional paging:
  - `--page-size 10`
  - `--page-token <token>`

## Clear messages
- `./build/zing message clear` — manually clear your inbox on the server

## Config
Zing stores settings and your auth token in a config file.
- Location (XDG conventions):
  - macOS/Linux: `~/.config/zing/config.toml`
  - Windows: `%AppData%/zing/config.toml`

Default keys:
- `server_addr = "localhost"`
- `server_port = 50051`
- `token = ""`  (filled after successful register/login)
- `plaintext = false`  (set with `-p`)
- `insecure = false`   (set with `-k`)

Notes:
- Register/Login store your token in the config.
- Commands use `server_addr:server_port` from the config.
- If your session is no longer valid, the server responds with: `your session has expired. please log in again`.
