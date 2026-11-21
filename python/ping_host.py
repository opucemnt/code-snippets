#!/usr/bin/env python3
"""检查主机连通：对给定的主机做 TCP 连接测试（默认 80 端口）。

用法:
    python3 ping_host.py <主机> [端口，默认 80] [超时秒，默认 3]
"""
import socket
import sys


def main():
    if len(sys.argv) < 2:
        print(f"用法: {sys.argv[0]} <主机> [端口] [超时秒]")
        sys.exit(1)
    host = sys.argv[1]
    port = int(sys.argv[2]) if len(sys.argv) > 2 else 80
    timeout = float(sys.argv[3]) if len(sys.argv) > 3 else 3
    try:
        with socket.create_connection((host, port), timeout=timeout):
            print(f"{host}:{port} 连通")
    except OSError as e:
        print(f"{host}:{port} 不通：{e}")


if __name__ == "__main__":
    main()
