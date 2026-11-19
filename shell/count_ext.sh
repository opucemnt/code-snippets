#!/usr/bin/env bash
# 按扩展名统计：数一数当前目录下每种后缀各有多少文件。
# 用法: ./count_ext.sh [目录]，默认当前目录
set -euo pipefail

dir="${1:-.}"
find "$dir" -maxdepth 1 -type f -name '*.*' -printf '%f\n' |
  sed 's/.*\.//' |
  sort |
  uniq -c |
  sort -rn
