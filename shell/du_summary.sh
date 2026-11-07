#!/bin/bash
# 目录占用汇总：列出指定目录下各子目录的大小，按从大到小排序。
#
# 用法: bash du_summary.sh [目录，默认当前目录] [显示条数，默认 10]
dir="${1:-.}"
n="${2:-10}"
du -sh "$dir"/*/ 2>/dev/null | sort -rh | head -n "$n"
