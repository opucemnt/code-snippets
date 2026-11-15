#!/bin/bash
# 端口监听检查：检查给定端口是否在监听，并显示占用进程。
#
# 用法: bash portcheck.sh <端口1> [端口2 ...]
if [ $# -eq 0 ]; then
  echo "用法: $0 <端口1> [端口2 ...]"
  exit 1
fi
for port in "$@"; do
  if ss -tlnp 2>/dev/null | grep -q ":$port "; then
    echo "端口 $port: 监听中"
    ss -tlnp 2>/dev/null | grep ":$port " | head -1
  else
    echo "端口 $port: 未监听"
  fi
done
