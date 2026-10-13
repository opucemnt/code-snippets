#!/usr/bin/env bash
# 磁盘使用率检查: 超过阈值则输出告警, 可接入 cron 定时跑。
#
# 用法: ./disk_alert.sh [阈值%, 默认 80]
set -euo pipefail

threshold="${1:-80}"

df -h --output=target,pcent | tail -n +2 | while read -r target pcent; do
  used="${pcent%%%}"
  if (( used >= threshold )); then
    echo "告警: ${target} 磁盘使用率 ${pcent}, 超过阈值 ${threshold}%"
  fi
done
echo "检查完成"
