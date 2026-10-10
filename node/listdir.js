// 列出目录下所有文件及其大小 (人类可读)。
//
// 运行: node listdir.js [目录, 默认当前目录]
const fs = require("fs");
const path = require("path");

const dir = process.argv[2] || ".";

function human(bytes) {
  const units = ["B", "KB", "MB", "GB"];
  let i = 0;
  while (bytes >= 1024 && i < units.length - 1) {
    bytes /= 1024;
    i += 1;
  }
  return `${bytes.toFixed(1)} ${units[i]}`;
}

for (const name of fs.readdirSync(dir)) {
  const full = path.join(dir, name);
  const stat = fs.statSync(full);
  const type = stat.isDirectory() ? "d" : "-";
  console.log(`${type} ${human(stat.size).padStart(10)}  ${name}`);
}
