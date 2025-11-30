# Zing

Lightweight, self‑hosted messaging for the command line. Zing provides a tiny gRPC server backed by Redis and a CLI
that lets you register, log in, send, and read messages directly from your terminal. It is ideal for lightweight 
team messaging, quick notifications from scripts/CI, and communicating with users on headless machines.

## Why Zing
- CLI‑first workflow: send, read, and clear messages without leaving the terminal.
- Scriptable: send messages from shell scripts, CI jobs, or other automation by piping stdin.
- Self‑hosted and minimal: a single binary server with Redis; deploy locally, in Docker, or to Kubernetes.
- Secure by default: password hashing on the server and short‑lived sessions (1 week) stored in Redis.

## Features
- Register and login via CLI; a session token is stored in your config.
- Send messages via flag, piped stdin, or your editor (`$VISUAL`/`$EDITOR`, fallback to `nano`).
- Read messages with colorized, paged output and simple pagination options.
- Clear your inbox on demand.
- Configuration via file or environment variables; TLS options for clients.

## Installation

Install the CLI (and server) with Go:

```
go install github.com/djcopley/zing@latest
```

This installs the `zing` binary into your `$GOPATH/bin` or `$GOBIN` (ensure it’s on your `PATH`). Requires Go 1.24+.

You can also build from source:
- Using Go: `go build -o build/zing ./main.go`
- Using Make: `make zing` (outputs `build/zing`)
- Run tests: `make test`

## Quick start

1) Start Redis (required by the server)

```
docker run -d --name redis -p 6379:6379 redis:7
```

2) Start the Zing server (point it at Redis)

```
REDIS_ADDR=localhost:6379 zing serve
```

By default the server binds to `localhost:50051`. You can change the bind address/port:

```
zing serve -a 0.0.0.0 --port 50051
```

Configure the Redis connection via config or environment variables (see Configuration). For example:

```
REDIS_ADDR=localhost:6379 zing serve
```

3) Register a user and store the token

The CLI connects with TLS by default. The built‑in server runs without TLS, so for local development pass `--plaintext` once during register/login; this setting is saved to your config for future commands.

```
zing --plaintext register localhost:50051 -u alice
```

You will be prompted to set and confirm a password. On success, your token and connection settings are saved.

4) Send and read messages

Send a message using a flag:

```
zing message send bob -m "Hello from Zing!"
```

Or pipe content from stdin:

```
echo "Build finished ✅" | zing message send bob
```

Read your inbox (paged and colorized by default):

```
zing message read
```

If your session expires, you’ll see: `your session has expired. please log in again`.

## CLI usage

- Start server
  - `zing serve` (binds to `localhost:50051`)
  - Change bind address/port: `zing serve -a 0.0.0.0 --port 50051`

- Register/login (first‑time connection settings are saved)
  - `zing --plaintext register localhost:50051 -u alice`
  - `zing --plaintext login localhost:50051 -u alice`
  - TLS options (client‑side):
    - `--plaintext, -p` Use a non‑TLS connection (useful for local dev)
    - `--insecure, -k`  Skip TLS cert verification when using TLS

- Send messages
  - `zing message send bob -m "Hello"`
  - From stdin: `echo "Hi" | zing message send bob`
  - If neither `-m` nor stdin is provided, your editor opens (uses `$VISUAL` or `$EDITOR`, falls back to `nano`).

- Read messages
  - `zing message read`
  - Paging options:
    - `--page-size <n>` (default 50, max 1000)
    - `--page-token <token>` (returned by previous page)
  - Output options:
    - `--no-color` to disable color
    - `--width <n>` to wrap to a specific terminal width
    - You can also export `NO_COLOR=1` to disable colorization globally.

- Clear messages
  - `zing message clear`

- Config helpers
  - `zing config edit` (opens your config in `$VISUAL`/`$EDITOR`)

- Version
  - `zing version`

## Configuration

Zing uses a simple TOML config and environment variables (via Viper). The config file is created when you register or log in for the first time.

Config file locations (OS defaults):
- Linux:   `~/.config/zing/config.toml`
- macOS:   `~/Library/Application Support/zing/config.toml`
- Windows: `%AppData%/zing/config.toml`

Keys you may see/set:

```
server_addr = ""        # Set automatically on register/login (e.g., "localhost:50051")
token = ""              # Session token stored after register/login
plaintext = false        # Set when you pass --plaintext
insecure = false         # Set when you pass --insecure (TLS only)

[redis]
addr = ""               # e.g., "127.0.0.1:6379"
username = ""
password = ""
db = 0
tls = false
```

Environment variable equivalents are supported; use upper‑case with underscores:
- `SERVER_ADDR`, `TOKEN`, `PLAINTEXT`, `INSECURE`
- `REDIS_ADDR`, `REDIS_USERNAME`, `REDIS_PASSWORD`, `REDIS_DB`, `REDIS_TLS`

## Deployment

Docker
- Build the image: `docker build -t zing .`
- Run the server:
  - `docker run --rm -p 50051:50051 -e REDIS_ADDR=host.docker.internal:6379 zing`

Kubernetes (Helm)
- A Helm chart is included under `charts/zing/` (Redis is listed as a dependency).
- Example (local chart):
  - `helm install zing ./charts/zing --set redis.enabled=true`
  - Configure values (address, ports, etc.) via `charts/zing/values.yaml`.

## Security notes
- The server binary in this repository starts without TLS; for production, terminate TLS in a reverse proxy or service mesh, or extend the server to use TLS credentials.
- Passwords are hashed with bcrypt on the server.
- Session tokens are stored in Redis with a 7‑day TTL; expired sessions prompt a re‑login.

## Troubleshooting
- Connection refused to Redis: set `REDIS_ADDR` or configure `[redis].addr` correctly.
- Unable to connect with TLS to the default server: use `--plaintext` for local development, or run the server behind TLS.
- Color output looks wrong: pass `--no-color` or set `NO_COLOR=1`.
- Pager issues: set `PAGER` to your preferred pager (defaults to `less -FRX`).
- Editor not found: set `VISUAL` or `EDITOR` (defaults to `nano`).

## License

GPL‑3.0. See `LICENSE` for details.
