#!/usr/bin/env python3
"""打印前 n 个斐波那契数。

用法:
    python3 fibonacci.py [n]
"""
import sys


def fibs(n):
    out = []
    a, b = 0, 1
    for _ in range(n):
        out.append(a)
        a, b = b, a + b
    return out


if __name__ == "__main__":
    n = int(sys.argv[1]) if len(sys.argv) > 1 else 10
    print(fibs(n))
