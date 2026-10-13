// 命令行倒计时。
//
// 运行: node countdown.js [秒数, 默认 10]
const total = parseInt(process.argv[2] || "10", 10);
let left = total;

const timer = setInterval(() => {
  process.stdout.write(`\r剩余 ${left} 秒...  `);
  left -= 1;
  if (left < 0) {
    clearInterval(timer);
    console.log("\n时间到!");
  }
}, 1000);
