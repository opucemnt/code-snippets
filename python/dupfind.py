#!/usr/bin/env python3
"""重复文件查找：按内容哈希找出目录下内容相同的文件。

用法:
    python3 dupfind.py <目录>
"""
import hashlib
import os
import sys
from collections import defaultdict


def file_hash(path):
    h = hashlib.md5()
    with open(path, "rb") as f:
        for chunk in iter(lambda: f.read(65536), b""):
            h.update(chunk)
    return h.hexdigest()


def main():
    if len(sys.argv) != 2:
        print(f"用法: {sys.argv[0]} <目录>")
        sys.exit(1)
    groups = defaultdict(list)
    for root, _dirs, files in os.walk(sys.argv[1]):
        for name in files:
            path = os.path.join(root, name)
            groups[file_hash(path)].append(path)
    found = False
    for digest, paths in groups.items():
        if len(paths) > 1:
            found = True
            print(f"重复组 {digest[:8]}:")
            for p in paths:
                print(f"  {p}")
    if not found:
        print("没有发现重复文件")


if __name__ == "__main__":
    main()
