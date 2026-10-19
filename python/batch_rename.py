#!/usr/bin/env python3
"""批量文件重命名工具。

用法:
    python3 batch_rename.py <目录> <前缀>

把目录下的文件按序号统一重命名为 "<前缀>_001.扩展名" 的风格。
只重命名文件, 跳过子目录。
"""
import os
import sys


def main():
    if len(sys.argv) != 3:
        print(f"用法: {sys.argv[0]} <目录> <前缀>")
        sys.exit(1)
    directory, prefix = sys.argv[1], sys.argv[2]
    if not os.path.isdir(directory):
        print(f"错误: 目录不存在: {directory}")
        sys.exit(1)

    count = 0
    for i, name in enumerate(sorted(os.listdir(directory)), start=1):
        src = os.path.join(directory, name)
        if not os.path.isfile(src):
            continue
        _, ext = os.path.splitext(name)
        dst = os.path.join(directory, f"{prefix}_{i:03d}{ext}")
        os.rename(src, dst)
        print(f"{name} -> {os.path.basename(dst)}")
        count += 1
    print(f"共重命名 {count} 个文件")


if __name__ == "__main__":
    main()
