package main

import (
	"context"
	pb "github.com/iBaiYang/go-blog-tag-service-eddycjy/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	/*验证客户端是否与服务端建立起了连接*/
	//ctx := context.Background()
	//clientConn, _ := GetClientConn(ctx, "localhost:8004", nil)
	//defer clientConn.Close()

	/*验证客户端是否与服务端建立起了连接，阻塞等待连接*/
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
	resp, _ := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequest{Name: "Go"}) // 发起指定 RPC 方法的调用

	log.Printf("resp: %v", resp)
}

/*
创建给定目标的客户端连接
*/
func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	// 禁用了此 ClientConn 的传输安全性验证
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}
