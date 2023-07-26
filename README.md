# go-blog-tag-service-eddycjy

无意中发现了这本书，[Go 语言编程之旅](https://golang2.eddycjy.com/)，
这本书中的 [示例代码](https://github.com/go-programming-tour-book) 地址也找到了。

书的作者 煎鱼 的博客地址是 [煎鱼博客](https://eddycjy.com/) ，
博客的 [Github地址](https://github.com/eddycjy/blog) ，
作者的 [Github地址](https://github.com/eddycjy)。

里面看到了 [RPC应用] 的相关内容。

这是一个使用 gRPC 服务通过 HTTP 调用我们 go-gin-blog-eddycjy 的博客后端服务，以此来获得标签列表的业务数据的学习示例。

## 准备

win10下，在github <https://github.com/protocolbuffers/protobuf/releases/ >下载 Protocol Buffers 。
解压后，设置环境变量：`G:\WorkSoft\protoc\bin` 。重启，查看版本(libprotoc 23.4)：
```
> protoc --version
```

protoc 插件安装：
> go get -u github.com/golang/protobuf/protoc-gen-go@v1.3.2

到目录 `G:\GoPath\pkg\mod\github.com\golang\protobuf@v1.3.2\protoc-gen-go` 下执行 `go build .`，
把生成的 `protoc-gen-go.exe` 移动到目录 `G:\GoPath\bin` 下。

## 命令

### 包安装命令

gRPC 库安装：
> go get -u google.golang.org/grpc@v1.29.1

编译和生成 proto 文件：
> protoc --go_out=plugins=grpc:. ./proto/*.proto 

grpcurl 调试工具：
> go get github.com/fullstorydev/grpcurl
>
> go install github.com/fullstorydev/grpcurl/cmd/grpcurl

在win10上安装grpcurl会出问题，可以直接到目录 `G:\GoPath\pkg\mod\github.com\fullstorydev\grpcurl@v1.8.7\cmd\grpcurl` 下，
修改 `grpcurl.go` 文件名为 `main.go`, 然后执行 `go build .`，把生成的 `grpcurl.exe` 移动到目录 `G:\GoPath\bin` 下，
别忘了把文件名修改回来，修改 `main.go` 文件名为 `grpcurl.go`。

cmux 多路复用：
> go get -u github.com/soheilhy/cmux@v0.1.4

protoc-gen-grpc-gateway 同端口同方法提供双流量支持插件安装：
> go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.14.5

在win10上，直接到目录 `G:\GoPath\pkg\mod\github.com\grpc-ecosystem\grpc-gateway@v1.16.0\protoc-gen-grpc-gateway` 下，
然后执行 `go build .`，把生成的 `protoc-gen-grpc-gateway.exe` 移动到目录 `G:\GoPath\bin` 下。



### 操作命令

获取该服务的 RPC 方法列表信息：
> grpcurl -plaintext localhost:8001 list

获取自定义的 RPC Service 方法：
> grpcurl -plaintext localhost:8001 list proto.TagService

调用 RPC 方法：
* 不指定参数
> grpcurl -plaintext localhost:8001 proto.TagService.GetTagList
* 指定参数
> grpcurl -plaintext -d {\"name\":\"Go\"} localhost:8001 proto.TagService.GetTagList

服务间调用：
> go run client\client.go

不同端口同时监听：
> grpcurl -plaintext localhost:8001 proto.TagService.GetTagList
>
> curl localhost:8002/ping

同一端口同时监听：
> grpcurl -plaintext localhost:8003 proto.TagService.GetTagList
> 
> curl localhost:8003/ping

重新编译 proto 文件：
```
protoc -I/usr/local/include -I. \
       -I$GOPATH/src \
       -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
       --grpc-gateway_out=logtostderr=true:. \
       ./proto/*.proto
```

在win10上，如下操作
```
protoc -I. -IG:\GoPath -IG:\GoPath\pkg\mod\github.com\grpc-ecosystem\grpc-gateway@v1.16.0\third_party\googleapis --grpc-gateway_out=logtostderr=true:. ./proto/*.proto
```



