#!/usr/bin/env bash
# 最新文件：打印指定目录里最近修改的一个文件。
# 用法: ./latest.sh [目录]，默认当前目录
set -euo pipefail

dir="${1:-.}"
latest="$(find "$dir" -maxdepth 1 -type f -printf '%T@ %p\n' | sort -rn | head -n 1 | cut -d' ' -f2-)"
if [[ -z "$latest" ]]; then
  echo "目录里没有文件：$dir" >&2
  exit 1
fi
echo "$latest"
