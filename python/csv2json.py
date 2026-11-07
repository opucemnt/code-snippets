#!/usr/bin/env python3
"""CSV 转 JSON：第一行是表头，输出 JSON 数组。

用法:
    python3 csv2json.py <csv 文件>
"""
import csv
import json
import sys


def main():
    if len(sys.argv) != 2:
        print(f"用法: {sys.argv[0]} <csv 文件>")
        sys.exit(1)
    with open(sys.argv[1], newline="", encoding="utf-8") as f:
        rows = list(csv.DictReader(f))
    print(json.dumps(rows, ensure_ascii=False, indent=2))


if __name__ == "__main__":
    main()
