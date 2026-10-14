#!/usr/bin/env python3
"""JSON 格式化工具。

用法:
    cat data.json | python3 json_fmt.py
    python3 json_fmt.py < data.json

从标准输入读取 JSON, 美化后输出到标准输出。
"""
import json
import sys


def main():
    try:
        data = json.load(sys.stdin)
    except json.JSONDecodeError as e:
        print(f"JSON 解析失败: {e}", file=sys.stderr)
        sys.exit(1)
    print(json.dumps(data, ensure_ascii=False, indent=2, sort_keys=True))


if __name__ == "__main__":
    main()
