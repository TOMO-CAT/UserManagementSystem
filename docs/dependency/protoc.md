# protoc

已经在 Dockerfile 中增加了，如果使用 docker 的话可以跳过这一步：

```
sudo apt install protobuf-compiler
```

还需要安装 `protoc-gen-go`，用于生成 `xx.pb.go` 文件：

```bash
# 网络不通时使用代理:
# export GOPROXY=https://proxy.golang.com.cn,direct

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

再安装 `protoc-gen-go-grpc` 插件，用于生成 `xx_grpc.pb.go` 文件：

```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

此时这两个插件会安装在 `~/go/bin` 下，我们可以在 `~/.bashrc` 中加入一行：

```bash
export PATH=~/go/bin:$PATH
```

检查插件是否安装成功：

```bash
$ protoc-gen-go --version
protoc-gen-go v1.31.0
```
