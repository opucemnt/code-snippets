#!/usr/bin/env bash
# 清理旧文件：删掉指定目录下超过 N 天没改过的文件。
# 默认只列出会删哪些，加 --apply 才真正删除。
# 用法: ./clean_tmp.sh [目录] [天数] [--apply]，默认 /tmp 7 天
set -euo pipefail

dir="${1:-/tmp}"
days="${2:-7}"
apply=0
[[ "${3:-}" == "--apply" ]] && apply=1

while IFS= read -r -d '' f; do
  if [[ "$apply" -eq 1 ]]; then
    rm -f "$f"
    echo "已删除：$f"
  else
    echo "会删除：$f（加 --apply 执行）"
  fi
done < <(find "$dir" -maxdepth 1 -type f -mtime +"$days" -print0)
