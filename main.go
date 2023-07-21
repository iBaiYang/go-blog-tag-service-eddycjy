package main

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/iBaiYang/go-blog-tag-service-eddycjy/proto"
	"github.com/iBaiYang/go-blog-tag-service-eddycjy/server"
	"github.com/soheilhy/cmux"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
	"strings"
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

//var port string
//
//func init() {
//	flag.StringVar(&port, "port", "8003", "启动端口号")
//	flag.Parse()
//}
//
//func RunTCPServer(port string) (net.Listener, error) {
//	return net.Listen("tcp", ":"+port)
//}
//
//func RunGrpcServer() *grpc.Server {
//	s := grpc.NewServer()
//	pb.RegisterTagServiceServer(s, server.NewTagServer())
//	reflection.Register(s)
//
//	return s
//}
//
//func RunHttpServer(port string) *http.Server {
//	serveMux := http.NewServeMux()
//	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
//		_, _ = w.Write([]byte(`pong`))
//	})
//
//	return &http.Server{
//		Addr:    ":" + port,
//		Handler: serveMux,
//	}
//}
//
//func main() {
//	l, err := RunTCPServer(port)
//	if err != nil {
//		log.Fatalf("Run TCP Server err: %v", err)
//	}
//
//	m := cmux.New(l)
//	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldPrefixSendSettings("content-type", "application/grpc"))
//	httpL := m.Match(cmux.HTTP1Fast())
//
//	grpcS := RunGrpcServer()
//	httpS := RunHttpServer(port)
//	go grpcS.Serve(grpcL)
//	go httpS.Serve(httpL)
//
//	err = m.Serve()
//	if err != nil {
//		log.Fatalf("Run Serve err: %v", err)
//	}
//}

var port string

func init() {
	flag.StringVar(&port, "port", "8004", "启动端口号")
	flag.Parse()
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func RunServer(port string) error {
	httpMux := runHttpServer()
	grpcS := runGrpcServer()

	//endpoint := "0.0.0.0:" + port
	//runtime.HTTPError = grpcGatewayError
	//gwmux := runtime.NewServeMux()

	gatewayMux := runGrpcGatewayServer()

	httpMux.Handle("/", gatewayMux)

	return http.ListenAndServe(":"+port, grpcHandlerFunc(grpcS, httpMux))
}

func runHttpServer() *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))
	})

	return serveMux
}

func runGrpcServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	return s
}

func runGrpcGatewayServer() *runtime.ServeMux {
	endpoint := "0.0.0.0:" + port
	gwmux := runtime.NewServeMux()
	dopts := []grpc.DialOption{grpc.WithInsecure()}
	_ = pb.RegisterTagServiceHandlerFromEndpoint(context.Background(), gwmux, endpoint, dopts)

	return gwmux
}

func main() {
	err := RunServer(port)
	if err != nil {
		log.Fatalf("Run Serve err: %v", err)
	}
}

type httpError struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func grpcGatewayError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	s, ok := status.FromError(err)
	if !ok {
		s = status.New(codes.Unknown, err.Error())
	}

	httpError := httpError{Code: int32(s.Code()), Message: s.Message()}
	details := s.Details()
	for _, detail := range details {
		if v, ok := detail.(*pb.Error); ok {
			httpError.Code = v.Code
			httpError.Message = v.Message
		}
	}

	resp, _ := json.Marshal(httpError)
	w.Header().Set("Content-type", marshaler.ContentType())
	w.WriteHeader(runtime.HTTPStatusFromCode(s.Code()))
	_, _ = w.Write(resp)
}
