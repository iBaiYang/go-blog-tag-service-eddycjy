package main

import (
	"flag"
	pb "github.com/iBaiYang/go-blog-tag-service-eddycjy/proto"
	"github.com/iBaiYang/go-blog-tag-service-eddycjy/server"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

//var port string
//
//func init() {
//	flag.StringVar(&port, "p", "8000", "启动端口号")
//	flag.Parse()
//}
//
//func main() {
//	s := grpc.NewServer()
//	pb.RegisterTagServiceServer(s, server.NewTagServer())
//	reflection.Register(s)
//
//	lis, err := net.Listen("tcp", ":"+port)
//	if err != nil {
//		log.Fatalf("net.Listen err: %v", err)
//	}
//
//	err = s.Serve(lis)
//	if err != nil {
//		log.Fatalf("server.Serve err: %v", err)
//	}
//}

//var grpcPort string
//var httpPort string
//
//func init() {
//	flag.StringVar(&grpcPort, "grpc_port", "8001", "gRPC 启动端口号")
//	flag.StringVar(&httpPort, "http_port", "9001", "HTTP 启动端口号")
//	flag.Parse()
//}
//
//func main() {
//	errs := make(chan error)
//	go func() {
//		err := RunHttpServer(httpPort)
//		if err != nil {
//			errs <- err
//		}
//	}()
//	go func() {
//		err := RunGrpcServer(grpcPort)
//		if err != nil {
//			errs <- err
//		}
//	}()
//
//	select {
//	case err := <-errs:
//		log.Fatalf("Run Server err: %v", err)
//	}
//}
//
//func RunHttpServer(port string) error {
//	serveMux := http.NewServeMux()
//	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
//		_, _ = w.Write([]byte(`pong`))
//	})
//
//	return http.ListenAndServe(":"+port, serveMux)
//}
//
//func RunGrpcServer(port string) error {
//	s := grpc.NewServer()
//	pb.RegisterTagServiceServer(s, server.NewTagServer())
//	reflection.Register(s)
//	lis, err := net.Listen("tcp", ":"+port)
//	if err != nil {
//		return err
//	}
//
//	return s.Serve(lis)
//}

var port string

func init() {
	flag.StringVar(&port, "port", "8003", "启动端口号")
	flag.Parse()
}

func RunTCPServer(port string) (net.Listener, error) {
	return net.Listen("tcp", ":"+port)
}

func RunGrpcServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	return s
}

func RunHttpServer(port string) *http.Server {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))
	})

	return &http.Server{
		Addr:    ":" + port,
		Handler: serveMux,
	}
}

func main() {
	l, err := RunTCPServer(port)
	if err != nil {
		log.Fatalf("Run TCP Server err: %v", err)
	}

	m := cmux.New(l)
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldPrefixSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())

	grpcS := RunGrpcServer()
	httpS := RunHttpServer(port)
	go grpcS.Serve(grpcL)
	go httpS.Serve(httpL)

	err = m.Serve()
	if err != nil {
		log.Fatalf("Run Serve err: %v", err)
	}
}
