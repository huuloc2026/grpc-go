# gRPC with Golang: Basic Implementation Tutorial

This tutorial guides you through implementing a basic gRPC service in Go.

## Prerequisites

- Go installed ([Download](https://golang.org/dl/))
- `protoc` Protocol Buffers compiler ([Install Guide](https://grpc.io/docs/protoc-installation/))
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins:

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## 1. Define the Service

Create a `proto/helloworld.proto` file:

```proto
syntax = "proto3";

package helloworld;

service Greeter {
    rpc SayHello (HelloRequest) returns (HelloReply) {}
}

message HelloRequest {
    string name = 1;
}

message HelloReply {
    string message = 1;
}
```

## 2. Generate Go Code

```sh
protoc --go_out=. --go-grpc_out=. proto/helloworld.proto
```

## 3. Implement the Server

Create `server/main.go`:

```go
package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    pb "github.com/huuloc2026/grpc-go/github.com/huuloc2026/grpc-demo/greeterpb"
)

type server struct {
    pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterGreeterServer(s, &server{})
    log.Println("Server listening at :50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
```

## 4. Implement the Client

Create `client/main.go`:

```go
package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "github.com/huuloc2026/grpc-go/github.com/huuloc2026/grpc-demo/greeterpb"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewGreeterClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "World"})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Greeting: %s", r.Message)
}
```

## 5. Run the Example

1. Start the server:

     ```sh
     go run server/main.go
     ```

2. In another terminal, run the client:

     ```sh
     go run client/main.go
     ```

Test
```sh
go test ./server
```


You should see the greeting message from the server.

---

For more, see the [gRPC-Go Quick Start](https://grpc.io/docs/languages/go/quickstart/).