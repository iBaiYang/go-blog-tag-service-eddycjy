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

gRPC 库安装：
> go get -u google.golang.org/grpc@v1.29.1

编译和生成 proto 文件：
> protoc --go_out=plugins=grpc:. ./proto/*.proto 

