#!/usr/bin/env python3
"""命令行倒计时。

用法:
    python3 countdown.py [秒数, 默认 10]

每秒刷新一行显示剩余时间, 结束时打印提示。
"""
import sys
import time


def main():
    seconds = int(sys.argv[1]) if len(sys.argv) > 1 else 10
    if seconds <= 0:
        print("秒数必须是正整数")
        sys.exit(1)
    for remaining in range(seconds, 0, -1):
        print(f"\r剩余 {remaining:>3d} 秒...", end="", flush=True)
        time.sleep(1)
    print("\r时间到!            ")


if __name__ == "__main__":
    main()
