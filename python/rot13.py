#!/usr/bin/env python3
"""ROT13：字母表轮换 13 位，加一次是加密，再加一次就解回来了。

用法:
    python3 rot13.py <文本>
    echo hello | python3 rot13.py
"""
import codecs
import sys


def main():
    if len(sys.argv) > 1:
        text = " ".join(sys.argv[1:])
    else:
        text = sys.stdin.read()
    print(codecs.encode(text, "rot_13"), end="")


if __name__ == "__main__":
    main()
