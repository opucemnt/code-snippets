// 读取文件并统计行数、单词数。
//
// 运行: node count.js <文件路径>
const fs = require("fs");

const file = process.argv[2];
if (!file) {
  console.error("用法: node count.js <文件路径>");
  process.exit(1);
}

const text = fs.readFileSync(file, "utf8");
const lines = text.split("\n").length;
const words = text.split(/\s+/).filter(Boolean).length;
console.log(`文件: ${file}`);
console.log(`行数: ${lines}, 单词数: ${words}`);
