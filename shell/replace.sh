#!/usr/bin/env bash
# 批量替换：把一批文件里的旧字符串换成新字符串。
# 默认只打印会改哪些文件，加 --apply 才真正写盘。
# 用法: ./replace.sh <旧串> <新串> <文件...> [--apply]
set -euo pipefail

apply=0
args=()
for a in "$@"; do
  if [[ "$a" == "--apply" ]]; then apply=1; else args+=("$a"); fi
done
if [[ "${#args[@]}" -lt 3 ]]; then
  echo "用法: $0 <旧串> <新串> <文件...> [--apply]" >&2
  exit 1
fi
old="${args[0]}"; new="${args[1]}"; files=("${args[@]:2}")

for f in "${files[@]}"; do
  if grep -qF "$old" "$f"; then
    if [[ "$apply" -eq 1 ]]; then
      sed -i "s|$old|$new|g" "$f"
      echo "已替换：$f"
    else
      echo "会替换：$f（加 --apply 执行）"
    fi
  fi
done
