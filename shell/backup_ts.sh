#!/bin/bash
# 带时间戳备份：把目录打包成带日期的 tar.gz，放到备份目录。
#
# 用法: bash backup_ts.sh <源目录> [备份目录，默认 ./backups]
set -e
src="$1"
dest="${2:-./backups}"
if [ -z "$src" ] || [ ! -d "$src" ]; then
  echo "用法: $0 <源目录> [备份目录]"
  exit 1
fi
mkdir -p "$dest"
name=$(basename "$src")
stamp=$(date +%Y%m%d_%H%M%S)
tar -czf "$dest/${name}_${stamp}.tar.gz" -C "$(dirname "$src")" "$name"
echo "已备份到 $dest/${name}_${stamp}.tar.gz"
