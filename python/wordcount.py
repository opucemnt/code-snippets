#!/usr/bin/env python3
"""词频统计：读文件，输出行数、单词数、字符数。

用法:
    python3 wordcount.py <文件路径>
"""
import sys


def main():
    if len(sys.argv) != 2:
        print(f"用法: {sys.argv[0]} <文件路径>")
        sys.exit(1)
    with open(sys.argv[1], encoding="utf-8") as f:
        text = f.read()
    lines = sum(1 for line in text.splitlines() if line)
    words = len(text.split())
    print(f"行数: {lines}  单词: {words}  字符: {len(text)}")


if __name__ == "__main__":
    main()
