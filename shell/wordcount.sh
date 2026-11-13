#!/bin/bash
# 词频统计：输出文件的行数、单词数、字节数。
#
# 用法: bash wordcount.sh <文件路径>
if [ $# -ne 1 ]; then
  echo "用法: $0 <文件路径>"
  exit 1
fi
if [ ! -f "$1" ]; then
  echo "文件不存在: $1"
  exit 1
fi
lines=$(grep -c . "$1")
words=$(wc -w < "$1")
bytes=$(wc -c < "$1")
echo "行数: $lines  单词: $words  字节: $bytes"
