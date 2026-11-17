#!/usr/bin/env python3
"""随机密码生成器：默认 16 位，可指定长度。

用法:
    python3 password.py [长度，默认 16]
"""
import secrets
import string
import sys


def main():
    length = int(sys.argv[1]) if len(sys.argv) > 1 else 16
    alphabet = string.ascii_letters + string.digits + "!@#$%^&*"
    print("".join(secrets.choice(alphabet) for _ in range(length)))


if __name__ == "__main__":
    main()
