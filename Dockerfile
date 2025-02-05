# 指定基础镜像
FROM ubuntu:24.04

# 安装依赖
RUN sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list \
    && apt-get clean && apt-get update && apt-get install -y --fix-missing \
    git \
    wget \
    golang-go \
    protobuf-compiler \
    vim \
    make \
    sudo \
    curl

# CGO 环境
RUN apt-get install -y --fix-missing \
    build-essential \
    clang

# 设置代理, 解决 go install 的网络问题
ENV GOPROXY=https://proxy.golang.com.cn,direct

# 安装 proto 相关二进制
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
