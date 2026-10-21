#!/usr/bin/env python3
"""猜数字小游戏。

程序随机生成 1-100 的整数, 玩家输入猜测, 程序提示偏大/偏小,
猜中后输出尝试次数。

运行: python3 guess_number.py
"""
import random


def main():
    secret = random.randint(1, 100)
    tries = 0
    print("猜数字游戏: 我想了一个 1-100 的数, 来猜猜看!")
    while True:
        try:
            guess = int(input("请输入你的猜测: "))
        except ValueError:
            print("请输入一个整数")
            continue
        tries += 1
        if guess < secret:
            print("小了, 再大一点")
        elif guess > secret:
            print("大了, 再小一点")
        else:
            print(f"猜中了! 一共用了 {tries} 次")
            break


if __name__ == "__main__":
    main()
