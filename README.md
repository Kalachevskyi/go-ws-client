# go-ws-client
The WebSocket command-line client, written on Golang, 
intended for debugging applications that use WebSocket technology. 
With the ability to output formatted data in different formats (for example, JSON).
Multi-line data entry is supported.

P.S.: Tests are in the development process.
### Preconditions:
- Installed golangci-lint [How to...](https://github.com/golangci/golangci-lint#install)
- Golang version < 1.11

## Command Line Tool:
```
$ ws-client --help
  The WebSocket command-line client, written on Golang, 
  intended for debugging applications that use WebSocket technology.
  
  Usage:
    ws-client [flags]
  
  Flags:
    -d, --delimiter string   Delimiter for a multi-line request. (default ";")
    -f, --format string      Format the webSocket response (Only json is available at the moment)
    -h, --help               help for ws-client
    -m, --multi_line         Multi-line request.
    -u, --url string         Connect to WebSocket host.

```

## Installing dependencies:
- Run ```go mod download``` (it will install all dependencies from go.mod)

## Build binary:
- Run ```make build``` (it will build binary in ./bin/ws-client)

## Linters:

To start the linters, run the following command:

- Run `make lint`