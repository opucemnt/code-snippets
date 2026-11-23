#!/usr/bin/env python3
"""计时装饰器：给函数包一层，跑完打印耗时。

用法:
    python3 stopwatch.py
"""
import functools
import time


def stopwatch(func):
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        start = time.perf_counter()
        result = func(*args, **kwargs)
        elapsed = time.perf_counter() - start
        print(f"{func.__name__} took {elapsed:.3f}s")
        return result

    return wrapper


@stopwatch
def slow_job():
    time.sleep(0.2)
    return sum(range(100000))


if __name__ == "__main__":
    print("result:", slow_job())
