#!/bin/bash
# 检查主机连通：ping 一次看通不通，再试 TCP 80 端口。
#
# 用法: bash ping_host.sh <主机> [端口，默认 80]
if [ $# -lt 1 ]; then
  echo "用法: $0 <主机> [端口]"
  exit 1
fi
host="$1"
port="${2:-80}"
if ping -c 1 -W 2 "$host" >/dev/null 2>&1; then
  echo "$host ping 通"
else
  echo "$host ping 不通"
fi
if (exec 3<>"/dev/tcp/$host/$port") 2>/dev/null; then
  echo "$host:$port TCP 连通"
  exec 3>&- 3<&-
else
  echo "$host:$port TCP 不通"
fi
