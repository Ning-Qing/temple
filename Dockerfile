# 定义基础镜像版本
ARG GO_VERSION=1.17.13
ARG ALPINE_VERSION=3.16

# 构建运行环境
FROM alpine:${ALPINE_VERSION} AS base
# 换源，同步时区
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

# 构建编译时环境
FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS build

# 设置go env
ENV GOPATH="/go" \
    GO111MODULE="on" \
    GOPROXY="https://goproxy.cn" \
    CGO_ENABLED="0"

# 换源，安装编译依赖
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && apk add --no-cache \
    git \
    make

COPY . /go/src/github.com/Ning-Qing/temple
WORKDIR /go/src/github.com/Ning-Qing/temple
RUN make build

# 拷贝二进制文件
FROM base

WORKDIR /fabric-relayer
COPY --from=build /go/src/github.com/Ning-Qing/temple/build/fabric-relayer ./temple
ENTRYPOINT ["./temple"]

# 使用CMD传参
CMD ["-h"]