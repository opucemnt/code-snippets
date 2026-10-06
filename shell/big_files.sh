#!/usr/bin/env bash
# 查找目录下最大的 10 个文件, 按大小倒序输出。
#
# 用法: ./big_files.sh [目录, 默认当前目录]
set -euo pipefail

dir="${1:-.}"
find "$dir" -type f -exec du -h {} + 2>/dev/null \
  | sort -rh \
  | head -n 10
