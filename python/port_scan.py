#!/usr/bin/env python3
"""简易 TCP 端口扫描器。

用法:
    python3 port_scan.py <host> <起始端口> <结束端口>

对目标主机的端口范围做快速 TCP connect 扫描, 输出开放的端口。
仅用于自己拥有或授权测试的主机。
"""
import socket
import sys


def scan(host, start, end, timeout=0.5):
    open_ports = []
    for port in range(start, end + 1):
        with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
            s.settimeout(timeout)
            if s.connect_ex((host, port)) == 0:
                open_ports.append(port)
                print(f"[+] {port}/tcp open")
    return open_ports


def main():
    if len(sys.argv) != 4:
        print(f"用法: {sys.argv[0]} <host> <起始端口> <结束端口>")
        sys.exit(1)
    host = sys.argv[1]
    start, end = int(sys.argv[2]), int(sys.argv[3])
    print(f"扫描 {host} 端口 {start}-{end} ...")
    found = scan(host, start, end)
    print(f"完成, 开放端口: {found}")


if __name__ == "__main__":
    main()
