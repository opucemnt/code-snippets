#!/usr/bin/env python3
"""合并 JSON：把多个 JSON 对象文件合并成一个，后面的覆盖前面的。

用法:
    python3 merge_json.py a.json b.json [-o out.json]
"""
import argparse
import json


def main():
    parser = argparse.ArgumentParser(description="合并多个 JSON 对象文件")
    parser.add_argument("files", nargs="+", help="要合并的 JSON 文件")
    parser.add_argument("-o", "--output", default=None, help="输出文件，默认打印")
    args = parser.parse_args()

    merged = {}
    for path in args.files:
        with open(path, encoding="utf-8") as f:
            data = json.load(f)
        if not isinstance(data, dict):
            raise SystemExit(f"{path} 不是 JSON 对象，跳过合并")
        merged.update(data)

    text = json.dumps(merged, ensure_ascii=False, indent=2)
    if args.output:
        with open(args.output, "w", encoding="utf-8") as f:
            f.write(text + "\n")
    else:
        print(text)


if __name__ == "__main__":
    main()
