// 词频统计：读文件，输出行数、单词数、字符数。
//
// 运行: node wordcount.js <文件路径>
const fs = require("fs");

const file = process.argv[2];
if (!file) {
  console.log("用法: node wordcount.js <文件路径>");
  process.exit(1);
}
const text = fs.readFileSync(file, "utf8");
const lines = text.split("\n").filter((l) => l.length > 0).length;
const words = text.split(/\s+/).filter(Boolean).length;
console.log("行数:", lines, " 单词:", words, " 字符:", [...text].length);
