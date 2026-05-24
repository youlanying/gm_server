# =============================================
# Stage 1: 编译阶段 (Builder)
# =============================================
FROM golang:1.24-bookworm AS builder

# 安装 gcc（go-sqlite3 需要 CGO）+ git（go mod 拉取依赖）
RUN apt-get update && apt-get install -y --no-install-recommends gcc git && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# 先拷贝依赖文件，利用 Docker 缓存层
COPY go.mod go.sum ./
RUN go mod download

# 拷贝源码并编译
COPY src/ ./src/
RUN cd src/beegoWeb_main && \
    CGO_ENABLED=1 GOOS=linux go build -o /app/gm_server .

# =============================================
# Stage 2: 运行阶段 (Runtime)
# =============================================
FROM debian:bookworm-slim

# 安装运行时依赖（时区数据 + CA证书）
RUN apt-get update && apt-get install -y --no-install-recommends \
    tzdata \
    ca-certificates \
    libsqlite3-0 \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# 从编译阶段拷贝可执行文件
COPY --from=builder /app/gm_server .

# 拷贝运行所需的配置和资源
COPY configs/   ./configs/
COPY beegoWeb/  ./beegoWeb/

# 创建日志目录和数据目录
RUN mkdir -p log data

# 暴露端口（根据 config.ini）
EXPOSE 2021
EXPOSE 4025
EXPOSE 2020
EXPOSE 3020/udp

# 启动 GM 服务器
CMD ["./gm_server"]
