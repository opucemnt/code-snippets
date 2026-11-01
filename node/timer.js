// 命令行倒计时：每秒打印剩余秒数，到 0 提示完成。
//
// 运行: node timer.js [秒数，默认 10]
let secs = parseInt(process.argv[2], 10) || 10;
const id = setInterval(() => {
  process.stdout.write("\r剩余 " + secs + " 秒 ");
  secs--;
  if (secs < 0) {
    clearInterval(id);
    console.log("\r时间到！");
  }
}, 1000);
