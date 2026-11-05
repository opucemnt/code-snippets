// CSV 转 JSON：第一行是表头，输出 JSON 数组。
//
// 运行: node csv2json.js <csv 文件>
const fs = require("fs");

const file = process.argv[2];
if (!file) {
  console.log("用法: node csv2json.js <csv 文件>");
  process.exit(1);
}
const rows = fs.readFileSync(file, "utf8").trim().split("\n").map((l) => l.split(","));
const [header, ...body] = rows;
const out = body.map((row) => Object.fromEntries(header.map((h, i) => [h.trim(), (row[i] || "").trim()])));
console.log(JSON.stringify(out, null, 2));
