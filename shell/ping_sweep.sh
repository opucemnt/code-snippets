#!/usr/bin/env bash
# 批量 ping 一个 C 段, 输出存活的主机 (32 并发)。
#
# 用法: ./ping_sweep.sh 192.168.1
set -euo pipefail

prefix="${1:?用法: $0 <网段前缀, 如 192.168.1>}"

for i in $(seq 1 254); do
  if ping -c 1 -W 1 "${prefix}.${i}" >/dev/null 2>&1; then
    echo "${prefix}.${i} 存活"
  fi &
  # 每 32 个等待一波, 控制并发
  if (( i % 32 == 0 )); then
    wait
  fi
done
wait
echo "扫描完成"
