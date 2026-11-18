#!/usr/bin/env python3
"""快速排序：取中间元素做基准，递归分治。

用法:
    python3 quicksort.py
"""
import random


def quicksort(items):
    if len(items) <= 1:
        return items
    pivot = items[len(items) // 2]
    left = [x for x in items if x < pivot]
    mid = [x for x in items if x == pivot]
    right = [x for x in items if x > pivot]
    return quicksort(left) + mid + quicksort(right)


if __name__ == "__main__":
    data = [random.randint(0, 99) for _ in range(12)]
    print("before:", data)
    print("after: ", quicksort(data))
