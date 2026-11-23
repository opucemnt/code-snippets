// 计时包装：给函数包一层，跑完打印耗时。
// 用法: node stopwatch.js

function stopwatch(name, fn) {
  return async (...args) => {
    const start = performance.now();
    const result = await fn(...args);
    console.log(`${name} took ${((performance.now() - start) / 1000).toFixed(3)}s`);
    return result;
  };
}

const slowJob = stopwatch("slowJob", async () => {
  await new Promise((r) => setTimeout(r, 200));
  return Array.from({ length: 100000 }, (_, i) => i).reduce((a, b) => a + b, 0);
});

slowJob().then((result) => console.log("result:", result));
