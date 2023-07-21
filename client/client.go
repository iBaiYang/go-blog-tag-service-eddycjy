package main

import (
	"context"
	pb "github.com/iBaiYang/go-blog-tag-service-eddycjy/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	//ctx := context.Background()
	//clientConn, _ := GetClientConn(
	//	ctx,
	//	"localhost:8004",
	//	[]grpc.DialOption{grpc.WithBlock()},
	//)
	//defer clientConn.Close()

	ctx := context.Background()
	clientConn, _ := GetClientConn(ctx, "localhost:8004", nil)
	defer clientConn.Close()

	tagServiceClient := pb.NewTagServiceClient(clientConn)
	resp, _ := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequest{Name: "Go"})

	log.Printf("resp: %v", resp)
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}
