#!/usr/bin/env bash
# 目录备份脚本: 把指定目录打包成带日期的 tar.gz, 只保留最近 7 份。
#
# 用法: ./backup.sh <要备份的目录> [存放目录, 默认 /tmp/backup]
set -euo pipefail

src="${1:?用法: $0 <要备份的目录> [存放目录]}"
dest="${2:-/tmp/backup}"
name="$(basename "$src")"
stamp="$(date +%Y%m%d-%H%M%S)"

mkdir -p "$dest"
tar -czf "$dest/${name}-${stamp}.tar.gz" -C "$(dirname "$src")" "$name"
echo "已备份到: $dest/${name}-${stamp}.tar.gz"

# 只保留最近 7 份
ls -t "$dest"/"${name}"-*.tar.gz 2>/dev/null | tail -n +8 | xargs -r rm -f
echo "清理完成, 只保留最近 7 份"
