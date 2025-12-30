#!/bin/sh
set -e

echo "编译Go项目..."
go build -o gin-init main.go

echo "构建Docker镜像..."
docker build -f deploy/docker/Dockerfile -t gin-init:latest .
