# gRPC Streaming Example

[![Go Report Card](https://goreportcard.com/badge/github.com/diptesh2082/grpc-streaming-example)](https://goreportcard.com/report/github.com/diptesh2082/grpc-streaming-example)
[![GoDoc](https://godoc.org/github.com/diptesh2082/grpc-streaming-example?status.svg)](https://godoc.org/github.com/diptesh2082/grpc-streaming-example)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This project demonstrates all four types of gRPC streaming in Go:

## Features

- **Unary RPC:** Simple request-response pattern (SayHello)
- **Server-Side Streaming:** Server streams multiple responses (StreamStockPricesServer)
- **Client-Side Streaming:** Client streams multiple requests (StreamStockPricesClient)
- **Bidirectional Streaming:** Both client and server stream messages (StreamStockPricesBi)

## Prerequisites

- Go 1.21 or later
- Protocol Buffers compiler (`protoc`)
- Go plugins for Protocol Buffers

## Installation

1. Install the required Go plugins:
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

2. Generate gRPC code from proto files:
```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/user.proto
```

## Project Structure
```
.
├── client/
│   ├── main.go
│   ├── grpc_client.go
│   └── stock_client.go
├── proto/
│   ├── user.proto
│   ├── user.pb.go
│   └── user_grpc.pb.go
└── server/
    ├── grpc_server.go
    └── stock_server.go
```

## Running the Application

1. Start the server:
```bash
go run server/grpc_server.go
```

2. In another terminal, run the client:
```bash
go run client/main.go
```

## Example Usage

### Server-Side Streaming
```go
// Client code
stream, err := client.StreamStockPricesServer(context.Background(), &pb.StockRequest{
    StockSymbol: "AAPL",
})
```

### Client-Side Streaming
```go
// Client code
stream, err := client.StreamStockPricesClient(context.Background())
```

### Bidirectional Streaming
```go
// Client code
stream, err := client.StreamStockPricesBi(context.Background())
```

## License

MIT License

Copyright (c) 2024 Diptesh Mandal

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
