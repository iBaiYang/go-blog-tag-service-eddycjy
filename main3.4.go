package main

/*初步运行一个gRPC服务*/

import (
	pb "github.com/iBaiYang/go-blog-tag-service-eddycjy/proto"
	"github.com/iBaiYang/go-blog-tag-service-eddycjy/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	var port string = "8001"
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("server.Serve err: %v", err)
	}
}
