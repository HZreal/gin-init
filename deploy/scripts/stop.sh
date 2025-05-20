#!/bin/sh
set -e

if [ "$1" = "docker" ]; then
  echo "停止并移除Docker容器..."
  docker stop gin-init && docker rm gin-init
else
  echo "停止本地进程（需手动kill）"
fi 