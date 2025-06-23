package main

import (
	"context"
	"testing"

	pb "github.com/huuloc2026/grpc-go/github.com/huuloc2026/grpc-demo/greeterpb"
)

func TestSayHello(t *testing.T) {
	s := &greeterServer{}

	req := &pb.HelloRequest{Name: "TestUser"}
	resp, err := s.SayHello(context.Background(), req)

	if err != nil {
		t.Fatalf("SayHello returned error: %v", err)
	}
	if resp.GetMessage() != "Hello, TestUser" {
		t.Errorf("Unexpected response: got %q, want %q", resp.GetMessage(), "Hello, TestUser")
	}
}
