#!/bin/bash
# 按日期前缀重命名：给目录下所有文件加上 YYYYMMDD_ 前缀。
#
# 用法: bash rename_date.sh <目录>
dir="$1"
if [ -z "$dir" ] || [ ! -d "$dir" ]; then
  echo "用法: $0 <目录>"
  exit 1
fi
prefix=$(date +%Y%m%d)
for f in "$dir"/*; do
  [ -f "$f" ] || continue
  base=$(basename "$f")
  case "$base" in
    "$prefix"_*) continue ;;
  esac
  mv "$f" "$dir/${prefix}_${base}"
  echo "已重命名: $base -> ${prefix}_${base}"
done
