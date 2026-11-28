#!/bin/bash
# CSV 求和：对 CSV 里指定列（按列号）的数值求和，跳过表头。
#
# 用法: bash sum_csv.sh <csv 文件> <列号，从 1 开始>
if [ $# -ne 2 ]; then
  echo "用法: $0 <csv 文件> <列号>"
  exit 1
fi
awk -F, -v col="$2" 'NR > 1 && $col ~ /^-?[0-9.]+$/ {sum += $col; n++} END {printf "第 %d 列：共 %d 行数值，总和 %s\n", col, n, sum}' "$1"
