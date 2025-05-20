#!/bin/sh
set -e

if [ "$1" = "docker" ]; then
  echo "以Docker方式启动..."
  docker run -d --name gin-init -p 40020:40020 --env-file=deploy/env/.env.prod gin-init:latest
else
  echo "以本地二进制方式启动..."
  ./gin-init
fi 