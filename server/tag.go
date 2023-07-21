package server

import (
	"context"
	"encoding/json"
	"github.com/iBaiYang/go-blog-tag-service-eddycjy/pkg/bapi"
	"github.com/iBaiYang/go-blog-tag-service-eddycjy/pkg/errcode"
	pb "github.com/iBaiYang/go-blog-tag-service-eddycjy/proto"
	"net/http"
)

type TagServer struct{}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	//body, err := http.Get("http://127.0.0.1:8000/api/v1/tags?name=" + r.GetName())
	//if err != nil {
	//	return nil, errcode.TogRPCError(errcode.ErrorGetTagListFail)
	//}

	api := bapi.NewAPI("http://127.0.0.1:8000")
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, err
	}

	tagList := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}

	return &tagList, nil
}
