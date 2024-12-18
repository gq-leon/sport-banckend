FROM golang:1.23.3-alpine AS builder

WORKDIR /app

#ENV GOPROXY=https://goproxy.cn,direct

# 复制go.mod和go.sum文件先于源代码，以利用缓存
COPY go.mod go.sum ./

# 下载Go依赖，使用大陆加速器
RUN go mod download

# 复制整个项目源代码
COPY . .

# 编译Go程序，关闭CGO，减少依赖，静态链接
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sport-backend ./cmd/main.go

# 第二阶段：运行阶段，使用Alpine Linux作为基础镜像
FROM alpine:latest

# 安装基本的运行时依赖
RUN apk --no-cache add ca-certificates tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata

# 从构建阶段复制必要文件到运行阶段
COPY --from=builder /app/sport-backend ./
COPY --from=builder /app/.env ./

# 设置环境变量
ENV GIN_MODE=release

# 设置容器启动命令
CMD ["./sport-backend"]