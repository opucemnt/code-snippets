#!/bin/bash
# 判断字符串是不是回文（正反读都一样）
# 用法: bash palindrome.sh [字符串]
s=${1:-level}
rev=$(echo "$s" | rev)
if [ "$s" = "$rev" ]; then
    echo "$s 是回文"
else
    echo "$s 不是回文"
fi
