package main

import (
    "context"
    "log"
    "net"
    "google.golang.org/grpc"

    pb "github.com/omar-khawaja/grpc-nomad-demo/hellopb"
)

type server struct {}

func (s *server) Say(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
    msg := "Whats up " + in.Name + "!"
    return &pb.HelloResponse{Message: msg}, nil
}

func main() {
    l, err := net.Listen("tcp", "0.0.0.0:8080")
    if err != nil {
        log.Println(err)
        return
    }
    s := grpc.NewServer()
    pb.RegisterHelloServer(s, &server{})
    s.Serve(l)
}
