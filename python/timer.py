#!/usr/bin/env python3
"""命令行倒计时：每秒打印剩余秒数，到 0 提示完成。

用法:
    python3 timer.py [秒数，默认 10]
"""
import sys
import time


def main():
    secs = int(sys.argv[1]) if len(sys.argv) > 1 else 10
    for i in range(secs, 0, -1):
        print(f"\r剩余 {i} 秒 ", end="", flush=True)
        time.sleep(1)
    print("\r时间到！")


if __name__ == "__main__":
    main()
