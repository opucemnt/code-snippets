// 文件监听：文件内容变化时打印时间和新大小。
//
// 运行: node filewatch.js <文件路径>
const fs = require("fs");

const file = process.argv[2];
if (!file) {
  console.log("用法: node filewatch.js <文件路径>");
  process.exit(1);
}
let last = fs.statSync(file).mtimeMs;
console.log("正在监听", file, "按 Ctrl+C 退出");
fs.watchFile(file, { interval: 500 }, (curr) => {
  if (curr.mtimeMs !== last) {
    last = curr.mtimeMs;
    console.log(new Date().toLocaleTimeString(), "变化，新大小:", curr.size, "字节");
  }
});
