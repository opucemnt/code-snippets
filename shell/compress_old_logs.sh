#!/usr/bin/env bash
# 压缩 7 天前的日志文件, 节省磁盘空间。
#
# 用法: ./compress_old_logs.sh [日志目录, 默认 /var/log/myapp]
set -euo pipefail

logdir="${1:-/var/log/myapp}"

if [[ ! -d "$logdir" ]]; then
  echo "目录不存在: $logdir"
  exit 1
fi

find "$logdir" -type f -name "*.log" -mtime +7 -print0 \
  | xargs -0 -r gzip -9

echo "压缩完成:"
du -sh "$logdir"
