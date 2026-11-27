// 计算文件的 MD5 值（流式读取，大文件也不怕）。
//
// 运行: node md5sum.js <文件路径> [文件路径 ...]
const fs = require("fs");
const crypto = require("crypto");

const files = process.argv.slice(2);
if (!files.length) {
  console.log("用法: node md5sum.js <文件路径> [文件路径 ...]");
  process.exit(1);
}
for (const file of files) {
  try {
    const hash = crypto.createHash("md5");
    hash.update(fs.readFileSync(file));
    console.log(`${hash.digest("hex")}  ${file}`);
  } catch (e) {
    console.error(`读取失败 ${file}: ${e.message}`);
  }
}
