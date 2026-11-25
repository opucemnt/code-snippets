// 倒序输出文件：把文件的每一行按相反顺序打印。
//
// 运行: node reverse_lines.js <文件路径>
const fs = require("fs");

const file = process.argv[2];
if (!file) {
  console.log("用法: node reverse_lines.js <文件路径>");
  process.exit(1);
}
const lines = fs.readFileSync(file, "utf8").split("\n");
if (lines.length && lines[lines.length - 1] === "") lines.pop();
for (let i = lines.length - 1; i >= 0; i--) {
  console.log(lines[i]);
}
