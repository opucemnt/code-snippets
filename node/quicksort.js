// 快速排序：取中间元素做基准，递归分治。
// 用法: node quicksort.js

function quicksort(items) {
  if (items.length <= 1) return items;
  const pivot = items[Math.floor(items.length / 2)];
  const left = items.filter((x) => x < pivot);
  const mid = items.filter((x) => x === pivot);
  const right = items.filter((x) => x > pivot);
  return [...quicksort(left), ...mid, ...quicksort(right)];
}

const data = Array.from({ length: 12 }, () => Math.floor(Math.random() * 100));
console.log("before:", data);
console.log("after: ", quicksort(data));
