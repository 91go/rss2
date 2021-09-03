FROM golang:1.16 AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o rss2 .

FROM scratch AS releaser
COPY --from=builder /build/rss2 /
EXPOSE 8090
CMD ["/rss2"]