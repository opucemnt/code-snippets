#!/usr/bin/env python3
"""代码行数统计工具。

用法:
    python3 loc.py <目录> [扩展名, 默认 .py]

递归统计目录下指定扩展名文件的总行数、空行数和注释行数。
"""
import os
import sys


def count_file(path):
    total = blank = comment = 0
    with open(path, encoding="utf-8", errors="ignore") as f:
        for line in f:
            total += 1
            s = line.strip()
            if not s:
                blank += 1
            elif s.startswith("#"):
                comment += 1
    return total, blank, comment


def main():
    if len(sys.argv) < 2:
        print(f"用法: {sys.argv[0]} <目录> [扩展名]")
        sys.exit(1)
    root = sys.argv[1]
    ext = sys.argv[2] if len(sys.argv) > 2 else ".py"
    t = b = c = n = 0
    for dirpath, _, filenames in os.walk(root):
        for name in filenames:
            if name.endswith(ext):
                ft, fb, fc = count_file(os.path.join(dirpath, name))
                t, b, c, n = t + ft, b + fb, c + fc, n + 1
    print(f"文件数: {n}")
    print(f"总行数: {t} (空行 {b}, 注释 {c}, 代码 {t - b - c})")


if __name__ == "__main__":
    main()
